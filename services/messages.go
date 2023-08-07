package services

import (
	"context"
	"fmt"

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
	results, err := s.DB.GetMessages(req.TopicId, "", req.PageKey, int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp = &protos.ListMessagesResponse{Messages: gut.Map(results, MessageToProto)}
	return
}

// Create a new Message
func (s *MessageService) CreateMessage(ctx context.Context, req *protos.CreateMessageRequest) (resp *protos.CreateMessageResponse, err error) {
	topic, err := s.DB.GetTopic(req.Message.TopicId)
	if topic == nil {
		return nil, fmt.Errorf("topic not found: %s", req.Message.TopicId)
	}
	if err != nil {
		return nil, err
	}
	// Add a new message entity here
	message := req.Message
	if message.Id != "" {
		// see if it already exists
		curr, _ := s.DB.GetMessage(message.Id)
		if curr != nil {
			return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Message with id '%s' already exists", message.Id))
		}
	} else {
		message.Id = s.DB.NextId("Message")
	}
	message.UserId = GetAuthedUser(ctx)
	dbmsg := MessageFromProto(message)
	if err := s.DB.CreateMessage(dbmsg); err != nil {
		return nil, err
	}

	resp = &protos.CreateMessageResponse{
		Message: MessageToProto(dbmsg),
	}
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
