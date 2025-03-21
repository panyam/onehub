// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: onehub/v1/search.proto

package protos

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchTopicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchTopicsRequest) Reset() {
	*x = SearchTopicsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onehub_v1_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTopicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTopicsRequest) ProtoMessage() {}

func (x *SearchTopicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onehub_v1_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTopicsRequest.ProtoReflect.Descriptor instead.
func (*SearchTopicsRequest) Descriptor() ([]byte, []int) {
	return file_onehub_v1_search_proto_rawDescGZIP(), []int{0}
}

var File_onehub_v1_search_proto protoreflect.FileDescriptor

var file_onehub_v1_search_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x32, 0x7a, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x12, 0x1e, 0x2e, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x42, 0x7b,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x65, 0x68, 0x75,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x4f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4f, 0x6e, 0x65,
	0x68, 0x75, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x4f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_onehub_v1_search_proto_rawDescOnce sync.Once
	file_onehub_v1_search_proto_rawDescData = file_onehub_v1_search_proto_rawDesc
)

func file_onehub_v1_search_proto_rawDescGZIP() []byte {
	file_onehub_v1_search_proto_rawDescOnce.Do(func() {
		file_onehub_v1_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_onehub_v1_search_proto_rawDescData)
	})
	return file_onehub_v1_search_proto_rawDescData
}

var file_onehub_v1_search_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_onehub_v1_search_proto_goTypes = []interface{}{
	(*SearchTopicsRequest)(nil), // 0: onehub.v1.SearchTopicsRequest
}
var file_onehub_v1_search_proto_depIdxs = []int32{
	0, // 0: onehub.v1.SearchService.SearchTopics:input_type -> onehub.v1.SearchTopicsRequest
	0, // 1: onehub.v1.SearchService.SearchTopics:output_type -> onehub.v1.SearchTopicsRequest
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_onehub_v1_search_proto_init() }
func file_onehub_v1_search_proto_init() {
	if File_onehub_v1_search_proto != nil {
		return
	}
	file_onehub_v1_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_onehub_v1_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTopicsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_onehub_v1_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_onehub_v1_search_proto_goTypes,
		DependencyIndexes: file_onehub_v1_search_proto_depIdxs,
		MessageInfos:      file_onehub_v1_search_proto_msgTypes,
	}.Build()
	File_onehub_v1_search_proto = out.File
	file_onehub_v1_search_proto_rawDesc = nil
	file_onehub_v1_search_proto_goTypes = nil
	file_onehub_v1_search_proto_depIdxs = nil
}
