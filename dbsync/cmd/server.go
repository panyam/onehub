package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	dbsync "dbsync/core"

	"github.com/jackc/pglogrepl"
	_ "github.com/lib/pq"
)

func GetEnvOrDefault(envvar string, defaultValue string) string {
	out := os.Getenv(envvar)
	if out == "" {
		out = defaultValue
	}
	return out
}

const DEFAULT_POSTGRES_HOST = "localhost"
const DEFAULT_POSTGRES_NAME = "onehubdb"
const DEFAULT_POSTGRES_USER = "postgres"
const DEFAULT_POSTGRES_PASSWORD = "docker"
const DEFAULT_POSTGRES_PORT = "54321"

const DEFAULT_DBSYNC_CTRL_NAMESPACE = "dbsync_ctrl"
const DEFAULT_DBSYNC_WM_TABLENAME = "dbsync_wmtable"
const DEFAULT_DBSYNC_PUBNAME = "dbsync_mypub"
const DEFAULT_DBSYNC_REPLSLOT = "dbsync_replslot"

func getConnStr() string {
	dbname := GetEnvOrDefault("POSTGRES_NAME", DEFAULT_POSTGRES_NAME)
	dbhost := GetEnvOrDefault("POSTGRES_HOST", DEFAULT_POSTGRES_HOST)
	dbuser := GetEnvOrDefault("POSTGRES_USER", DEFAULT_POSTGRES_USER)
	dbpassword := GetEnvOrDefault("POSTGRES_PASSWORD", DEFAULT_POSTGRES_PASSWORD)
	dbport := GetEnvOrDefault("POSTGRES_PORT", DEFAULT_POSTGRES_PORT)
	if portval, err := strconv.Atoi(dbport); err != nil {
		panic(err)
	} else {
		return dbsync.ConnStr(dbname, dbhost, portval, dbuser, dbpassword)
	}
}

func setupPGDB() (p *dbsync.PGDB) {
	connstr := getConnStr()
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}

	ctrl_namespace := GetEnvOrDefault("DBSYNC_CTRL_NAMESPACE", DEFAULT_DBSYNC_CTRL_NAMESPACE)
	wm_table_name := GetEnvOrDefault("DBSYNC_WM_TABLENAME", DEFAULT_DBSYNC_WM_TABLENAME)
	pubname := GetEnvOrDefault("DBSYNC_PUBNAME", DEFAULT_DBSYNC_PUBNAME)
	replslot := GetEnvOrDefault("DBSYNC_REPLSLOT", DEFAULT_DBSYNC_REPLSLOT)
	p = &dbsync.PGDB{
		CtrlNamespace: ctrl_namespace,
		WMTableName:   wm_table_name,
		Publication:   pubname,
		ReplSlotName:  replslot,
	}

	// Create publications etc here otherwise Setup will fail
	if err := p.Setup(db); err != nil {
		panic(err)
	}
	return
}

func MessageToMap(p *dbsync.PGDB, msg *pglogrepl.TupleData, reln *pglogrepl.RelationMessage) (pkey string, out map[string]interface{}, errors map[string]error) {
	msgcols := msg.Columns
	relcols := reln.Columns
	if len(msgcols) != len(relcols) {
		log.Printf("Msg cols (%d) and Rel cols (%d) dont match", len(msgcols), len(relcols))
	}
	fullschema := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
	log.Printf("Namespace: %s, RelName: %s, FullSchema: %s", reln.Namespace, reln.RelationName, fullschema)
	pkey = "id"
	if out == nil {
		out = make(map[string]interface{})
	}
	tableinfo := p.GetTableInfo(reln.RelationID)
	for i, col := range reln.Columns {
		val := msgcols[i]
		colinfo := tableinfo.ColInfo[col.Name]
		log.Println("Cols: ", i, col.Name, val, colinfo)
		var err error
		if val.DataType == pglogrepl.TupleDataTypeText {
			out[col.Name], err = colinfo.DecodeText(val.Data)
		} else if val.DataType == pglogrepl.TupleDataTypeBinary {
			out[col.Name], err = colinfo.DecodeBytes(val.Data)
		}
		if err != nil {
			if errors == nil {
				errors = make(map[string]error)
			}
			errors[col.Name] = err
		}
	}
	return
}

func main() {
	p := setupPGDB()

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
		HandleRelationMessage: func(idx int, msg *pglogrepl.RelationMessage, tableInfo *dbsync.TableInfo) error {
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
