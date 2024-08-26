package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pglogrepl"
	_ "github.com/lib/pq"
	"github.com/panyam/dbsync"
	ohds "github.com/panyam/onehub/clients"
	// "github.com/typesense/typesense-go/typesense"
)

type StringMap = map[string]any

type OurMessageHandler struct {
	dbsync.DefaultMessageHandler
	ds *dbsync.Syncer
}

func (m *OurMessageHandler) HandleRelationMessage(idx int, msg *pglogrepl.RelationMessage, tableInfo *dbsync.PGTableInfo) error {
	// Make sure we ahve an equivalent TS schema (or we could do this proactively at the start)
	// Typically we wouldnt be doing this when handling log events but rather
	// on startup time
	// doctype := fmt.Sprintf("%s.%s", msg.Namespace, msg.RelationName)
	// _, fieldMap := PGTableInfoToSchema(tableInfo)
	// log.Println(fmt.Sprintf("Relation Message (%s): ", doctype), m.LastBegin, msg, "Fields: ", fieldMap)
	return nil
}

func (m *OurMessageHandler) HandleInsertMessage(idx int, msg *pglogrepl.InsertMessage, reln *pglogrepl.RelationMessage, tableInfo *dbsync.PGTableInfo) error {
	// log.Println("Insert Message: ", m.LastBegin, msg, reln)
	// Now write this to our typesense index

	pkey, outmap, errors := m.ds.MessageToMap(msg.Tuple, reln)
	if errors != nil {
		log.Println("Error converting to map: ", pkey, errors)
	}
	// log.Println("Converted: ", pkey, out)

	if _, ok := outmap["created_at"]; ok {
		outmap["created_at"] = outmap["created_at"].(time.Time).Unix()
	}
	if _, ok := outmap["updated_at"]; ok {
		outmap["updated_at"] = outmap["updated_at"].(time.Time).Unix()
	}
	doctype := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
	// result, err := tsclient.Collection(doctype).Documents().Upsert(out)
	// docid := outmap["id"].(string)
	docid := tableInfo.GetRecordID(msg.Tuple, reln)

	m.ds.MarkUpdated(doctype, docid, outmap)

	/* - Uncomment to do single gets instead of batch
	result, err := tsclient.Upsert(doctype, docid, outmap)
	if err != nil && clients.TSErrorCode(err) != clients.ErrCodeEntityNotFound {
		schema, err2 := tsclient.GetCollection(doctype)
		log.Println("Error Upserting: ", result, err)
		log.Println("Old Schema: ", schema, err2)
		panic(err)
	}
	*/

	return nil
}

func (m *OurMessageHandler) HandleDeleteMessage(idx int, msg *pglogrepl.DeleteMessage, reln *pglogrepl.RelationMessage, tableInfo *dbsync.PGTableInfo) error {
	// Instead of individual deletes we will batch them by collections
	doctype := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
	docid := tableInfo.GetRecordID(msg.OldTuple, reln)
	m.ds.MarkDeleted(doctype, docid)
	return nil
}
func (m *OurMessageHandler) HandleUpdateMessage(idx int, msg *pglogrepl.UpdateMessage, reln *pglogrepl.RelationMessage, tableInfo *dbsync.PGTableInfo) error {
	// log.Println("Update Message: ", m.LastBegin, msg, reln)
	return nil
}

func main() {
	d, err := dbsync.NewSyncer(
		dbsync.ForTables("users", "messages", "topics"),
	)
	if err != nil {
		panic(err)
	}
	tsclient := tsClient()
	d.Batcher = tsclient
	d.MessageHandler = &OurMessageHandler{ds: d}

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
	d.Start()
}

func tsClient() *ohds.TSClient {
	tsclient := ohds.NewClient("", "")
	// Ensure right schemas on TS
	tsclient.EnsureSchema("public.users", []StringMap{
		{"name": "id", "type": "string"},
		{"name": "version", "type": "int64"},
		{"name": "created_at", "type": "int64"},
		{"name": "updated_at", "type": "int64"},
		{"name": "name", "type": "string"},
		{"name": "avatar", "type": "string", "optional": true},
		{"name": "profile_data", "type": "object", "optional": true},
	})
	tsclient.EnsureSchema("public.messages", []StringMap{
		{"name": "id", "type": "string"},
		{"name": "version", "type": "int64"},
		{"name": "created_at", "type": "int64"},
		{"name": "updated_at", "type": "int64"},
		{"name": "creator_id", "type": "string"},
		{"name": "topic_id", "type": "string"},
		{"name": "parent_id", "type": "string", "optional": true},
		{"name": "source_id", "type": "string", "optional": true},
		{"name": "content_type", "type": "string", "optional": true},
		{"name": "content_text", "type": "string", "optional": true},
		{"name": "content_data", "type": "object", "optional": true},
	})
	tsclient.EnsureSchema("public.topics", []StringMap{
		{"name": "id", "type": "string"},
		{"name": "version", "type": "int64"},
		{"name": "created_at", "type": "int64"},
		{"name": "updated_at", "type": "int64"},
		{"name": "users", "type": "string[]", "optional": true},
	})
	return tsclient
}

/*
func PGTableInfoToSchema(tableInfo *dbsync.PGTableInfo) (fields []StringMap, fieldMap map[string]StringMap) {
	fieldMap = make(map[string]StringMap)
	for colname, colinfo := range tableInfo.ColInfo {
		coltype := colinfo.ColumnType
		field := StringMap{
			"name":     colname,
			"optional": true,
			"type":     "string",
		}
		if coltype == "smallint" || coltype == "integer" {
			field["type"] = "int32"
		} else if colinfo.ColumnType == "bigint" {
			field["type"] = "int64"
		} else if strings.HasPrefix(colinfo.ColumnType, "timestamp") {
			field["type"] = "int64"
		} else if strings.HasPrefix(colinfo.ColumnType, "text") {
			field["type"] = "string"
		} else if strings.HasPrefix(colinfo.ColumnType, "json") {
			field["type"] = "object"
		} else if strings.ToLower(colinfo.ColumnType) == "array" {
			field["type"] = "string[]"
		} else {
			log.Println("Invalid type in decoding: ", colinfo.ColumnType)
			panic("Invalid type")
		}
		fields = append(fields, field)
		fieldMap[colname] = field
	}
	return
}
*/
