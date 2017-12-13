# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account/actions/account.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from account.models import types_pb2 as account_dot_models_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='account/actions/account.proto',
  package='account',
  syntax='proto3',
  serialized_pb=_b('\n\x1d\x61\x63\x63ount/actions/account.proto\x12\x07\x61\x63\x63ount\x1a\x1a\x61\x63\x63ount/models/types.proto\"9\n\x13ValidateUserRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"?\n\x11ValidateUserReply\x12\r\n\x05valid\x18\x01 \x01(\x08\x12\x1b\n\x04user\x18\x02 \x01(\x0b\x32\r.account.UserBP\n\x18\x63loud.ysitd.grpc.accountP\x01Z2github.com/ysitd-cloud/grpc-schema/account/actionsb\x06proto3')
  ,
  dependencies=[account_dot_models_dot_types__pb2.DESCRIPTOR,])




_VALIDATEUSERREQUEST = _descriptor.Descriptor(
  name='ValidateUserRequest',
  full_name='account.ValidateUserRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='username', full_name='account.ValidateUserRequest.username', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='password', full_name='account.ValidateUserRequest.password', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=70,
  serialized_end=127,
)


_VALIDATEUSERREPLY = _descriptor.Descriptor(
  name='ValidateUserReply',
  full_name='account.ValidateUserReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='valid', full_name='account.ValidateUserReply.valid', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='user', full_name='account.ValidateUserReply.user', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=129,
  serialized_end=192,
)

_VALIDATEUSERREPLY.fields_by_name['user'].message_type = account_dot_models_dot_types__pb2._USER
DESCRIPTOR.message_types_by_name['ValidateUserRequest'] = _VALIDATEUSERREQUEST
DESCRIPTOR.message_types_by_name['ValidateUserReply'] = _VALIDATEUSERREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ValidateUserRequest = _reflection.GeneratedProtocolMessageType('ValidateUserRequest', (_message.Message,), dict(
  DESCRIPTOR = _VALIDATEUSERREQUEST,
  __module__ = 'account.actions.account_pb2'
  # @@protoc_insertion_point(class_scope:account.ValidateUserRequest)
  ))
_sym_db.RegisterMessage(ValidateUserRequest)

ValidateUserReply = _reflection.GeneratedProtocolMessageType('ValidateUserReply', (_message.Message,), dict(
  DESCRIPTOR = _VALIDATEUSERREPLY,
  __module__ = 'account.actions.account_pb2'
  # @@protoc_insertion_point(class_scope:account.ValidateUserReply)
  ))
_sym_db.RegisterMessage(ValidateUserReply)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\030cloud.ysitd.grpc.accountP\001Z2github.com/ysitd-cloud/grpc-schema/account/actions'))
# @@protoc_insertion_point(module_scope)
