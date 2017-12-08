# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: judge/actions/judge.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='judge/actions/judge.proto',
  package='judge',
  syntax='proto3',
  serialized_pb=_b('\n\x19judge/actions/judge.proto\x12\x05judge\"Q\n\rAccessRequest\x12\x0f\n\x07subject\x18\x01 \x01(\t\x12\x0e\n\x06\x61\x63tion\x18\x02 \x01(\t\x12\x10\n\x08resource\x18\x03 \x01(\t\x12\r\n\x05token\x18\x04 \x01(\t\",\n\x0b\x41\x63\x63\x65ssReply\x12\r\n\x05\x61llow\x18\x01 \x01(\x08\x12\x0e\n\x06reason\x18\x02 \x01(\tBL\n\x16\x63loud.ysitd.grpc.judgeP\x01Z0github.com/ysitd-cloud/grpc-schema/judge/actionsb\x06proto3')
)




_ACCESSREQUEST = _descriptor.Descriptor(
  name='AccessRequest',
  full_name='judge.AccessRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='subject', full_name='judge.AccessRequest.subject', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='action', full_name='judge.AccessRequest.action', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='resource', full_name='judge.AccessRequest.resource', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='token', full_name='judge.AccessRequest.token', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=36,
  serialized_end=117,
)


_ACCESSREPLY = _descriptor.Descriptor(
  name='AccessReply',
  full_name='judge.AccessReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='allow', full_name='judge.AccessReply.allow', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='reason', full_name='judge.AccessReply.reason', index=1,
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
  serialized_start=119,
  serialized_end=163,
)

DESCRIPTOR.message_types_by_name['AccessRequest'] = _ACCESSREQUEST
DESCRIPTOR.message_types_by_name['AccessReply'] = _ACCESSREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AccessRequest = _reflection.GeneratedProtocolMessageType('AccessRequest', (_message.Message,), dict(
  DESCRIPTOR = _ACCESSREQUEST,
  __module__ = 'judge.actions.judge_pb2'
  # @@protoc_insertion_point(class_scope:judge.AccessRequest)
  ))
_sym_db.RegisterMessage(AccessRequest)

AccessReply = _reflection.GeneratedProtocolMessageType('AccessReply', (_message.Message,), dict(
  DESCRIPTOR = _ACCESSREPLY,
  __module__ = 'judge.actions.judge_pb2'
  # @@protoc_insertion_point(class_scope:judge.AccessReply)
  ))
_sym_db.RegisterMessage(AccessReply)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\026cloud.ysitd.grpc.judgeP\001Z0github.com/ysitd-cloud/grpc-schema/judge/actions'))
# @@protoc_insertion_point(module_scope)