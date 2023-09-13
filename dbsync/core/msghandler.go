package dbsync

import (
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pglogrepl"
)

var ErrStopProcessingMessages = errors.New("message processing halted")

type PGMSGHandler struct {
	DB                    *PGDB
	HandleBeginMessage    func(idx int, msg *pglogrepl.BeginMessage) error
	HandleCommitMessage   func(idx int, msg *pglogrepl.CommitMessage) error
	HandleRelationMessage func(idx int, msg *pglogrepl.RelationMessage, tableInfo *TableInfo) error
	HandleUpdateMessage   func(idx int, msg *pglogrepl.UpdateMessage, reln *pglogrepl.RelationMessage) error
	HandleDeleteMessage   func(idx int, msg *pglogrepl.DeleteMessage, reln *pglogrepl.RelationMessage) error
	HandleInsertMessage   func(idx int, msg *pglogrepl.InsertMessage, reln *pglogrepl.RelationMessage) error
	relnCache             map[uint32]*pglogrepl.RelationMessage
}

func (p *PGMSGHandler) AddRelation(reln *pglogrepl.RelationMessage) *TableInfo {
	if p.relnCache == nil {
		p.relnCache = make(map[uint32]*pglogrepl.RelationMessage)
	}
	p.relnCache[reln.RelationID] = reln

	// Query to get info on pkeys
	tableInfo, _ := p.DB.RefreshTableInfo(reln.RelationID, reln.Namespace, reln.RelationName)
	return tableInfo

	/*
			getpkeyquer := fmt.Sprintf(`SELECT a.attname, format_type(a.atttypid, a.atttypmod) AS data_type
								FROM   pg_index i
								JOIN   pg_attribute a ON a.attrelid = i.indrelid AND a.attnum = ANY(i.indkey)
								WHERE  i.indrelid = 'tablename'::regclass AND i.indisprimary`, reln.Namespace, reln.RelationName)
		log.Println("Query for pkey: ", getpkeyquer)
	*/
}

func (p *PGMSGHandler) GetRelation(relationId uint32) *pglogrepl.RelationMessage {
	if p.relnCache == nil {
		p.relnCache = make(map[uint32]*pglogrepl.RelationMessage)
	}
	reln, ok := p.relnCache[relationId]
	if !ok {
		panic("Could not find relation - Need to query DB manually")
		return nil
	}
	return reln
}

func (p *PGMSGHandler) HandleMessage(idx int, rawmsg *PGMSG) (err error) {
	msgtype := rawmsg.Data[0]
	switch msgtype {
	case 'B':
		if p.HandleBeginMessage != nil {
			var msg pglogrepl.BeginMessage
			msg.Decode(rawmsg.Data[1:])
			return p.HandleBeginMessage(idx, &msg)
		}
	case 'C':
		if p.HandleCommitMessage != nil {
			var msg pglogrepl.CommitMessage
			msg.Decode(rawmsg.Data[1:])
			return p.HandleCommitMessage(idx, &msg)
		}
	case 'R':
		if p.HandleRelationMessage != nil {
			var msg pglogrepl.RelationMessage
			msg.Decode(rawmsg.Data[1:])
			tableInfo := p.AddRelation(&msg)
			return p.HandleRelationMessage(idx, &msg, tableInfo)
		}
	case 'I':
		if p.HandleInsertMessage != nil {
			var msg pglogrepl.InsertMessage
			msg.Decode(rawmsg.Data[1:])
			reln := p.GetRelation(msg.RelationID)
			return p.HandleInsertMessage(idx, &msg, reln)
		}
	case 'D':
		if p.HandleDeleteMessage != nil {
			var msg pglogrepl.DeleteMessage
			msg.Decode(rawmsg.Data[1:])
			reln := p.GetRelation(msg.RelationID)
			return p.HandleDeleteMessage(idx, &msg, reln)
		}
	case 'U':
		if p.HandleUpdateMessage != nil {
			var msg pglogrepl.UpdateMessage
			msg.Decode(rawmsg.Data[1:])
			reln := p.GetRelation(msg.RelationID)
			return p.HandleUpdateMessage(idx, &msg, reln)
		}
	default:
		log.Println(fmt.Sprintf("Processing Messages (%c): ", msgtype), rawmsg)
		panic(fmt.Errorf("invalid Message Type: %c", msgtype))
	}
	return nil
}
