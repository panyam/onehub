# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: onehub/v1/agents.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'onehub/v1/agents.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from onehub.v1 import models_pb2 as onehub_dot_v1_dot_models__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16onehub/v1/agents.proto\x12\tonehub.v1\x1a google/protobuf/field_mask.proto\x1a\x16onehub/v1/models.proto\x1a\x1cgoogle/api/annotations.proto\"\xc0\x01\n\x04Tool\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\x12\x37\n\x0cinput_params\x18\x04 \x03(\x0b\x32\x14.onehub.v1.ParamInfoR\x0binputParams\x12\x39\n\routput_params\x18\x05 \x03(\x0b\x32\x14.onehub.v1.ParamInfoR\x0coutputParams\"`\n\tParamInfo\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x02 \x01(\tR\x0b\x64\x65scription\x12\x1d\n\nparam_type\x18\x03 \x01(\tR\tparamType\"S\n\x05\x41gent\x12\x14\n\x05\x63lass\x18\x01 \x01(\tR\x05\x63lass\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription2\x13\n\x11\x41gentPlaneService2\x0f\n\rAgentsServiceB{\n\rcom.onehub.v1B\x0b\x41gentsProtoP\x01Z\x18github.com/onehub/protos\xa2\x02\x03OXX\xaa\x02\tOnehub.V1\xca\x02\tOnehub\\V1\xe2\x02\x15Onehub\\V1\\GPBMetadata\xea\x02\nOnehub::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'onehub.v1.agents_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\rcom.onehub.v1B\013AgentsProtoP\001Z\030github.com/onehub/protos\242\002\003OXX\252\002\tOnehub.V1\312\002\tOnehub\\V1\342\002\025Onehub\\V1\\GPBMetadata\352\002\nOnehub::V1'
  _globals['_TOOL']._serialized_start=126
  _globals['_TOOL']._serialized_end=318
  _globals['_PARAMINFO']._serialized_start=320
  _globals['_PARAMINFO']._serialized_end=416
  _globals['_AGENT']._serialized_start=418
  _globals['_AGENT']._serialized_end=501
  _globals['_AGENTPLANESERVICE']._serialized_start=503
  _globals['_AGENTPLANESERVICE']._serialized_end=522
  _globals['_AGENTSSERVICE']._serialized_start=524
  _globals['_AGENTSSERVICE']._serialized_end=539
# @@protoc_insertion_point(module_scope)