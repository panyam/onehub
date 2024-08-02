package services

import (
	"context"
	"fmt"
	"log"
	"time"

	gfn "github.com/panyam/goutils/fn"
	ds "github.com/panyam/onehub/datastore"
	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageService struct {
	protos.UnimplementedMessageServiceServer
	DB *ds.OneHubDB
}

func NewMessageService(db *ds.OneHubDB) *MessageService {
	return &MessageService{
		DB: db,
	}
}

func ensureMessageBase(msg *protos.Message) *protos.Message {
	if msg.ContentBase == nil {
		msg.ContentBase = &protos.ContentBase{}
	}
	if msg.Base == nil {
		msg.Base = &protos.MessageBase{}
	}
	return msg
}

func (s *MessageService) ListMessages(ctx context.Context, req *protos.ListMessagesRequest) (resp *protos.ListMessagesResponse, err error) {
	ctx, span := tracer.Start(ctx, "ListMessages")
	defer span.End()

	pageKey := ""
	pageSize := 100
	if req.Pagination != nil {
		pageKey = req.Pagination.PageKey
		pageSize = int(req.Pagination.PageSize)
	}
	results, err := s.DB.GetMessages(req.TopicId, "", pageKey, pageSize)
	if err != nil {
		return nil, err
	}

	resp = &protos.ListMessagesResponse{Messages: gfn.Map(results, MessageToProto)}
	return
}

// Create a new Message
func (s *MessageService) CreateMessages(ctx context.Context, req *protos.CreateMessagesRequest) (resp *protos.CreateMessagesResponse, err error) {
	ctx, span := tracer.Start(ctx, "CreateMessages")
	defer span.End()

	topic, err := s.DB.GetTopic(req.TopicId)
	if topic == nil {
		return nil, fmt.Errorf("topic not found: %s", req.TopicId)
	}
	if err != nil {
		return nil, err
	}

	// Add a new message entity here
	numNewIDs := 0
	authedUser := GetAuthedUser(ctx)
	for _, message := range req.Messages {
		// TODO - do this on batch
		ensureMessageBase(message)
		if !req.AllowUserids {
			message.Base.CreatorId = authedUser
		}
		message.TopicId = req.TopicId
		if message.Base.Id != "" {
			// see if it already exists
			curr, _ := s.DB.GetMessage(message.Base.Id)
			if curr != nil {
				logger.InfoContext(ctx, "Message with id already exists", "messageId", message.Base.Id)
				return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Message with id '%s' already exists", message.Base.Id))
			}
		} else {
			numNewIDs++
		}
	}
	var dbmsgs []*ds.Message
	newIDs := s.DB.NewIDs("Message", numNewIDs)
	for _, message := range req.Messages {
		// get an ID from the pool
		if message.Base.Id == "" {
			numNewIDs--
			message.Base.Id = newIDs[numNewIDs]
		}
		dbmsg := MessageFromProto(message)
		dbmsgs = append(dbmsgs, dbmsg)
	}

	// Also set time stamps
	for _, msg := range dbmsgs {
		msg.CreatedAt = time.Now()
		msg.UpdatedAt = time.Now()
	}

	if err := s.DB.CreateMessages(dbmsgs); err != nil {
		logger.ErrorContext(ctx, err.Error())
		return nil, err
	}

	resp = &protos.CreateMessagesResponse{
		Messages: gfn.Map(dbmsgs, MessageToProto),
	}
	messageCnt.Add(ctx, int64(len(dbmsgs)))
	return
}

// Create a new Message
func (s *MessageService) ImportMessages(ctx context.Context, req *protos.ImportMessagesRequest) (resp *protos.ImportMessagesResponse, err error) {
	ctx, span := tracer.Start(ctx, "ImportMessages")
	defer span.End()
	// Add a new message entity here
	numNewIDs := 0
	for idx, message := range req.Messages {
		ensureMessageBase(message)
		if message.TopicId == "" {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("TopicId not set in message %d", idx))
		}
		if message.Base.CreatorId == "" {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("CreatorId not set in message %d", idx))
		}
		if message.Base.Id != "" {
			// see if it already exists
			curr, _ := s.DB.GetMessage(message.Base.Id)
			if curr != nil {
				return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Message with id '%s' already exists", message.Base.Id))
			}
		} else {
			numNewIDs++
		}
	}
	var dbmsgs []*ds.Message
	newIDs := s.DB.NewIDs("Message", numNewIDs)
	for _, message := range req.Messages {
		// get an ID from the pool
		if message.Base.Id == "" {
			numNewIDs--
			message.Base.Id = newIDs[numNewIDs]
		}
		dbmsg := MessageFromProto(message)
		dbmsgs = append(dbmsgs, dbmsg)
	}

	if err := s.DB.CreateMessages(dbmsgs); err != nil {
		return nil, err
	}

	resp = &protos.ImportMessagesResponse{
		Messages: gfn.Map(dbmsgs, MessageToProto),
	}
	log.Printf("Imported %d messages", len(req.Messages))
	return
}

func (s *MessageService) UpdateMessage(ctx context.Context, req *protos.UpdateMessageRequest) (resp *protos.UpdateMessageResponse, err error) {
	ensureMessageBase(req.Message)
	msg, err := s.DB.GetMessage(req.Message.Base.Id)
	if msg == nil {
		return nil, fmt.Errorf("message not found: %s", req.Message.Base.Id)
	}
	dbmsg := MessageFromProto(req.Message)
	if err := s.DB.SaveMessage(dbmsg); err != nil {
		return nil, err
	}
	resp = &protos.UpdateMessageResponse{
		Message: MessageToProto(dbmsg),
	}
	return
}

func (s *MessageService) GetMessage(ctx context.Context, req *protos.GetMessageRequest) (resp *protos.GetMessageResponse, err error) {
	curr, _ := s.DB.GetMessage(req.Id)
	if curr == nil {
		err = status.Error(codes.NotFound, fmt.Sprintf("Message with id '%s' not found", req.Id))
	} else {
		resp = &protos.GetMessageResponse{Message: MessageToProto(curr)}
	}
	return
}

// Deletes an message from our system.
func (s *MessageService) DeleteMessage(ctx context.Context, req *protos.DeleteMessageRequest) (resp *protos.DeleteMessageResponse, err error) {
	resp = &protos.DeleteMessageResponse{}
	err = s.DB.DeleteMessage(req.Id)
	return
}

// Search a new Message
func (s *MessageService) SearchMessages(ctx context.Context, req *protos.SearchMessagesRequest) (resp *protos.SearchMessagesResponse, err error) {
	resp = &protos.SearchMessagesResponse{}
	return
}
