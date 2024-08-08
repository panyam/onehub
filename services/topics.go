package services

import (
	"context"
	"fmt"

	gfn "github.com/panyam/goutils/fn"
	ds "github.com/panyam/onehub/datastore"
	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"github.com/panyam/onehub/obs"
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

func ensureTopicBase(topic *protos.Topic) *protos.Topic {
	if topic.Base == nil {
		topic.Base = &protos.MessageBase{}
	}
	return topic
}

// Create a new Topic
func (s *TopicService) CreateTopic(ctx context.Context, req *protos.CreateTopicRequest) (resp *protos.CreateTopicResponse, err error) {
	ctx, span := obs.Tracer.Start(ctx, "CreateTopic")
	defer span.End()
	ensureTopicBase(req.Topic)
	req.Topic.Base.CreatorId = GetAuthedUser(ctx)
	if req.Topic.Base.CreatorId == "" {
		obs.Logger.InfoContext(ctx, "User is not authenticated to create a topic.")
		return nil, status.Error(codes.PermissionDenied, "User is not authenticated to create a topic.")
	}
	topic := req.Topic
	if topic.Base.Id != "" {
		// see if it already exists
		curr, _ := s.DB.GetTopic(ctx, topic.Base.Id)
		if curr != nil {
			obs.Logger.InfoContext(ctx, "Topic with id already exists", "topicId", topic.Base.Id)
			return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Topic with id '%s' already exists", topic.Base.Id))
		}
	} else {
		topic.Base.Id = s.DB.NewID("Topic")
	}
	if topic.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name not found")
	}

	dbTopic := TopicFromProto(topic)
	err = s.DB.SaveTopic(ctx, dbTopic)
	if err == nil {
		resp = &protos.CreateTopicResponse{
			Topic: TopicToProto(dbTopic),
		}
		// Increatement our topic
		topicCnt.Add(ctx, 1)
	}
	return resp, err
}

// Deletes an topic from our system.
func (s *TopicService) DeleteTopic(ctx context.Context, req *protos.DeleteTopicRequest) (resp *protos.DeleteTopicResponse, err error) {
	ctx, span := obs.Tracer.Start(ctx, "DeleteTopic")
	defer span.End()
	resp = &protos.DeleteTopicResponse{}
	s.DB.DeleteTopic(ctx, req.Id)
	return
}

// Finds and retrieves topics matching the particular criteria.
func (s *TopicService) ListTopics(ctx context.Context, req *protos.ListTopicsRequest) (resp *protos.ListTopicsResponse, err error) {
	ctx, span := obs.Tracer.Start(ctx, "ListTopics")
	defer span.End()
	results, err := s.DB.ListTopics(ctx, "", 100)
	if err != nil {
		return nil, err
	}
	resp = &protos.ListTopicsResponse{Topics: gfn.Map(results, TopicToProto)}
	return
}

func (s *TopicService) GetTopic(ctx context.Context, req *protos.GetTopicRequest) (resp *protos.GetTopicResponse, err error) {
	ctx, span := obs.Tracer.Start(ctx, "GetTopic")
	defer span.End()
	curr, _ := s.DB.GetTopic(ctx, req.Id)
	if curr == nil {
		err = status.Error(codes.NotFound, fmt.Sprintf("Topic with id '%s' not found", req.Id))
	} else {
		resp = &protos.GetTopicResponse{Topic: TopicToProto(curr)}
	}
	return
}

func (s *TopicService) GetTopics(ctx context.Context, req *protos.GetTopicsRequest) (resp *protos.GetTopicsResponse, err error) {
	ctx, span := obs.Tracer.Start(ctx, "GetTopics")
	defer span.End()
	topics := gfn.BatchGet(req.Ids, func(id string) (out *protos.Topic, err error) {
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
	ctx, span := obs.Tracer.Start(ctx, "UpdateTopic")
	defer span.End()
	currtopic, err := s.GetTopic(ctx, &protos.GetTopicRequest{Id: req.Topic.Base.Id})
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
	err = s.DB.SaveTopic(ctx, TopicFromProto(currtopic.Topic))
	if err == nil {
		resp = &protos.UpdateTopicResponse{
			Topic: currtopic.Topic,
		}
	}
	return resp, err
}
