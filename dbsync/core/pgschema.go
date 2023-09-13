package dbsync

import (
	"fmt"
	"strconv"
	"time"
)

const PG_TIMESTAMP_FORMAT = "2006-01-02 15:04:05.999999+00"

type PGTableInfo struct {
	RelationID uint32
	ColInfo    map[string]*PGColumnInfo
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
