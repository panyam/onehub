package services

import (
	"context"
	"fmt"
	"time"

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

// Create a new Message
func (s *MessageService) CreateMessage(ctx context.Context, req *protos.CreateMessageRequest) (resp *protos.CreateMessageResponse, err error) {
	topic, err := s.DB.GetTopic(req.Message.TopicId)
	if topic == nil {
		return nil, fmt.Errorf("Topic not found: %s", req.Message.TopicId)
	}
	if err != nil {
		return nil, err
	}
	// Add a new message entity here
	message := req.Message
	message.Id = fmt.Sprintf("%s:%d", req.Message.TopicId, time.Now().UnixMilli())
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
