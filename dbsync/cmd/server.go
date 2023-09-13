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
)

func main() {
	p := dbsync.PGDBFromEnv()

	selChan := make(chan dbsync.Selection)
	var currSelection dbsync.Selection

	// State of our processing
	lastBegin := -1
	lastCommit := -1

	tsclient := dbsync.NewTSClient("", "")
	log.Println("Created Typesense Client: ", tsclient)

	// Ensure right schemas on TS

	pgmsghandler := dbsync.PGMSGHandler{
		DB: p,
		HandleBeginMessage: func(idx int, msg *pglogrepl.BeginMessage) error {
			lastBegin = idx
			log.Println("Begin Transaction: ", msg)
			return nil
		},
		HandleCommitMessage: func(idx int, msg *pglogrepl.CommitMessage) error {
			lastCommit = -1
			log.Println("Commit Transaction: ", lastBegin, msg)
			return nil
		},
		HandleRelationMessage: func(idx int, msg *pglogrepl.RelationMessage, tableInfo *dbsync.PGTableInfo) error {
			log.Println("Relation Message: ", lastBegin, msg)
			// Make sure we ahve an equivalent TS schema (or we could do this proactively at the start)
			// Typically we wouldnt be doing this when handling log events but rather
			// on startup time
			doctype := fmt.Sprintf("%s.%s", msg.Namespace, msg.RelationName)
			dbsync.EnsureSchema(tsclient, doctype, tableInfo)
			return nil
		},
		HandleInsertMessage: func(idx int, msg *pglogrepl.InsertMessage, reln *pglogrepl.RelationMessage) error {
			log.Println("Insert Message: ", lastBegin, msg, reln)
			// Now write this to our typesense index

			pkey, out, errors := MessageToMap(p, msg.Tuple, reln)
			if errors != nil {
				log.Println("Error converting to map: ", errors)
			}
			log.Println("Converted: ", pkey, out)

			if _, ok := out["created_at"]; ok {
				out["created_at"] = out["created_at"].(time.Time).Unix()
			}
			if _, ok := out["updated_at"]; ok {
				out["updated_at"] = out["updated_at"].(time.Time).Unix()
			}
			doctype := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
			result, err := tsclient.Collection(doctype).Documents().Upsert(out)
			if err != nil {
				schema, err2 := tsclient.Collection(doctype).Retrieve()
				log.Println("Error Upserting: ", result, err)
				log.Println("Old Schema: ", schema, err2)
				panic(err)
			}

			return nil
		},
		HandleDeleteMessage: func(idx int, msg *pglogrepl.DeleteMessage, reln *pglogrepl.RelationMessage) error {
			log.Println("Delete Message: ", lastBegin, msg, reln)
			return nil
		},
		HandleUpdateMessage: func(idx int, msg *pglogrepl.UpdateMessage, reln *pglogrepl.RelationMessage) error {
			log.Println("Update Message: ", lastBegin, msg, reln)
			return nil
		},
	}

	logQueue := dbsync.NewLogQueue(p, func(msgs []dbsync.PGMSG, err error) (numProcessed int, stop bool) {
		log.Println("Curr Selection:", currSelection)
		if err != nil {
			log.Println("Error processing messsages: ", err)
			return 0, false
		}
		for i, rawmsg := range msgs {
			err := pgmsghandler.HandleMessage(i, &rawmsg)
			if err == dbsync.ErrStopProcessingMessages {
				break
			} else if err != nil {
				log.Println("Error handling message: ", i, err)
			}
		}
		if lastCommit < 0 {
			return lastCommit + 1, false
		} else {
			return len(msgs), false
		}
	})
	go logQueue.Start()

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

	// Now we start the syncer.  This is responsible for:
	//  Starting/Stopping the logQueue (above)
	//  Getting Selection requests, executing them (either in a transaction or not)
	for selReq := range selChan {
		logQueue.Stop()
		selReq.Execute()
		currSelection = selReq
		logQueue.Start()
	}
}
