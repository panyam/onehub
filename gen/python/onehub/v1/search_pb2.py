# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: onehub/v1/search.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from protoc_gen_openapiv2.options import annotations_pb2 as protoc__gen__openapiv2_dot_options_dot_annotations__pb2
from onehub.v1 import models_pb2 as onehub_dot_v1_dot_models__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16onehub/v1/search.proto\x12\tonehub.v1\x1a google/protobuf/field_mask.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x16onehub/v1/models.proto\x1a\x1cgoogle/api/annotations.proto\"\x15\n\x13SearchTopicsRequest2z\n\rSearchService\x12i\n\x0cSearchTopics\x12\x1e.onehub.v1.SearchTopicsRequest\x1a\x1e.onehub.v1.SearchTopicsRequest\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/search/topicsB{\n\rcom.onehub.v1B\x0bSearchProtoP\x01Z\x18github.com/onehub/protos\xa2\x02\x03OXX\xaa\x02\tOnehub.V1\xca\x02\tOnehub\\V1\xe2\x02\x15Onehub\\V1\\GPBMetadata\xea\x02\nOnehub::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'onehub.v1.search_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\rcom.onehub.v1B\013SearchProtoP\001Z\030github.com/onehub/protos\242\002\003OXX\252\002\tOnehub.V1\312\002\tOnehub\\V1\342\002\025Onehub\\V1\\GPBMetadata\352\002\nOnehub::V1'
  _SEARCHSERVICE.methods_by_name['SearchTopics']._options = None
  _SEARCHSERVICE.methods_by_name['SearchTopics']._serialized_options = b'\202\323\344\223\002\023\022\021/v1/search/topics'
  _globals['_SEARCHTOPICSREQUEST']._serialized_start=173
  _globals['_SEARCHTOPICSREQUEST']._serialized_end=194
  _globals['_SEARCHSERVICE']._serialized_start=196
  _globals['_SEARCHSERVICE']._serialized_end=318
# @@protoc_insertion_point(module_scope)
