// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: onehub/v1/topics.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TopicService_CreateTopic_FullMethodName = "/onehub.v1.TopicService/CreateTopic"
	TopicService_ListTopics_FullMethodName  = "/onehub.v1.TopicService/ListTopics"
	TopicService_GetTopic_FullMethodName    = "/onehub.v1.TopicService/GetTopic"
	TopicService_GetTopics_FullMethodName   = "/onehub.v1.TopicService/GetTopics"
	TopicService_DeleteTopic_FullMethodName = "/onehub.v1.TopicService/DeleteTopic"
	TopicService_UpdateTopic_FullMethodName = "/onehub.v1.TopicService/UpdateTopic"
)

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	// *
	// Create a new sesssion
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error)
	// *
	// List all topics from a user.
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	// *
	// Get a particular topic
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error)
	// *
	// Batch get multiple topics by ID
	GetTopics(ctx context.Context, in *GetTopicsRequest, opts ...grpc.CallOption) (*GetTopicsResponse, error)
	// *
	// Delete a particular topic
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error)
	// *
	// Updates specific fields of a topic
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*UpdateTopicResponse, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_CreateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, TopicService_ListTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error) {
	out := new(GetTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_GetTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopics(ctx context.Context, in *GetTopicsRequest, opts ...grpc.CallOption) (*GetTopicsResponse, error) {
	out := new(GetTopicsResponse)
	err := c.cc.Invoke(ctx, TopicService_GetTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error) {
	out := new(DeleteTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_DeleteTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*UpdateTopicResponse, error) {
	out := new(UpdateTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_UpdateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations should embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	// *
	// Create a new sesssion
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
	// *
	// List all topics from a user.
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	// *
	// Get a particular topic
	GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error)
	// *
	// Batch get multiple topics by ID
	GetTopics(context.Context, *GetTopicsRequest) (*GetTopicsResponse, error)
	// *
	// Delete a particular topic
	DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error)
	// *
	// Updates specific fields of a topic
	UpdateTopic(context.Context, *UpdateTopicRequest) (*UpdateTopicResponse, error)
}

// UnimplementedTopicServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedTopicServiceServer) ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedTopicServiceServer) GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedTopicServiceServer) GetTopics(context.Context, *GetTopicsRequest) (*GetTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (UnimplementedTopicServiceServer) DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedTopicServiceServer) UpdateTopic(context.Context, *UpdateTopicRequest) (*UpdateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_ListTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopics(ctx, req.(*GetTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_UpdateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "onehub.v1.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _TopicService_ListTopics_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicService_GetTopic_Handler,
		},
		{
			MethodName: "GetTopics",
			Handler:    _TopicService_GetTopics_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _TopicService_UpdateTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onehub/v1/topics.proto",
}
