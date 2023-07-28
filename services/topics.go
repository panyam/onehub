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
	estore.IDSetter = func(topic *protos.Topic, id string) { topic.Id = id }
	estore.IDGetter = func(topic *protos.Topic) string { return topic.Id }

	estore.CreatedAtSetter = func(topic *protos.Topic, val *tspb.Timestamp) { topic.CreatedAt = val }
	estore.CreatedAtGetter = func(topic *protos.Topic) *tspb.Timestamp { return topic.CreatedAt }

	estore.UpdatedAtSetter = func(topic *protos.Topic, val *tspb.Timestamp) { topic.UpdatedAt = val }
	estore.UpdatedAtGetter = func(topic *protos.Topic) *tspb.Timestamp { return topic.UpdatedAt }

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

// Get a single topic by id
func (s *TopicService) GetTopic(ctx context.Context, req *protos.GetTopicRequest) (resp *protos.GetTopicResponse, err error) {
	log.Println("Getting Topic by ID: ", req.Id)
	resp = &protos.GetTopicResponse{
		Topic: s.EntityStore.Get(req.Id),
	}
	return
}

// Batch gets multiple topics.
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

// Deletes an topic from our system.
func (s *TopicService) DeleteTopic(ctx context.Context, req *protos.DeleteTopicRequest) (resp *protos.DeleteTopicResponse, err error) {
	resp = &protos.DeleteTopicResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves topics matching the particular criteria.
func (s *TopicService) ListTopics(ctx context.Context, req *protos.ListTopicsRequest) (resp *protos.ListTopicsResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Topic) bool {
		return strings.Compare(s1.Name, s2.Name) < 0
	}, nil)
	log.Println("Found Topics: ", results)
	resp = &protos.ListTopicsResponse{Topics: results}
	return
}
