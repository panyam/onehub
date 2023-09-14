package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	dbsync "dbsync/core"

	"github.com/jackc/pglogrepl"
	_ "github.com/lib/pq"
	ohds "github.com/panyam/onehub/clients"
	"github.com/typesense/typesense-go/typesense"
)

type PG2TS struct {
	tsclient      *typesense.Client
	tsclient2     *ohds.TSClient
	pgdb          *dbsync.PGDB
	selChan       chan dbsync.Selection
	currSelection dbsync.Selection
	msghandler    dbsync.PGMSGHandler
}

func NewPG2TS() *PG2TS {
	tsclient := dbsync.NewTSClient("", "")
	out := &PG2TS{
		tsclient: tsclient,
		pgdb:     dbsync.PGDBFromEnv(),
		selChan:  make(chan dbsync.Selection),
	}
	out.msghandler = dbsync.PGMSGHandler{
		DB: out.pgdb,
		HandleRelationMessage: func(m *dbsync.PGMSGHandler, idx int, msg *pglogrepl.RelationMessage, tableInfo *dbsync.PGTableInfo) error {
			log.Println("Relation Message: ", m.LastBegin, msg)
			// Make sure we ahve an equivalent TS schema (or we could do this proactively at the start)
			// Typically we wouldnt be doing this when handling log events but rather
			// on startup time
			doctype := fmt.Sprintf("%s.%s", msg.Namespace, msg.RelationName)
			dbsync.EnsureSchema(tsclient, doctype, tableInfo)
			return nil
		},
		HandleInsertMessage: func(m *dbsync.PGMSGHandler, idx int, msg *pglogrepl.InsertMessage, reln *pglogrepl.RelationMessage) error {
			log.Println("Insert Message: ", m.LastBegin, msg, reln)
			// Now write this to our typesense index

			pkey, out, errors := dbsync.MessageToMap(out.pgdb, msg.Tuple, reln)
			if errors != nil {
				log.Println("Error converting to map: ", pkey, errors)
			}
			// log.Println("Converted: ", pkey, out)

			if _, ok := out["created_at"]; ok {
				out["created_at"] = out["created_at"].(time.Time).Unix()
			}
			if _, ok := out["updated_at"]; ok {
				out["updated_at"] = out["updated_at"].(time.Time).Unix()
			}
			doctype := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
			// result, err := tsclient.Collection(doctype).Documents().Upsert(out)
			result, err := tsclient2.Upsert(doctype, out["id"], out)
			if err != nil {
				schema, err2 := tsclient.Collection(doctype).Retrieve()
				log.Println("Error Upserting: ", result, err)
				log.Println("Old Schema: ", schema, err2)
				panic(err)
			}

			return nil
		},
		HandleDeleteMessage: func(m *dbsync.PGMSGHandler, idx int, msg *pglogrepl.DeleteMessage, reln *pglogrepl.RelationMessage) error {
			tableinfo := out.pgdb.GetTableInfo(reln.RelationID)
			doctype := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
			docid := tableinfo.GetRecordID(msg.OldTuple, reln)
			log.Println(fmt.Sprintf("Delete Message (%s/%s): ", doctype, docid), m.LastBegin, msg, reln)
			doc := tsclient.Collection(doctype).Document(docid)
			result, err := doc.Delete()
			// result, err := tsclient.Collections(doctype).Documents(docid).Delete()
			if err != nil && err.Status != 404 {
				schema, err2 := tsclient.Collection(doctype).Delete()
				log.Println("Error Deleting: ", result, err)
				log.Println("Old Schema: ", schema, err2)
				panic(err)
			}
			return nil
		},
		HandleUpdateMessage: func(m *dbsync.PGMSGHandler, idx int, msg *pglogrepl.UpdateMessage, reln *pglogrepl.RelationMessage) error {
			log.Println("Update Message: ", m.LastBegin, msg, reln)
			return nil
		},
	}
	return out
}

func (p *PG2TS) NewLogQueue() *dbsync.LogQueue {
	logQueue := dbsync.NewLogQueue(p.pgdb, func(msgs []dbsync.PGMSG, err error) (numProcessed int, stop bool) {
		log.Println("Curr Selection:", p.currSelection)
		if err != nil {
			log.Println("Error processing messsages: ", err)
			return 0, false
		}
		for i, rawmsg := range msgs {
			err := p.msghandler.HandleMessage(i, &rawmsg)
			if err == dbsync.ErrStopProcessingMessages {
				break
			} else if err != nil {
				log.Println("Error handling message: ", i, err)
			}
		}
		if p.msghandler.LastCommit > 0 {
			return p.msghandler.LastCommit + 1, false
		} else {
			return len(msgs), false
		}
	})
	go logQueue.Start()
	return logQueue
}

func (p *PG2TS) Start() {
	log.Println("Created Typesense Client: ", p.tsclient)
	// Start log processing
	logQueue := p.NewLogQueue()

	// Now we start the syncer.  This is responsible for:
	//  Starting/Stopping the logQueue (above)
	//  Getting Selection requests, executing them (either in a transaction or not)
	for selReq := range p.selChan {
		logQueue.Stop()
		selReq.Execute()
		p.currSelection = selReq
		logQueue.Start()
	}
}

func main() {
	// State of our processing
	p := NewPG2TS()

	// Ensure right schemas on TS

	// Start a simple http server that listens to commands to control the replicator
	// and to "introduce" selective dumps
	go func() {
		http.HandleFunc("/select", func(w http.ResponseWriter, r *http.Request) {
			// Add a new selection For now we just submit SELECT statements
			// Simple ones - our SELECT query will as a query param in the req
			log.Println("Query: ", r.URL.Query())
			io.WriteString(w, "This is my website!\n")
		})
		if err := http.ListenAndServe(":3333", nil); err != nil {
			panic(err)
		}
	}()

	p.Start()
}
