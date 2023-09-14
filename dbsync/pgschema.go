package dbsync

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pglogrepl"
)

const PG_TIMESTAMP_FORMAT = "2006-01-02 15:04:05.999999+00"

type PGTableInfo struct {
	RelationID uint32
	ColInfo    map[string]*PGColumnInfo
}

func (t *PGTableInfo) GetRecordID(msg *pglogrepl.TupleData, reln *pglogrepl.RelationMessage) string {
	// TODO - Handle non "id" pkeys and composite pkeys
	msgcols := msg.Columns
	relcols := reln.Columns
	if len(msgcols) != len(relcols) {
		log.Printf("Msg cols (%d) and Rel cols (%d) dont match", len(msgcols), len(relcols))
	}
	// fullschema := fmt.Sprintf("%s.%s", reln.Namespace, reln.RelationName)
	// log.Printf("Namespace: %s, RelName: %s, FullSchema: %s", reln.Namespace, reln.RelationName, fullschema)
	pkey := "id"
	for idx, col := range reln.Columns {
		if col.Name == pkey {
			colinfo, ok := t.ColInfo[col.Name]
			if !ok || colinfo == nil {
				return ""
			}
			return string(msg.Columns[idx].Data)
		}
	}
	return ""
}

type PGColumnInfo struct {
	DBName          string
	Namespace       string
	TableName       string
	ColumnName      string
	ColumnType      string
	OrdinalPosition int
}

func (c *PGColumnInfo) DecodeText(input []byte) (out interface{}, err error) {
	sval := string(input)
	if c.ColumnType == "text" {
		out = sval
	} else if c.ColumnType == "timestamp with time zone" {
		out, err = time.Parse(PG_TIMESTAMP_FORMAT, sval)
	} else if c.ColumnType == "bigint" {
		out, err = strconv.Atoi(sval)
	} else {
		panic(fmt.Errorf("invalid text type: %s", c.ColumnType))
	}
	return
}

func (c *PGColumnInfo) DecodeBytes(input []byte) (out interface{}, err error) {
	panic(fmt.Errorf("invalid binary type: %s", c.ColumnType))
}
