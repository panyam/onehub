package dbsync

import (
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pglogrepl"
)

var ErrStopProcessingMessages = errors.New("message processing halted")

type PGMSGHandler struct {
	HandleBeginMessage    func(idx int, msg *pglogrepl.BeginMessage) error
	HandleCommitMessage   func(idx int, msg *pglogrepl.CommitMessage) error
	HandleRelationMessage func(idx int, msg *pglogrepl.RelationMessage) error
	HandleUpdateMessage   func(idx int, msg *pglogrepl.UpdateMessage, reln *pglogrepl.RelationMessage) error
	HandleDeleteMessage   func(idx int, msg *pglogrepl.DeleteMessage, reln *pglogrepl.RelationMessage) error
	HandleInsertMessage   func(idx int, msg *pglogrepl.InsertMessage, reln *pglogrepl.RelationMessage) error
	relnCache             map[uint32]*pglogrepl.RelationMessage
}

func (p *PGMSGHandler) GetRelation(relationId uint32) *pglogrepl.RelationMessage {
	reln, ok := p.relnCache[relationId]
	if !ok {
		panic("Could not find relation - Need to query DB manually")
		return nil
	}
	return reln
}

func (p *PGMSGHandler) HandleMessage(idx int, rawmsg *PGMSG) (err error) {
	if p.relnCache == nil {
		p.relnCache = make(map[uint32]*pglogrepl.RelationMessage)
	}
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
			p.relnCache[msg.RelationID] = &msg
			return p.HandleRelationMessage(idx, &msg)
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
