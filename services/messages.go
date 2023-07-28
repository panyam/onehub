package services

import (
	"context"
	"log"

	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

type MessageService struct {
	protos.UnimplementedMessageServiceServer
	*EntityStore[protos.Message]
}

func NewMessageService(estore *EntityStore[protos.Message]) *MessageService {
	if estore == nil {
		estore = NewEntityStore[protos.Message]()
	}
	estore.IDSetter = func(message *protos.Message, id string) { message.Id = id }
	estore.IDGetter = func(message *protos.Message) string { return message.Id }

	estore.CreatedAtSetter = func(message *protos.Message, val *tspb.Timestamp) { message.CreatedAt = val }
	estore.CreatedAtGetter = func(message *protos.Message) *tspb.Timestamp { return message.CreatedAt }

	estore.UpdatedAtSetter = func(message *protos.Message, val *tspb.Timestamp) { message.UpdatedAt = val }
	estore.UpdatedAtGetter = func(message *protos.Message) *tspb.Timestamp { return message.UpdatedAt }

	return &MessageService{
		EntityStore: estore,
	}
}

// Create a new Message
func (s *MessageService) CreateMessage(ctx context.Context, req *protos.CreateMessageRequest) (resp *protos.CreateMessageResponse, err error) {
	resp = &protos.CreateMessageResponse{}
	resp.Message = s.EntityStore.Create(req.Message)
	return
}

// Get a single topic by id
func (s *MessageService) GetMessage(ctx context.Context, req *protos.GetMessageRequest) (resp *protos.GetMessageResponse, err error) {
	log.Println("Getting Message by ID: ", req.Id)
	resp = &protos.GetMessageResponse{
		Message: s.EntityStore.Get(req.Id),
	}
	return
}

// Batch gets multiple messages.
func (s *MessageService) GetMessages(ctx context.Context, req *protos.GetMessagesRequest) (resp *protos.GetMessagesResponse, err error) {
	log.Println("BatchGet for IDs: ", req.Ids)
	resp = &protos.GetMessagesResponse{
		Messages: s.EntityStore.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Message
func (s *MessageService) UpdateMessage(ctx context.Context, req *protos.UpdateMessageRequest) (resp *protos.UpdateMessageResponse, err error) {
	resp = &protos.UpdateMessageResponse{
		Message: s.EntityStore.Update(req.Message),
	}
	return
}

// Deletes an message from our system.
func (s *MessageService) DeleteMessage(ctx context.Context, req *protos.DeleteMessageRequest) (resp *protos.DeleteMessageResponse, err error) {
	resp = &protos.DeleteMessageResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves messages matching the particular criteria.
func (s *MessageService) ListMessages(ctx context.Context, req *protos.ListMessagesRequest) (resp *protos.ListMessagesResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Message) bool {
		return s1.CreatedAt.AsTime() < s2.CreatedAt.AsTime()
	},
		func(m *protos.Message) bool {
			return m.TopicId == req.TopicId
		})
	log.Println("Found Messages: ", results)
	resp = &protos.ListMessagesResponse{Messages: results}
	return
}
