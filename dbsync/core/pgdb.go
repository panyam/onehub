package dbsync

import (
	"database/sql"
	_ "encoding/binary"
	"fmt"
	"log"
	"strconv"
	"time"
	// "github.com/jackc/pgproto3/v2"
)

const PG_TIMESTAMP_FORMAT = "2006-01-02 15:04:05.999999+00"

type TableInfo struct {
	RelationID uint32
	ColInfo    map[string]*ColumnInfo
}

type ColumnInfo struct {
	DBName          string
	Namespace       string
	TableName       string
	ColumnName      string
	ColumnType      string
	OrdinalPosition int
}

func (c *ColumnInfo) DecodeText(input []byte) (out interface{}, err error) {
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

func (c *ColumnInfo) DecodeBytes(input []byte) (out interface{}, err error) {
	panic(fmt.Errorf("invalid binary type: %s", c.ColumnType))
}

type ColumnType uint8

const (
	ColumnTypeInvalid     ColumnType = iota
	ColumnTypeBoolean                // 1 byte
	ColumnTypeSmallInt               // int16
	ColumnTypeInteger                // int32
	ColumnTypeBigInt                 // int64
	ColumnTypeDecimal                // var length precision
	ColumnTypeNumeric                // var length precision
	ColumnTypeReal                   // float
	ColumnTypeDouble                 // float64
	ColumnTypeSmallSerial            // uint16
	ColumnTypeSerial                 // uint32
	ColumnTypeBigSerial              // uint64

	ColumnTypeVarChar // Varying length characters
	ColumnTypeChar    // Fixed length characters
	ColumnTypeText    // text
	ColumnTypeJson    // json

	ColumnTypeBytea // bytearray

	// https://www.postgresql.org/docs/current/datatype-datetime.html
	ColumnTypeTimestamp // timestamp (p) [ with/without timezone ] - 8 bytes
	ColumnTypeDate      // date - 4 bytes
	ColumnTypeTime      // time with/without timezone - 8 or 12 bytes
	ColumnTypeInterval  // interval - 16 bytes

	// https://www.postgresql.org/docs/current/datatype-geometric.html
	ColumnTypePoint   // 16 bytes
	ColumnTypeLine    // 32 bytes
	ColumnTypeLSeg    // 32 bytes
	ColumnTypeBox     // 32 bytes
	ColumnTypePath    // 16 + 16n bytes
	ColumnTypePolygon // 40 + 16n bytes
	ColumnTypeCircle  // 24 bytes
)

var stringToColumnType = map[string]ColumnType{}

func ToColumnType(coltype string) ColumnType {
	return ColumnTypeInvalid
}

func (t ColumnType) String() string {
	switch t {
	case ColumnTypeInvalid:
		return "invalid type"
	case ColumnTypeBoolean:
		return "boolean"
	case ColumnTypeSmallInt:
		return "smallint"
	case ColumnTypeInteger:
		return "integer"
	case ColumnTypeBigInt:
		return "bigint"
	case ColumnTypeDecimal:
		return "decimal"
	case ColumnTypeNumeric:
		return "numeric"
	case ColumnTypeReal:
		return "real"
	case ColumnTypeDouble:
		return "double"
	case ColumnTypeSmallSerial:
		return "smallserial"
	case ColumnTypeSerial:
		return "serial"
	case ColumnTypeBigSerial:
		return "bigserial"
	case ColumnTypeVarChar:
		return "varchar"
	case ColumnTypeChar:
		return "char"
	case ColumnTypeText:
		return "text"
	case ColumnTypeJson:
		return "json"
	case ColumnTypeBytea:
		return "bytea"
	case ColumnTypeTimestamp:
		return "timestamp"
	case ColumnTypeDate:
		return "date"
	case ColumnTypeTime:
		return "time"
	case ColumnTypeInterval:
		return "time"
	default:
		panic(fmt.Sprintf("unknown col type: %d", t))
	}
}

type PGMSG struct {
	LSN  string
	Xid  uint64
	Data []byte
}

type PGDB struct {
	db *sql.DB

	// The name space where we will create sync specific tables (outside user space)
	CtrlNamespace string

	// Name of this sync
	Name string

	// Table where we will write watermarks for this sync
	WMTableName string

	// Which publication are we tracking with this sync?
	Publication string

	// The replication slot we will use to subscribe to change
	// log events from for this sync
	ReplSlotName string

	relnToTableInfo map[uint32]*TableInfo
}

func (p *PGDB) GetTableInfo(relationID uint32) *TableInfo {
	tableinfo, ok := p.relnToTableInfo[relationID]
	if !ok {
		tableinfo = &TableInfo{
			RelationID: relationID,
			ColInfo:    make(map[string]*ColumnInfo),
		}
		p.relnToTableInfo[relationID] = tableinfo
	}
	return tableinfo
}

func (p *PGDB) RefreshTableInfo(relationID uint32, namespace string, table_name string) (tableInfo *TableInfo, err error) {
	field_info_query := fmt.Sprintf(`SELECT table_schema, table_name, column_name, ordinal_position, data_type, table_catalog from information_schema.columns WHERE table_schema = '%s' and table_name = '%s' ;`, namespace, table_name)
	log.Println("Query for field types: ", field_info_query)
	rows, err := p.db.Query(field_info_query)
	if err != nil {
		log.Println("Error getting table info: ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var col ColumnInfo
		if err := rows.Scan(&col.Namespace, &col.TableName, &col.ColumnName, &col.OrdinalPosition, &col.ColumnType, &col.DBName); err != nil {
			log.Println("Could not scan row: ", err)
		} else {
			if p.relnToTableInfo == nil {
				p.relnToTableInfo = make(map[uint32]*TableInfo)
			}
			tableInfo = p.GetTableInfo(relationID)
			if colinfo, ok := tableInfo.ColInfo[col.ColumnName]; !ok {
				tableInfo.ColInfo[col.ColumnName] = &col
			} else {
				colinfo.DBName = col.DBName
				colinfo.Namespace = col.Namespace
				colinfo.TableName = col.TableName
				colinfo.ColumnName = col.ColumnName
				colinfo.ColumnType = col.ColumnType
				colinfo.OrdinalPosition = col.OrdinalPosition
			}
		}
	}
	return
}

func (p *PGDB) Setup(db *sql.DB) (err error) {
	p.db = db
	err = p.ensureNamespace()

	if err == nil {
		err = p.ensureWMTable()
	}

	if err == nil {
		err = p.registerWithPublication()
	}

	if err == nil {
		err = p.setupReplicationSlots()
	}
	return
}

func (p *PGDB) DB() *sql.DB {
	return p.db
}

func (p *PGDB) GetMessages(numMessages int, consume bool, out []PGMSG) (msgs []PGMSG, err error) {
	msgs = out
	changesfuncname := "pg_logical_slot_peek_binary_changes"
	if consume {
		changesfuncname = "pg_logical_slot_get_binary_changes"
	}
	q := fmt.Sprintf(`select * from %s(
					'%s', NULL, %d,
					'publication_names', '%s',
					'proto_version', '1') ;`,
		changesfuncname, p.ReplSlotName, numMessages, p.Publication)
	rows, err := p.db.Query(q)
	if err != nil {
		log.Println("SELECT NAMESPACE ERROR: ", err)
		return nil, err
	}

	for rows.Next() {
		var msg PGMSG
		err = rows.Scan(&msg.LSN, &msg.Xid, &msg.Data)
		if err != nil {
			log.Println("Error scanning change: ", err)
			return
		}
		msgs = append(msgs, msg)
	}
	return
}

func (p *PGDB) Forward(nummsgs int) error {
	changesfuncname := "pg_logical_slot_get_binary_changes"
	q := fmt.Sprintf(`select * from %s('%s', NULL, %d,
					'publication_names', '%s',
					'proto_version', '1') ;`,
		changesfuncname, p.ReplSlotName, nummsgs, p.Publication)
	rows, err := p.db.Query(q)
	if err != nil {
		log.Println("SELECT NAMESPACE ERROR: ", err)
		return err
	}
	// We dont actually need the results
	defer rows.Close()

	// Now update our peek offset
	// peekOffset tells where to do the next "limit" function from
	return nil
}

func (p *PGDB) ensureNamespace() (err error) {
	rows, err := p.db.Query("SELECT * from pg_catalog.pg_namespace where nspname = $1", p.CtrlNamespace)
	if err != nil {
		log.Println("SELECT NAMESPACE ERROR: ", err)
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		// Name space does not exist so create it
		create_schema_query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s AUTHORIZATION CURRENT_USER", p.CtrlNamespace)
		_, err := p.db.Exec(create_schema_query)
		if err != nil {
			log.Println("CREATE SCHEMA ERROR: ", err)
			return err
		}
	}

	return nil
}

func (p *PGDB) ensureWMTable() (err error) {
	// Check if our WM table exists
	rows, err := p.db.Query("SELECT relname, relnamespace, reltype FROM pg_catalog.pg_class WHERE relname = $1 AND relkind = 'r'", p.WMTableName)
	if err != nil {
		log.Println("Get WMTable Error: ", err)
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		// create this table
		create_wmtable_query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
				selectionid varchar(50) PRIMARY KEY,
				low_wm varchar(50),
				high_wm varchar(50)
			)`, p.CtrlNamespace, p.WMTableName)
		_, err = p.db.Exec(create_wmtable_query)
		if err != nil {
			log.Println("WMTable Creation Error: ", err)
			return
		}
	}

	return nil
}

func (p *PGDB) registerWithPublication() error {
	// Now ensure our WM table is assigned to the publication
	q := fmt.Sprintf(`select pubname from pg_publication_tables where schemaname = '%s' and tablename = '%s'`, p.CtrlNamespace, p.WMTableName)
	rows, err := p.db.Query(q)
	if err != nil {
		log.Println("Could not query pb_publication_tables: ", err)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		// There is a row - so make sure our pubname matches the given publication
		// if it doesnt then it means we have an error
		var pubname string
		if err := rows.Scan(&pubname); err != nil {
			log.Println("Error scanning pubname: ", err)
			return err
		}
		if pubname != p.Publication {
			return fmt.Errorf("table %s.%s is already assigned to Publication '%s'", p.CtrlNamespace, p.WMTableName, pubname)
		}
	} else {
		// our table is not part of the publication so add to it
		alterpub := fmt.Sprintf(`ALTER PUBLICATION %s ADD TABLE %s.%s`, p.Publication, p.CtrlNamespace, p.WMTableName)
		_, err := p.db.Exec(alterpub)
		if err != nil {
			log.Println("ALTER PUBLICATION Error : ", err)
			createpubsql := fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE table1, table2, ..., tableN ;", p.Publication)
			log.Printf("Did you create the publication?  Try: %s", createpubsql)
			return err
		}
	}
	return nil
}

/**
 * Create our replication slots and prepare it to be ready for peek/geting events
 * from our publication.  If a slot already exists, then ensures it is a pgoutput type
 */
func (p *PGDB) setupReplicationSlots() error {
	q := fmt.Sprintf(`SELECT slot_name, plugin, slot_type, restart_lsn, confirmed_flush_lsn
			FROM pg_replication_slots
			WHERE slot_name = '%s'`, p.ReplSlotName)
	rows, err := p.db.Query(q)
	if err != nil {
		log.Println("Error Getting Replication Slots: ", err)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		var slot_name string
		var plugin string
		var slot_type string
		var restart_lsn string
		var confirmed_flush_lsn string

		if err := rows.Scan(&slot_name, &plugin, &slot_type, &restart_lsn, &confirmed_flush_lsn); err != nil {
			log.Println("Error scanning slot_name, plugin, plot_type: ", err)
			return err
		}
		if slot_name != p.ReplSlotName {
			return fmt.Errorf("replication slot invalid: %s", p.ReplSlotName)
		}
		if plugin != "pgoutput" {
			return fmt.Errorf("invalid plugin (%s).  Only 'pgoutput' supported", plugin)
		}
		if slot_type != "logical" {
			return fmt.Errorf("invalid replication (%s).  Only 'logical' supported", slot_type)
		}
	} else {
		// Create it
		q := fmt.Sprintf(`SELECT * FROM pg_create_logical_replication_slot('%s', 'pgoutput', false, true);`, p.ReplSlotName)
		rows2, err := p.db.Query(q)
		if err != nil {
			log.Println("SELECT NAMESPACE ERROR: ", err)
			return err
		}
		defer rows2.Close()
		if !rows2.Next() {
			return fmt.Errorf("pg_create_logical_replication_slot returned no rows")
		}
	}
	return nil
}
