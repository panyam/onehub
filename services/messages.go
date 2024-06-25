package services

import (
	"context"
	"fmt"
	"log"
	"time"

	gut "github.com/panyam/goutils/utils"
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

func (s *MessageService) ListMessages(ctx context.Context, req *protos.ListMessagesRequest) (resp *protos.ListMessagesResponse, err error) {
	results, err := s.DB.GetMessages(req.TopicId, "", req.Pagination.PageKey, int(req.Pagination.PageSize))
	if err != nil {
		return nil, err
	}

	resp = &protos.ListMessagesResponse{Messages: gut.Map(results, MessageToProto)}
	return
}

// Create a new Message
func (s *MessageService) CreateMessages(ctx context.Context, req *protos.CreateMessagesRequest) (resp *protos.CreateMessagesResponse, err error) {
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
		if !req.AllowUserids {
			message.UserId = authedUser
		}
		message.TopicId = req.TopicId
		if message.Id != "" {
			// see if it already exists
			curr, _ := s.DB.GetMessage(message.Id)
			if curr != nil {
				return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Message with id '%s' already exists", message.Id))
			}
		} else {
			numNewIDs++
		}
	}
	var dbmsgs []*ds.Message
	newIDs := s.DB.NewIDs("Message", numNewIDs)
	for _, message := range req.Messages {
		// get an ID from the pool
		if message.Id == "" {
			numNewIDs--
			message.Id = newIDs[numNewIDs]
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
		return nil, err
	}

	resp = &protos.CreateMessagesResponse{
		Messages: gut.Map(dbmsgs, MessageToProto),
	}
	return
}

// Create a new Message
func (s *MessageService) ImportMessages(ctx context.Context, req *protos.ImportMessagesRequest) (resp *protos.ImportMessagesResponse, err error) {
	// Add a new message entity here
	numNewIDs := 0
	for idx, message := range req.Messages {
		if message.TopicId == "" {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("TopicId not set in message %d", idx))
		}
		if message.UserId == "" {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("UserId not set in message %d", idx))
		}
		if message.Id != "" {
			// see if it already exists
			curr, _ := s.DB.GetMessage(message.Id)
			if curr != nil {
				return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Message with id '%s' already exists", message.Id))
			}
		} else {
			numNewIDs++
		}
	}
	var dbmsgs []*ds.Message
	newIDs := s.DB.NewIDs("Message", numNewIDs)
	for _, message := range req.Messages {
		// get an ID from the pool
		if message.Id == "" {
			numNewIDs--
			message.Id = newIDs[numNewIDs]
		}
		dbmsg := MessageFromProto(message)
		dbmsgs = append(dbmsgs, dbmsg)
	}

	if err := s.DB.CreateMessages(dbmsgs); err != nil {
		return nil, err
	}

	resp = &protos.ImportMessagesResponse{
		Messages: gut.Map(dbmsgs, MessageToProto),
	}
	log.Printf("Imported %d messages", len(req.Messages))
	return
}

func (s *MessageService) UpdateMessage(ctx context.Context, req *protos.UpdateMessageRequest) (resp *protos.UpdateMessageResponse, err error) {
	msg, err := s.DB.GetMessage(req.Message.Id)
	if msg == nil {
		return nil, fmt.Errorf("message not found: %s", req.Message.Id)
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
	s.DB.DeleteMessage(req.Id)
	return
}
