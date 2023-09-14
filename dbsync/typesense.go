package dbsync

import (
	"log"
	"strings"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func NewTSClient(address string, apikey string) (tsclient *typesense.Client) {
	if address == "" {
		address = "http://localhost:8108"
	}
	if apikey == "" {
		apikey = "xyz"
	}
	return typesense.NewClient(
		typesense.WithServer(address),
		typesense.WithAPIKey(apikey),
		typesense.WithConnectionTimeout(5*time.Second),
		typesense.WithCircuitBreakerMaxRequests(50),
		typesense.WithCircuitBreakerInterval(2*time.Minute),
		typesense.WithCircuitBreakerTimeout(1*time.Minute),
	)
}

var TRUE = true
var FALSE = false

func PGTableInfoToSchema(tableInfo *PGTableInfo) (fields []api.Field, fieldMap map[string]api.Field) {
	fieldMap = make(map[string]api.Field)
	for colname, colinfo := range tableInfo.ColInfo {
		coltype := colinfo.ColumnType
		log.Println("C: ", colinfo)
		field := api.Field{
			Name:     colname,
			Optional: &TRUE,
			Type:     "string",
		}
		if coltype == "smallint" || coltype == "integer" {
			field.Type = "int32"
		} else if colinfo.ColumnType == "bigint" {
			field.Type = "int64"
		} else if strings.HasPrefix(colinfo.ColumnType, "timestamp") {
			field.Type = "int64"
		} else if strings.HasPrefix(colinfo.ColumnType, "text") {
			field.Type = "string"
		} else if strings.HasPrefix(colinfo.ColumnType, "json") {
			field.Type = "object"
		} else if strings.ToLower(colinfo.ColumnType) == "array" {
			field.Type = "string[]"
		} else {
			log.Println("Invalid type in decoding: ", colinfo.ColumnType)
			panic("Invalid type")
		}
		fields = append(fields, field)
		fieldMap[colname] = field
	}
	return
}

func EnsureSchema(tsclient *typesense.Client, doctype string, tableInfo *PGTableInfo) {
	fields, fieldMap := PGTableInfoToSchema(tableInfo)
	schema := &api.CollectionSchema{
		Name:               doctype,
		EnableNestedFields: &TRUE,
		Fields:             fields,
	}
	existing, err := tsclient.Collection(doctype).Retrieve()
	if err != nil {
		log.Println("Schema Fetch Error: ", err)
	}
	if existing == nil {
		res, err := tsclient.Collections().Create(schema)
		log.Println("Schema Creation: ", doctype, res, err)
		if err != nil {
			panic(err)
		}
	} else {
		// TODO - check there are *acutally* changes first
		// update it
		var newfields []api.Field
		for _, efield := range existing.Fields {
			newfield, ok := fieldMap[efield.Name]
			if !ok || newfield.Name != efield.Name {
				// New field added
				newfields = append(newfields, newfield)
			} else if newfield.Type != efield.Type || efield.Optional != newfield.Optional {
				// drop and reload it
				newfields = append(newfields, api.Field{
					Drop: &TRUE,
					Name: newfield.Name,
					Type: efield.Type,
				})

				// now added
				newfields = append(newfields, api.Field{
					Name:     newfield.Name,
					Type:     newfield.Type,
					Optional: &TRUE,
				})
			}
		}
		if newfields != nil {
			res, err := tsclient.Collection(doctype).Update(&api.CollectionUpdateSchema{
				Fields: newfields,
			})
			log.Println("Schema Update: ", doctype, res, err)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}
