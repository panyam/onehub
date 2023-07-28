package services

import (
	"context"
	"log"
	"strings"

	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

type TopicService struct {
	protos.UnimplementedTopicServiceServer
	*EntityStore[protos.Topic]
}

func NewTopicService(estore *EntityStore[protos.Topic]) *TopicService {
	if estore == nil {
		estore = NewEntityStore[protos.Topic]()
	}
	estore.IDSetter = func(song *protos.Topic, id string) { song.Id = id }
	estore.IDGetter = func(song *protos.Topic) string { return song.Id }

	estore.CreatedAtSetter = func(song *protos.Topic, val *tspb.Timestamp) { song.CreatedAt = val }
	estore.CreatedAtGetter = func(song *protos.Topic) *tspb.Timestamp { return song.CreatedAt }

	estore.UpdatedAtSetter = func(song *protos.Topic, val *tspb.Timestamp) { song.UpdatedAt = val }
	estore.UpdatedAtGetter = func(song *protos.Topic) *tspb.Timestamp { return song.UpdatedAt }

	return &TopicService{
		EntityStore: estore,
	}
}

// Create a new Topic
func (s *TopicService) CreateTopic(ctx context.Context, req *protos.CreateTopicRequest) (resp *protos.CreateTopicResponse, err error) {
	resp = &protos.CreateTopicResponse{}
	resp.Topic = s.EntityStore.Create(req.Topic)
	return
}

// Batch gets multiple songs.
func (s *TopicService) GetTopics(ctx context.Context, req *protos.GetTopicsRequest) (resp *protos.GetTopicsResponse, err error) {
	log.Println("BatchGet for IDs: ", req.Ids)
	resp = &protos.GetTopicsResponse{
		Topics: s.EntityStore.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Topic
func (s *TopicService) UpdateTopic(ctx context.Context, req *protos.UpdateTopicRequest) (resp *protos.UpdateTopicResponse, err error) {
	resp = &protos.UpdateTopicResponse{
		Topic: s.EntityStore.Update(req.Topic),
	}
	return
}

// Deletes an song from our system.
func (s *TopicService) DeleteTopic(ctx context.Context, req *protos.DeleteTopicRequest) (resp *protos.DeleteTopicResponse, err error) {
	resp = &protos.DeleteTopicResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves songs matching the particular criteria.
func (s *TopicService) ListTopics(ctx context.Context, req *protos.ListTopicsRequest) (resp *protos.ListTopicsResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Topic) bool {
		return strings.Compare(s1.Name, s2.Name) < 0
	})
	log.Println("Found Topics: ", results)
	resp = &protos.ListTopicsResponse{Topics: results}
	return
}
