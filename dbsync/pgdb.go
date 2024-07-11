package dbsync

import (
	"database/sql"
	_ "encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	// "github.com/jackc/pgproto3/v2"
)

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

	relnToPGTableInfo map[uint32]*PGTableInfo
}

const DEFAULT_POSTGRES_HOST = "localhost"
const DEFAULT_POSTGRES_NAME = "onehubdb"
const DEFAULT_POSTGRES_USER = "postgres"
const DEFAULT_POSTGRES_PASSWORD = "docker"
const DEFAULT_POSTGRES_PORT = "5432"

const DEFAULT_DBSYNC_CTRL_NAMESPACE = "dbsync_ctrl"
const DEFAULT_DBSYNC_WM_TABLENAME = "dbsync_wmtable"
const DEFAULT_DBSYNC_PUBNAME = "dbsync_mypub"
const DEFAULT_DBSYNC_REPLSLOT = "dbsync_replslot"

func GetEnvOrDefault(envvar string, defaultValue string) string {
	out := os.Getenv(envvar)
	if out == "" {
		out = defaultValue
	}
	return out
}

func PGDBFromEnv() (p *PGDB) {
	dbname := GetEnvOrDefault("POSTGRES_NAME", DEFAULT_POSTGRES_NAME)
	dbhost := GetEnvOrDefault("POSTGRES_HOST", DEFAULT_POSTGRES_HOST)
	dbuser := GetEnvOrDefault("POSTGRES_USER", DEFAULT_POSTGRES_USER)
	dbpassword := GetEnvOrDefault("POSTGRES_PASSWORD", DEFAULT_POSTGRES_PASSWORD)
	dbport := GetEnvOrDefault("POSTGRES_PORT", DEFAULT_POSTGRES_PORT)
	portval, err := strconv.Atoi(dbport)
	if err != nil {
		panic(err)
	}
	connstr := ConnStr(dbname, dbhost, portval, dbuser, dbpassword)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}

	ctrl_namespace := GetEnvOrDefault("DBSYNC_CTRL_NAMESPACE", DEFAULT_DBSYNC_CTRL_NAMESPACE)
	wm_table_name := GetEnvOrDefault("DBSYNC_WM_TABLENAME", DEFAULT_DBSYNC_WM_TABLENAME)
	pubname := GetEnvOrDefault("DBSYNC_PUBNAME", DEFAULT_DBSYNC_PUBNAME)
	replslot := GetEnvOrDefault("DBSYNC_REPLSLOT", DEFAULT_DBSYNC_REPLSLOT)
	p = &PGDB{
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

func (p *PGDB) GetTableInfo(relationID uint32) *PGTableInfo {
	if p.relnToPGTableInfo == nil {
		p.relnToPGTableInfo = make(map[uint32]*PGTableInfo)
	}
	tableinfo, ok := p.relnToPGTableInfo[relationID]
	if !ok {
		tableinfo = &PGTableInfo{
			RelationID: relationID,
			ColInfo:    make(map[string]*PGColumnInfo),
		}
		p.relnToPGTableInfo[relationID] = tableinfo
	}
	return tableinfo
}

/**
 * Queries the DB for the latest schema of a given relation and stores it
 */
func (p *PGDB) RefreshTableInfo(relationID uint32, namespace string, table_name string) (tableInfo *PGTableInfo, err error) {
	field_info_query := fmt.Sprintf(`SELECT table_schema, table_name, column_name, ordinal_position, data_type, table_catalog from information_schema.columns WHERE table_schema = '%s' and table_name = '%s' ;`, namespace, table_name)
	log.Println("Query for field types: ", field_info_query)
	rows, err := p.db.Query(field_info_query)
	if err != nil {
		log.Println("Error getting table info: ", err)
		return nil, err
	}
	defer rows.Close()
	tableInfo = p.GetTableInfo(relationID)
	for rows.Next() {
		var col PGColumnInfo
		if err := rows.Scan(&col.Namespace, &col.TableName, &col.ColumnName, &col.OrdinalPosition, &col.ColumnType, &col.DBName); err != nil {
			log.Println("Could not scan row: ", err)
		} else {
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
