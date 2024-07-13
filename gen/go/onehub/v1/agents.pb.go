// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: onehub/v1/agents.proto

package protos

import (
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

type Tool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique ID for this tool
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A readable name for this tool - as understood by an LLM
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of this tool as understood by an LLM
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Info about input parameters
	InputParams []*ParamInfo `protobuf:"bytes,4,rep,name=input_params,json=inputParams,proto3" json:"input_params,omitempty"`
	// Info about output params
	OutputParams []*ParamInfo `protobuf:"bytes,5,rep,name=output_params,json=outputParams,proto3" json:"output_params,omitempty"`
}

func (x *Tool) Reset() {
	*x = Tool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onehub_v1_agents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tool) ProtoMessage() {}

func (x *Tool) ProtoReflect() protoreflect.Message {
	mi := &file_onehub_v1_agents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tool.ProtoReflect.Descriptor instead.
func (*Tool) Descriptor() ([]byte, []int) {
	return file_onehub_v1_agents_proto_rawDescGZIP(), []int{0}
}

func (x *Tool) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tool) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tool) GetInputParams() []*ParamInfo {
	if x != nil {
		return x.InputParams
	}
	return nil
}

func (x *Tool) GetOutputParams() []*ParamInfo {
	if x != nil {
		return x.OutputParams
	}
	return nil
}

type ParamInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the parameter
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the parameter.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// A string notation of the type
	ParamType string `protobuf:"bytes,3,opt,name=param_type,json=paramType,proto3" json:"param_type,omitempty"`
}

func (x *ParamInfo) Reset() {
	*x = ParamInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onehub_v1_agents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamInfo) ProtoMessage() {}

func (x *ParamInfo) ProtoReflect() protoreflect.Message {
	mi := &file_onehub_v1_agents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamInfo.ProtoReflect.Descriptor instead.
func (*ParamInfo) Descriptor() ([]byte, []int) {
	return file_onehub_v1_agents_proto_rawDescGZIP(), []int{1}
}

func (x *ParamInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ParamInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ParamInfo) GetParamType() string {
	if x != nil {
		return x.ParamType
	}
	return ""
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class       string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onehub_v1_agents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_onehub_v1_agents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_onehub_v1_agents_proto_rawDescGZIP(), []int{2}
}

func (x *Agent) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Agent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Agent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_onehub_v1_agents_proto protoreflect.FileDescriptor

var file_onehub_v1_agents_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x04,
	0x54, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0c, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x6e, 0x65,
	0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x60,
	0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x53, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x13, 0x0a, 0x11, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x0f, 0x0a, 0x0d, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x7b, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4f, 0x6e,
	0x65, 0x68, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4f, 0x6e, 0x65, 0x68, 0x75, 0x62,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4f, 0x6e, 0x65, 0x68, 0x75, 0x62, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4f, 0x6e,
	0x65, 0x68, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_onehub_v1_agents_proto_rawDescOnce sync.Once
	file_onehub_v1_agents_proto_rawDescData = file_onehub_v1_agents_proto_rawDesc
)

func file_onehub_v1_agents_proto_rawDescGZIP() []byte {
	file_onehub_v1_agents_proto_rawDescOnce.Do(func() {
		file_onehub_v1_agents_proto_rawDescData = protoimpl.X.CompressGZIP(file_onehub_v1_agents_proto_rawDescData)
	})
	return file_onehub_v1_agents_proto_rawDescData
}

var file_onehub_v1_agents_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_onehub_v1_agents_proto_goTypes = []interface{}{
	(*Tool)(nil),      // 0: onehub.v1.Tool
	(*ParamInfo)(nil), // 1: onehub.v1.ParamInfo
	(*Agent)(nil),     // 2: onehub.v1.Agent
}
var file_onehub_v1_agents_proto_depIdxs = []int32{
	1, // 0: onehub.v1.Tool.input_params:type_name -> onehub.v1.ParamInfo
	1, // 1: onehub.v1.Tool.output_params:type_name -> onehub.v1.ParamInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_onehub_v1_agents_proto_init() }
func file_onehub_v1_agents_proto_init() {
	if File_onehub_v1_agents_proto != nil {
		return
	}
	file_onehub_v1_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_onehub_v1_agents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tool); i {
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
		file_onehub_v1_agents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamInfo); i {
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
		file_onehub_v1_agents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
			RawDescriptor: file_onehub_v1_agents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_onehub_v1_agents_proto_goTypes,
		DependencyIndexes: file_onehub_v1_agents_proto_depIdxs,
		MessageInfos:      file_onehub_v1_agents_proto_msgTypes,
	}.Build()
	File_onehub_v1_agents_proto = out.File
	file_onehub_v1_agents_proto_rawDesc = nil
	file_onehub_v1_agents_proto_goTypes = nil
	file_onehub_v1_agents_proto_depIdxs = nil
}
