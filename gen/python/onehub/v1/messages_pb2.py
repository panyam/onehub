# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: onehub/v1/messages.proto
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
    'onehub/v1/messages.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from onehub.v1 import models_pb2 as onehub_dot_v1_dot_models__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18onehub/v1/messages.proto\x12\tonehub.v1\x1a google/protobuf/field_mask.proto\x1a\x16onehub/v1/models.proto\x1a\x1cgoogle/api/annotations.proto\"\x87\x01\n\x15\x43reateMessagesRequest\x12\x19\n\x08topic_id\x18\x01 \x01(\tR\x07topicId\x12.\n\x08messages\x18\x02 \x03(\x0b\x32\x12.onehub.v1.MessageR\x08messages\x12#\n\rallow_userids\x18\x03 \x01(\x08R\x0c\x61llowUserids\"H\n\x16\x43reateMessagesResponse\x12.\n\x08messages\x18\x01 \x03(\x0b\x32\x12.onehub.v1.MessageR\x08messages\"G\n\x15ImportMessagesRequest\x12.\n\x08messages\x18\x02 \x03(\x0b\x32\x12.onehub.v1.MessageR\x08messages\"H\n\x16ImportMessagesResponse\x12.\n\x08messages\x18\x01 \x03(\x0b\x32\x12.onehub.v1.MessageR\x08messages\"g\n\x13ListMessagesRequest\x12\x19\n\x08topic_id\x18\x01 \x01(\tR\x07topicId\x12\x35\n\npagination\x18\x02 \x01(\x0b\x32\x15.onehub.v1.PaginationR\npagination\"\x85\x01\n\x14ListMessagesResponse\x12.\n\x08messages\x18\x01 \x03(\x0b\x32\x12.onehub.v1.MessageR\x08messages\x12=\n\npagination\x18\x02 \x01(\x0b\x32\x1d.onehub.v1.PaginationResponseR\npagination\"\x8f\x01\n\x15SearchMessagesRequest\x12#\n\rsearch_phrase\x18\x01 \x01(\tR\x0csearchPhrase\x12\x1b\n\tsender_id\x18\x02 \x01(\tR\x08senderId\x12\x19\n\x08topic_id\x18\x03 \x01(\tR\x07topicId\x12\x19\n\x08order_by\x18\x04 \x03(\tR\x07orderBy\"\x87\x01\n\x16SearchMessagesResponse\x12.\n\x08messages\x18\x01 \x03(\x0b\x32\x12.onehub.v1.MessageR\x08messages\x12=\n\npagination\x18\x02 \x01(\x0b\x32\x1d.onehub.v1.PaginationResponseR\npagination\"#\n\x11GetMessageRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"B\n\x12GetMessageResponse\x12,\n\x07message\x18\x01 \x01(\x0b\x32\x12.onehub.v1.MessageR\x07message\"&\n\x12GetMessagesRequest\x12\x10\n\x03ids\x18\x01 \x03(\tR\x03ids\"\xb0\x01\n\x13GetMessagesResponse\x12H\n\x08messages\x18\x01 \x03(\x0b\x32,.onehub.v1.GetMessagesResponse.MessagesEntryR\x08messages\x1aO\n\rMessagesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12(\n\x05value\x18\x02 \x01(\x0b\x32\x12.onehub.v1.MessageR\x05value:\x02\x38\x01\"&\n\x14\x44\x65leteMessageRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\x17\n\x15\x44\x65leteMessageResponse\"\xbe\x01\n\x14UpdateMessageRequest\x12,\n\x07message\x18\x01 \x01(\x0b\x32\x12.onehub.v1.MessageR\x07message\x12;\n\x0bupdate_mask\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\nupdateMask\x12;\n\x0b\x61ppend_mask\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\nappendMask\"E\n\x15UpdateMessageResponse\x12,\n\x07message\x18\x01 \x01(\x0b\x32\x12.onehub.v1.MessageR\x07message2\xbb\x07\n\x0eMessageService\x12\x82\x01\n\x0e\x43reateMessages\x12 .onehub.v1.CreateMessagesRequest\x1a!.onehub.v1.CreateMessagesResponse\"+\x82\xd3\xe4\x93\x02%\" /v1/topics/{topic_id=*}/messages:\x01*\x12k\n\x0eSearchMessages\x12 .onehub.v1.SearchMessagesRequest\x1a!.onehub.v1.SearchMessagesResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\x0c/v1/messages\x12u\n\x0eImportMessages\x12 .onehub.v1.ImportMessagesRequest\x1a!.onehub.v1.ImportMessagesResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/messages:import:\x01*\x12y\n\x0cListMessages\x12\x1e.onehub.v1.ListMessagesRequest\x1a\x1f.onehub.v1.ListMessagesResponse\"(\x82\xd3\xe4\x93\x02\"\x12 /v1/topics/{topic_id=*}/messages\x12\x66\n\nGetMessage\x12\x1c.onehub.v1.GetMessageRequest\x1a\x1d.onehub.v1.GetMessageResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/messages/{id=*}\x12k\n\x0bGetMessages\x12\x1d.onehub.v1.GetMessagesRequest\x1a\x1e.onehub.v1.GetMessagesResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/messages:batchGet\x12o\n\rDeleteMessage\x12\x1f.onehub.v1.DeleteMessageRequest\x1a .onehub.v1.DeleteMessageResponse\"\x1b\x82\xd3\xe4\x93\x02\x15*\x13/v1/messages/{id=*}\x12\x7f\n\rUpdateMessage\x12\x1f.onehub.v1.UpdateMessageRequest\x1a .onehub.v1.UpdateMessageResponse\"+\x82\xd3\xe4\x93\x02%2 /v1/messages/{message.base.id=*}:\x01*B}\n\rcom.onehub.v1B\rMessagesProtoP\x01Z\x18github.com/onehub/protos\xa2\x02\x03OXX\xaa\x02\tOnehub.V1\xca\x02\tOnehub\\V1\xe2\x02\x15Onehub\\V1\\GPBMetadata\xea\x02\nOnehub::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'onehub.v1.messages_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\rcom.onehub.v1B\rMessagesProtoP\001Z\030github.com/onehub/protos\242\002\003OXX\252\002\tOnehub.V1\312\002\tOnehub\\V1\342\002\025Onehub\\V1\\GPBMetadata\352\002\nOnehub::V1'
  _globals['_GETMESSAGESRESPONSE_MESSAGESENTRY']._loaded_options = None
  _globals['_GETMESSAGESRESPONSE_MESSAGESENTRY']._serialized_options = b'8\001'
  _globals['_MESSAGESERVICE'].methods_by_name['CreateMessages']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['CreateMessages']._serialized_options = b'\202\323\344\223\002%\" /v1/topics/{topic_id=*}/messages:\001*'
  _globals['_MESSAGESERVICE'].methods_by_name['SearchMessages']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['SearchMessages']._serialized_options = b'\202\323\344\223\002\016\022\014/v1/messages'
  _globals['_MESSAGESERVICE'].methods_by_name['ImportMessages']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['ImportMessages']._serialized_options = b'\202\323\344\223\002\030\"\023/v1/messages:import:\001*'
  _globals['_MESSAGESERVICE'].methods_by_name['ListMessages']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['ListMessages']._serialized_options = b'\202\323\344\223\002\"\022 /v1/topics/{topic_id=*}/messages'
  _globals['_MESSAGESERVICE'].methods_by_name['GetMessage']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['GetMessage']._serialized_options = b'\202\323\344\223\002\025\022\023/v1/messages/{id=*}'
  _globals['_MESSAGESERVICE'].methods_by_name['GetMessages']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['GetMessages']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/messages:batchGet'
  _globals['_MESSAGESERVICE'].methods_by_name['DeleteMessage']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['DeleteMessage']._serialized_options = b'\202\323\344\223\002\025*\023/v1/messages/{id=*}'
  _globals['_MESSAGESERVICE'].methods_by_name['UpdateMessage']._loaded_options = None
  _globals['_MESSAGESERVICE'].methods_by_name['UpdateMessage']._serialized_options = b'\202\323\344\223\002%2 /v1/messages/{message.base.id=*}:\001*'
  _globals['_CREATEMESSAGESREQUEST']._serialized_start=128
  _globals['_CREATEMESSAGESREQUEST']._serialized_end=263
  _globals['_CREATEMESSAGESRESPONSE']._serialized_start=265
  _globals['_CREATEMESSAGESRESPONSE']._serialized_end=337
  _globals['_IMPORTMESSAGESREQUEST']._serialized_start=339
  _globals['_IMPORTMESSAGESREQUEST']._serialized_end=410
  _globals['_IMPORTMESSAGESRESPONSE']._serialized_start=412
  _globals['_IMPORTMESSAGESRESPONSE']._serialized_end=484
  _globals['_LISTMESSAGESREQUEST']._serialized_start=486
  _globals['_LISTMESSAGESREQUEST']._serialized_end=589
  _globals['_LISTMESSAGESRESPONSE']._serialized_start=592
  _globals['_LISTMESSAGESRESPONSE']._serialized_end=725
  _globals['_SEARCHMESSAGESREQUEST']._serialized_start=728
  _globals['_SEARCHMESSAGESREQUEST']._serialized_end=871
  _globals['_SEARCHMESSAGESRESPONSE']._serialized_start=874
  _globals['_SEARCHMESSAGESRESPONSE']._serialized_end=1009
  _globals['_GETMESSAGEREQUEST']._serialized_start=1011
  _globals['_GETMESSAGEREQUEST']._serialized_end=1046
  _globals['_GETMESSAGERESPONSE']._serialized_start=1048
  _globals['_GETMESSAGERESPONSE']._serialized_end=1114
  _globals['_GETMESSAGESREQUEST']._serialized_start=1116
  _globals['_GETMESSAGESREQUEST']._serialized_end=1154
  _globals['_GETMESSAGESRESPONSE']._serialized_start=1157
  _globals['_GETMESSAGESRESPONSE']._serialized_end=1333
  _globals['_GETMESSAGESRESPONSE_MESSAGESENTRY']._serialized_start=1254
  _globals['_GETMESSAGESRESPONSE_MESSAGESENTRY']._serialized_end=1333
  _globals['_DELETEMESSAGEREQUEST']._serialized_start=1335
  _globals['_DELETEMESSAGEREQUEST']._serialized_end=1373
  _globals['_DELETEMESSAGERESPONSE']._serialized_start=1375
  _globals['_DELETEMESSAGERESPONSE']._serialized_end=1398
  _globals['_UPDATEMESSAGEREQUEST']._serialized_start=1401
  _globals['_UPDATEMESSAGEREQUEST']._serialized_end=1591
  _globals['_UPDATEMESSAGERESPONSE']._serialized_start=1593
  _globals['_UPDATEMESSAGERESPONSE']._serialized_end=1662
  _globals['_MESSAGESERVICE']._serialized_start=1665
  _globals['_MESSAGESERVICE']._serialized_end=2620
# @@protoc_insertion_point(module_scope)
