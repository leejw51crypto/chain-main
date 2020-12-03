# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cosmos/upgrade/v1beta1/query.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.upgrade.v1beta1 import upgrade_pb2 as cosmos_dot_upgrade_dot_v1beta1_dot_upgrade__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='cosmos/upgrade/v1beta1/query.proto',
  package='cosmos.upgrade.v1beta1',
  syntax='proto3',
  serialized_options=b'Z,github.com/cosmos/cosmos-sdk/x/upgrade/types',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\"cosmos/upgrade/v1beta1/query.proto\x12\x16\x63osmos.upgrade.v1beta1\x1a\x1cgoogle/api/annotations.proto\x1a$cosmos/upgrade/v1beta1/upgrade.proto\"\x19\n\x17QueryCurrentPlanRequest\"F\n\x18QueryCurrentPlanResponse\x12*\n\x04plan\x18\x01 \x01(\x0b\x32\x1c.cosmos.upgrade.v1beta1.Plan\"\'\n\x17QueryAppliedPlanRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"*\n\x18QueryAppliedPlanResponse\x12\x0e\n\x06height\x18\x01 \x01(\x03\x32\xd0\x02\n\x05Query\x12\x9e\x01\n\x0b\x43urrentPlan\x12/.cosmos.upgrade.v1beta1.QueryCurrentPlanRequest\x1a\x30.cosmos.upgrade.v1beta1.QueryCurrentPlanResponse\",\x82\xd3\xe4\x93\x02&\x12$/cosmos/upgrade/v1beta1/current_plan\x12\xa5\x01\n\x0b\x41ppliedPlan\x12/.cosmos.upgrade.v1beta1.QueryAppliedPlanRequest\x1a\x30.cosmos.upgrade.v1beta1.QueryAppliedPlanResponse\"3\x82\xd3\xe4\x93\x02-\x12+/cosmos/upgrade/v1beta1/applied_plan/{name}B.Z,github.com/cosmos/cosmos-sdk/x/upgrade/typesb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,cosmos_dot_upgrade_dot_v1beta1_dot_upgrade__pb2.DESCRIPTOR,])




_QUERYCURRENTPLANREQUEST = _descriptor.Descriptor(
  name='QueryCurrentPlanRequest',
  full_name='cosmos.upgrade.v1beta1.QueryCurrentPlanRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=130,
  serialized_end=155,
)


_QUERYCURRENTPLANRESPONSE = _descriptor.Descriptor(
  name='QueryCurrentPlanResponse',
  full_name='cosmos.upgrade.v1beta1.QueryCurrentPlanResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='plan', full_name='cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=157,
  serialized_end=227,
)


_QUERYAPPLIEDPLANREQUEST = _descriptor.Descriptor(
  name='QueryAppliedPlanRequest',
  full_name='cosmos.upgrade.v1beta1.QueryAppliedPlanRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=229,
  serialized_end=268,
)


_QUERYAPPLIEDPLANRESPONSE = _descriptor.Descriptor(
  name='QueryAppliedPlanResponse',
  full_name='cosmos.upgrade.v1beta1.QueryAppliedPlanResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='height', full_name='cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=270,
  serialized_end=312,
)

_QUERYCURRENTPLANRESPONSE.fields_by_name['plan'].message_type = cosmos_dot_upgrade_dot_v1beta1_dot_upgrade__pb2._PLAN
DESCRIPTOR.message_types_by_name['QueryCurrentPlanRequest'] = _QUERYCURRENTPLANREQUEST
DESCRIPTOR.message_types_by_name['QueryCurrentPlanResponse'] = _QUERYCURRENTPLANRESPONSE
DESCRIPTOR.message_types_by_name['QueryAppliedPlanRequest'] = _QUERYAPPLIEDPLANREQUEST
DESCRIPTOR.message_types_by_name['QueryAppliedPlanResponse'] = _QUERYAPPLIEDPLANRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

QueryCurrentPlanRequest = _reflection.GeneratedProtocolMessageType('QueryCurrentPlanRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYCURRENTPLANREQUEST,
  '__module__' : 'cosmos.upgrade.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.upgrade.v1beta1.QueryCurrentPlanRequest)
  })
_sym_db.RegisterMessage(QueryCurrentPlanRequest)

QueryCurrentPlanResponse = _reflection.GeneratedProtocolMessageType('QueryCurrentPlanResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYCURRENTPLANRESPONSE,
  '__module__' : 'cosmos.upgrade.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.upgrade.v1beta1.QueryCurrentPlanResponse)
  })
_sym_db.RegisterMessage(QueryCurrentPlanResponse)

QueryAppliedPlanRequest = _reflection.GeneratedProtocolMessageType('QueryAppliedPlanRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYAPPLIEDPLANREQUEST,
  '__module__' : 'cosmos.upgrade.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.upgrade.v1beta1.QueryAppliedPlanRequest)
  })
_sym_db.RegisterMessage(QueryAppliedPlanRequest)

QueryAppliedPlanResponse = _reflection.GeneratedProtocolMessageType('QueryAppliedPlanResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYAPPLIEDPLANRESPONSE,
  '__module__' : 'cosmos.upgrade.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.upgrade.v1beta1.QueryAppliedPlanResponse)
  })
_sym_db.RegisterMessage(QueryAppliedPlanResponse)


DESCRIPTOR._options = None

_QUERY = _descriptor.ServiceDescriptor(
  name='Query',
  full_name='cosmos.upgrade.v1beta1.Query',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=315,
  serialized_end=651,
  methods=[
  _descriptor.MethodDescriptor(
    name='CurrentPlan',
    full_name='cosmos.upgrade.v1beta1.Query.CurrentPlan',
    index=0,
    containing_service=None,
    input_type=_QUERYCURRENTPLANREQUEST,
    output_type=_QUERYCURRENTPLANRESPONSE,
    serialized_options=b'\202\323\344\223\002&\022$/cosmos/upgrade/v1beta1/current_plan',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='AppliedPlan',
    full_name='cosmos.upgrade.v1beta1.Query.AppliedPlan',
    index=1,
    containing_service=None,
    input_type=_QUERYAPPLIEDPLANREQUEST,
    output_type=_QUERYAPPLIEDPLANRESPONSE,
    serialized_options=b'\202\323\344\223\002-\022+/cosmos/upgrade/v1beta1/applied_plan/{name}',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_QUERY)

DESCRIPTOR.services_by_name['Query'] = _QUERY

# @@protoc_insertion_point(module_scope)
