# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: onehub/v1/users.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15onehub/v1/users.proto\x12\tonehub.v1\x1a google/protobuf/field_mask.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x16onehub/v1/models.proto\x1a\x1cgoogle/api/annotations.proto\"8\n\x11\x43reateUserRequest\x12#\n\x04user\x18\x01 \x01(\x0b\x32\x0f.onehub.v1.UserR\x04user\"9\n\x12\x43reateUserResponse\x12#\n\x04user\x18\x01 \x01(\x0b\x32\x0f.onehub.v1.UserR\x04user\"J\n\x10ListUsersRequest\x12\x19\n\x08page_key\x18\x01 \x01(\tR\x07pageKey\x12\x1b\n\tpage_size\x18\x02 \x01(\x05R\x08pageSize\"^\n\x11ListUsersResponse\x12%\n\x05users\x18\x01 \x03(\x0b\x32\x0f.onehub.v1.UserR\x05users\x12\"\n\rnext_page_key\x18\x02 \x01(\tR\x0bnextPageKey\" \n\x0eGetUserRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"6\n\x0fGetUserResponse\x12#\n\x04user\x18\x01 \x01(\x0b\x32\x0f.onehub.v1.UserR\x04user\"#\n\x0fGetUsersRequest\x12\x10\n\x03ids\x18\x01 \x03(\tR\x03ids\"\x9b\x01\n\x10GetUsersResponse\x12<\n\x05users\x18\x01 \x03(\x0b\x32&.onehub.v1.GetUsersResponse.UsersEntryR\x05users\x1aI\n\nUsersEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12%\n\x05value\x18\x02 \x01(\x0b\x32\x0f.onehub.v1.UserR\x05value:\x02\x38\x01\"#\n\x11\x44\x65leteUserRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\x14\n\x12\x44\x65leteUserResponse\"\xcf\x01\n\x11UpdateUserRequest\x12#\n\x04user\x18\x01 \x01(\x0b\x32\x0f.onehub.v1.UserR\x04user\x12;\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\nupdateMask\x12\x1b\n\tadd_users\x18\x03 \x03(\tR\x08\x61\x64\x64Users\x12!\n\x0cremove_users\x18\x04 \x03(\tR\x0bremoveUsers:\x18\x92\x41\x15\n\x13*\x11UpdateUserRequest\"T\n\x12UpdateUserResponse\x12#\n\x04user\x18\x01 \x01(\x0b\x32\x0f.onehub.v1.UserR\x04user:\x19\x92\x41\x16\n\x14*\x12UpdateUserResponse2\xd8\x04\n\x0bUserService\x12_\n\nCreateUser\x12\x1c.onehub.v1.CreateUserRequest\x1a\x1d.onehub.v1.CreateUserResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\"\t/v1/users:\x01*\x12Y\n\tListUsers\x12\x1b.onehub.v1.ListUsersRequest\x1a\x1c.onehub.v1.ListUsersResponse\"\x11\x82\xd3\xe4\x93\x02\x0b\x12\t/v1/users\x12Z\n\x07GetUser\x12\x19.onehub.v1.GetUserRequest\x1a\x1a.onehub.v1.GetUserResponse\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/v1/users/{id=*}\x12_\n\x08GetUsers\x12\x1a.onehub.v1.GetUsersRequest\x1a\x1b.onehub.v1.GetUsersResponse\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/v1/users:batchGet\x12\x63\n\nDeleteUser\x12\x1c.onehub.v1.DeleteUserRequest\x1a\x1d.onehub.v1.DeleteUserResponse\"\x18\x82\xd3\xe4\x93\x02\x12*\x10/v1/users/{id=*}\x12k\n\nUpdateUser\x12\x1c.onehub.v1.UpdateUserRequest\x1a\x1d.onehub.v1.UpdateUserResponse\" \x82\xd3\xe4\x93\x02\x1a\x32\x15/v1/users/{user.id=*}:\x01*Bz\n\rcom.onehub.v1B\nUsersProtoP\x01Z\x18github.com/onehub/protos\xa2\x02\x03OXX\xaa\x02\tOnehub.V1\xca\x02\tOnehub\\V1\xe2\x02\x15Onehub\\V1\\GPBMetadata\xea\x02\nOnehub::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'onehub.v1.users_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\rcom.onehub.v1B\nUsersProtoP\001Z\030github.com/onehub/protos\242\002\003OXX\252\002\tOnehub.V1\312\002\tOnehub\\V1\342\002\025Onehub\\V1\\GPBMetadata\352\002\nOnehub::V1'
  _GETUSERSRESPONSE_USERSENTRY._options = None
  _GETUSERSRESPONSE_USERSENTRY._serialized_options = b'8\001'
  _UPDATEUSERREQUEST._options = None
  _UPDATEUSERREQUEST._serialized_options = b'\222A\025\n\023*\021UpdateUserRequest'
  _UPDATEUSERRESPONSE._options = None
  _UPDATEUSERRESPONSE._serialized_options = b'\222A\026\n\024*\022UpdateUserResponse'
  _USERSERVICE.methods_by_name['CreateUser']._options = None
  _USERSERVICE.methods_by_name['CreateUser']._serialized_options = b'\202\323\344\223\002\016\"\t/v1/users:\001*'
  _USERSERVICE.methods_by_name['ListUsers']._options = None
  _USERSERVICE.methods_by_name['ListUsers']._serialized_options = b'\202\323\344\223\002\013\022\t/v1/users'
  _USERSERVICE.methods_by_name['GetUser']._options = None
  _USERSERVICE.methods_by_name['GetUser']._serialized_options = b'\202\323\344\223\002\022\022\020/v1/users/{id=*}'
  _USERSERVICE.methods_by_name['GetUsers']._options = None
  _USERSERVICE.methods_by_name['GetUsers']._serialized_options = b'\202\323\344\223\002\024\022\022/v1/users:batchGet'
  _USERSERVICE.methods_by_name['DeleteUser']._options = None
  _USERSERVICE.methods_by_name['DeleteUser']._serialized_options = b'\202\323\344\223\002\022*\020/v1/users/{id=*}'
  _USERSERVICE.methods_by_name['UpdateUser']._options = None
  _USERSERVICE.methods_by_name['UpdateUser']._serialized_options = b'\202\323\344\223\002\0322\025/v1/users/{user.id=*}:\001*'
  _globals['_CREATEUSERREQUEST']._serialized_start=172
  _globals['_CREATEUSERREQUEST']._serialized_end=228
  _globals['_CREATEUSERRESPONSE']._serialized_start=230
  _globals['_CREATEUSERRESPONSE']._serialized_end=287
  _globals['_LISTUSERSREQUEST']._serialized_start=289
  _globals['_LISTUSERSREQUEST']._serialized_end=363
  _globals['_LISTUSERSRESPONSE']._serialized_start=365
  _globals['_LISTUSERSRESPONSE']._serialized_end=459
  _globals['_GETUSERREQUEST']._serialized_start=461
  _globals['_GETUSERREQUEST']._serialized_end=493
  _globals['_GETUSERRESPONSE']._serialized_start=495
  _globals['_GETUSERRESPONSE']._serialized_end=549
  _globals['_GETUSERSREQUEST']._serialized_start=551
  _globals['_GETUSERSREQUEST']._serialized_end=586
  _globals['_GETUSERSRESPONSE']._serialized_start=589
  _globals['_GETUSERSRESPONSE']._serialized_end=744
  _globals['_GETUSERSRESPONSE_USERSENTRY']._serialized_start=671
  _globals['_GETUSERSRESPONSE_USERSENTRY']._serialized_end=744
  _globals['_DELETEUSERREQUEST']._serialized_start=746
  _globals['_DELETEUSERREQUEST']._serialized_end=781
  _globals['_DELETEUSERRESPONSE']._serialized_start=783
  _globals['_DELETEUSERRESPONSE']._serialized_end=803
  _globals['_UPDATEUSERREQUEST']._serialized_start=806
  _globals['_UPDATEUSERREQUEST']._serialized_end=1013
  _globals['_UPDATEUSERRESPONSE']._serialized_start=1015
  _globals['_UPDATEUSERRESPONSE']._serialized_end=1099
  _globals['_USERSERVICE']._serialized_start=1102
  _globals['_USERSERVICE']._serialized_end=1702
# @@protoc_insertion_point(module_scope)