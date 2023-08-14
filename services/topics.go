package services

import (
	"context"
	"fmt"
	"log"

	gut "github.com/panyam/goutils/utils"
	ds "github.com/panyam/onehub/datastore"
	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TopicService struct {
	protos.UnimplementedTopicServiceServer
	DB *ds.OneHubDB
}

func NewTopicService(db *ds.OneHubDB) *TopicService {
	return &TopicService{
		DB: db,
	}
}

// Create a new Topic
func (s *TopicService) CreateTopic(ctx context.Context, req *protos.CreateTopicRequest) (resp *protos.CreateTopicResponse, err error) {
	req.Topic.CreatorId = GetAuthedUser(ctx)
	if req.Topic.CreatorId == "" {
		return nil, status.Error(codes.PermissionDenied, "User is not authenticated to create a topic.")
	}
	topic := req.Topic
	if topic.Id != "" {
		// see if it already exists
		curr, _ := s.DB.GetTopic(topic.Id)
		if curr != nil {
			return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Topic with id '%s' already exists", topic.Id))
		}
	} else {
		topic.Id = s.DB.NextId("Topic")
	}
	if topic.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name not found")
	}

	dbTopic := TopicFromProto(topic)
	err = s.DB.SaveTopic(dbTopic)
	if err == nil {
		resp = &protos.CreateTopicResponse{
			Topic: TopicToProto(dbTopic),
		}
	}
	return resp, err
}

// Deletes an topic from our system.
func (s *TopicService) DeleteTopic(ctx context.Context, req *protos.DeleteTopicRequest) (resp *protos.DeleteTopicResponse, err error) {
	resp = &protos.DeleteTopicResponse{}
	s.DB.DeleteTopic(req.Id)
	return
}

// Finds and retrieves topics matching the particular criteria.
func (s *TopicService) ListTopics(ctx context.Context, req *protos.ListTopicsRequest) (resp *protos.ListTopicsResponse, err error) {
	results, err := s.DB.ListTopics("", 100)
	if err != nil {
		return nil, err
	}
	log.Println("Found Topics: ", results)
	resp = &protos.ListTopicsResponse{Topics: gut.Map(results, TopicToProto)}
	return
}

func (s *TopicService) GetTopic(ctx context.Context, req *protos.GetTopicRequest) (resp *protos.GetTopicResponse, err error) {
	curr, _ := s.DB.GetTopic(req.Id)
	if curr == nil {
		err = status.Error(codes.NotFound, fmt.Sprintf("Topic with id '%s' not found", req.Id))
	} else {
		resp = &protos.GetTopicResponse{Topic: TopicToProto(curr)}
	}
	return
}

func (s *TopicService) GetTopics(ctx context.Context, req *protos.GetTopicsRequest) (resp *protos.GetTopicsResponse, err error) {
	topics := gut.BatchGet(req.Ids, func(id string) (out *protos.Topic, err error) {
		resp, err := s.GetTopic(ctx, &protos.GetTopicRequest{Id: id})
		if err != nil {
			return nil, err
		}
		return resp.Topic, nil
	})
	resp = &protos.GetTopicsResponse{
		Topics: topics,
	}
	return
}

// Update a new Topic
func (s *TopicService) UpdateTopic(ctx context.Context, req *protos.UpdateTopicRequest) (resp *protos.UpdateTopicResponse, err error) {
	currtopic, err := s.GetTopic(ctx, &protos.GetTopicRequest{Id: req.Topic.Id})
	if err != nil {
		return nil, err
	}

	update_mask := req.UpdateMask
	has_update_mask := update_mask != nil && len(update_mask.Paths) > 0
	if !has_update_mask && len(req.AddUsers) == 0 && len(req.RemoveUsers) == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"update_mask should specify (nested) fields to update *or* add_users *or* remove_users must have IDs to be added/removed")
	}

	if req.UpdateMask != nil {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "name":
				currtopic.Topic.Name = req.Topic.Name
			default:
				return nil, status.Errorf(codes.InvalidArgument, "UpdateTopic - update_mask contains invalid path: %s", path)
			}
		}
	}

	if currtopic.Topic.Users == nil {
		currtopic.Topic.Users = make(map[string]bool)
	}
	for _, userid := range req.AddUsers {
		currtopic.Topic.Users[userid] = true
	}
	for _, userid := range req.RemoveUsers {
		delete(currtopic.Topic.Users, userid)
	}
	err = s.DB.SaveTopic(TopicFromProto(currtopic.Topic))
	if err == nil {
		resp = &protos.UpdateTopicResponse{
			Topic: currtopic.Topic,
		}
	}
	return resp, err
}
