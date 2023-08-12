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

type UserService struct {
	protos.UnimplementedUserServiceServer
	DB *ds.OneHubDB
}

func NewUserService(db *ds.OneHubDB) *UserService {
	return &UserService{
		DB: db,
	}
}

// Create a new User
func (s *UserService) CreateUser(ctx context.Context, req *protos.CreateUserRequest) (resp *protos.CreateUserResponse, err error) {
	user := req.User
	if user.Id != "" {
		// see if it already exists
		curr, _ := s.DB.GetUser(user.Id)
		if curr != nil {
			return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("User with id '%s' already exists", user.Id))
		}
	} else {
		user.Id = s.DB.NextId("User")
	}
	if user.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name must be specified")
	}

	dbUser := UserFromProto(user)
	err = s.DB.SaveUser(dbUser)
	if err == nil {
		resp = &protos.CreateUserResponse{
			User: UserToProto(dbUser),
		}
	}
	return resp, err
}

// Deletes an user from our system.
func (s *UserService) DeleteUser(ctx context.Context, req *protos.DeleteUserRequest) (resp *protos.DeleteUserResponse, err error) {
	resp = &protos.DeleteUserResponse{}
	s.DB.DeleteUser(req.Id)
	return
}

// Finds and retrieves users matching the particular criteria.
func (s *UserService) ListUsers(ctx context.Context, req *protos.ListUsersRequest) (resp *protos.ListUsersResponse, err error) {
	results, err := s.DB.ListUsers("", 100)
	if err != nil {
		return nil, err
	}
	log.Println("Found Users: ", results)
	resp = &protos.ListUsersResponse{Users: gut.Map(results, UserToProto)}
	return
}

func (s *UserService) GetUser(ctx context.Context, req *protos.GetUserRequest) (resp *protos.GetUserResponse, err error) {
	curr, _ := s.DB.GetUser(req.Id)
	if curr == nil {
		err = status.Error(codes.NotFound, fmt.Sprintf("User with id '%s' not found", req.Id))
	} else {
		resp = &protos.GetUserResponse{User: UserToProto(curr)}
	}
	return
}

func (s *UserService) GetUsers(ctx context.Context, req *protos.GetUsersRequest) (resp *protos.GetUsersResponse, err error) {
	users := gut.BatchGet(req.Ids, func(id string) (out *protos.User, err error) {
		resp, err := s.GetUser(ctx, &protos.GetUserRequest{Id: id})
		if err != nil {
			return nil, err
		}
		return resp.User, nil
	})
	resp = &protos.GetUsersResponse{
		Users: users,
	}
	return
}

// Update a new User
func (s *UserService) UpdateUser(ctx context.Context, req *protos.UpdateUserRequest) (resp *protos.UpdateUserResponse, err error) {
	curruser, err := s.GetUser(ctx, &protos.GetUserRequest{Id: req.User.Id})
	if err != nil {
		return nil, err
	}

	update_mask := req.UpdateMask
	has_update_mask := update_mask != nil && len(update_mask.Paths) > 0
	if !has_update_mask && len(req.AddUsers) == 0 && len(req.RemoveUsers) == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"update_mask should specify (nested) fields to update")
	}

	if req.UpdateMask != nil {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "name":
				curruser.User.Name = req.User.Name
			case "avatar":
				curruser.User.Avatar = req.User.Avatar
			default:
				return nil, status.Errorf(codes.InvalidArgument, "UpdateUser - update_mask contains invalid path: %s", path)
			}
		}
	}

	err = s.DB.SaveUser(UserFromProto(curruser.User))
	if err == nil {
		resp = &protos.UpdateUserResponse{
			User: curruser.User,
		}
	}
	return resp, err
}
