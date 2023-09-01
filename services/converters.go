package services

import (
	"log"

	"github.com/lib/pq"

	ds "github.com/panyam/onehub/datastore"
	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"google.golang.org/protobuf/types/known/structpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

func TopicToProto(input *ds.Topic) (out *protos.Topic) {
	var userIds map[string]bool = make(map[string]bool)
	for _, userId := range input.Users {
		userIds[userId] = true
	}

	out = &protos.Topic{
		CreatedAt: tspb.New(input.BaseModel.CreatedAt),
		UpdatedAt: tspb.New(input.BaseModel.UpdatedAt),
		Name:      input.Name,
		Id:        input.BaseModel.Id,
		CreatorId: input.CreatorId,
		Users:     userIds,
	}
	return
}

func TopicFromProto(input *protos.Topic) (out *ds.Topic) {
	out = &ds.Topic{
		BaseModel: ds.BaseModel{
			CreatedAt: input.CreatedAt.AsTime(),
			UpdatedAt: input.UpdatedAt.AsTime(),
			Id:        input.Id,
		},
		Name:      input.Name,
		CreatorId: input.CreatorId,
	}
	if input.Users != nil {
		var userIds []string
		for userId := range input.Users {
			userIds = append(userIds, userId)
		}

		out.Users = pq.StringArray(userIds)
	}
	return
}

func MessageToProto(input *ds.Message) (out *protos.Message) {
	log.Println("Found DS Msg: ", input)
	out = &protos.Message{
		CreatedAt:   tspb.New(input.CreatedAt),
		UpdatedAt:   tspb.New(input.BaseModel.UpdatedAt),
		Id:          input.BaseModel.Id,
		UserId:      input.UserId,
		TopicId:     input.TopicId,
		ContentType: input.ContentType,
		ContentText: input.ContentText,
	}
	if input.ContentData != nil {
		if data, err := structpb.NewStruct(input.ContentData); err != nil {
			log.Println("Error converting ContentData: ", err)
		} else {
			out.ContentData = data
		}
	}
	return
}

func MessageFromProto(input *protos.Message) (out *ds.Message) {
	out = &ds.Message{
		BaseModel: ds.BaseModel{
			CreatedAt: input.CreatedAt.AsTime(),
			UpdatedAt: input.UpdatedAt.AsTime(),
			Id:        input.Id,
		},
		UserId:      input.UserId,
		TopicId:     input.TopicId,
		CreatedAt:   input.CreatedAt.AsTime(),
		ContentType: input.ContentType,
		ContentText: input.ContentText,
	}
	if input.ContentData != nil {
		out.ContentData = input.ContentData.AsMap()
	}
	return
}

func UserToProto(input *ds.User) (out *protos.User) {
	out = &protos.User{
		CreatedAt: tspb.New(input.CreatedAt),
		UpdatedAt: tspb.New(input.BaseModel.UpdatedAt),
		Name:      input.Name,
		Avatar:    input.Avatar,
		Id:        input.BaseModel.Id,
	}
	return
}

func UserFromProto(input *protos.User) (out *ds.User) {
	out = &ds.User{
		BaseModel: ds.BaseModel{
			CreatedAt: input.CreatedAt.AsTime(),
			UpdatedAt: input.UpdatedAt.AsTime(),
			Id:        input.Id,
		},
		Avatar: input.Avatar,
		Name:   input.Name,
	}
	return
}
