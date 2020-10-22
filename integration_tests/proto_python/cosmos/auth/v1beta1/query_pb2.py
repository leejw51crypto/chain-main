# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cosmos/auth/v1beta1/query.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.auth.v1beta1 import auth_pb2 as cosmos_dot_auth_dot_v1beta1_dot_auth__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='cosmos/auth/v1beta1/query.proto',
  package='cosmos.auth.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z)github.com/cosmos/cosmos-sdk/x/auth/types'),
  serialized_pb=_b('\n\x1f\x63osmos/auth/v1beta1/query.proto\x12\x13\x63osmos.auth.v1beta1\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1e\x63osmos/auth/v1beta1/auth.proto\x1a\x19\x63osmos_proto/cosmos.proto\"0\n\x13QueryAccountRequest\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t:\x08\xe8\xa0\x1f\x00\x88\xa0\x1f\x00\"K\n\x14QueryAccountResponse\x12\x33\n\x07\x61\x63\x63ount\x18\x01 \x01(\x0b\x32\x14.google.protobuf.AnyB\x0c\xca\xb4-\x08\x41\x63\x63ountI\"\x14\n\x12QueryParamsRequest\"H\n\x13QueryParamsResponse\x12\x31\n\x06params\x18\x01 \x01(\x0b\x32\x1b.cosmos.auth.v1beta1.ParamsB\x04\xc8\xde\x1f\x00\x32\x9c\x02\n\x05Query\x12\x8f\x01\n\x07\x41\x63\x63ount\x12(.cosmos.auth.v1beta1.QueryAccountRequest\x1a).cosmos.auth.v1beta1.QueryAccountResponse\"/\x82\xd3\xe4\x93\x02)\x12\'/cosmos/auth/v1beta1/accounts/{address}\x12\x80\x01\n\x06Params\x12\'.cosmos.auth.v1beta1.QueryParamsRequest\x1a(.cosmos.auth.v1beta1.QueryParamsResponse\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/cosmos/auth/v1beta1/paramsB+Z)github.com/cosmos/cosmos-sdk/x/auth/typesb\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,google_dot_protobuf_dot_any__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,cosmos_dot_auth_dot_v1beta1_dot_auth__pb2.DESCRIPTOR,cosmos__proto_dot_cosmos__pb2.DESCRIPTOR,])




_QUERYACCOUNTREQUEST = _descriptor.Descriptor(
  name='QueryAccountRequest',
  full_name='cosmos.auth.v1beta1.QueryAccountRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='cosmos.auth.v1beta1.QueryAccountRequest.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('\350\240\037\000\210\240\037\000'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=194,
  serialized_end=242,
)


_QUERYACCOUNTRESPONSE = _descriptor.Descriptor(
  name='QueryAccountResponse',
  full_name='cosmos.auth.v1beta1.QueryAccountResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account', full_name='cosmos.auth.v1beta1.QueryAccountResponse.account', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\312\264-\010AccountI'), file=DESCRIPTOR),
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
  serialized_start=244,
  serialized_end=319,
)


_QUERYPARAMSREQUEST = _descriptor.Descriptor(
  name='QueryParamsRequest',
  full_name='cosmos.auth.v1beta1.QueryParamsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
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
  serialized_start=321,
  serialized_end=341,
)


_QUERYPARAMSRESPONSE = _descriptor.Descriptor(
  name='QueryParamsResponse',
  full_name='cosmos.auth.v1beta1.QueryParamsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='params', full_name='cosmos.auth.v1beta1.QueryParamsResponse.params', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000'), file=DESCRIPTOR),
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
  serialized_start=343,
  serialized_end=415,
)

_QUERYACCOUNTRESPONSE.fields_by_name['account'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_QUERYPARAMSRESPONSE.fields_by_name['params'].message_type = cosmos_dot_auth_dot_v1beta1_dot_auth__pb2._PARAMS
DESCRIPTOR.message_types_by_name['QueryAccountRequest'] = _QUERYACCOUNTREQUEST
DESCRIPTOR.message_types_by_name['QueryAccountResponse'] = _QUERYACCOUNTRESPONSE
DESCRIPTOR.message_types_by_name['QueryParamsRequest'] = _QUERYPARAMSREQUEST
DESCRIPTOR.message_types_by_name['QueryParamsResponse'] = _QUERYPARAMSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

QueryAccountRequest = _reflection.GeneratedProtocolMessageType('QueryAccountRequest', (_message.Message,), dict(
  DESCRIPTOR = _QUERYACCOUNTREQUEST,
  __module__ = 'cosmos.auth.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.auth.v1beta1.QueryAccountRequest)
  ))
_sym_db.RegisterMessage(QueryAccountRequest)

QueryAccountResponse = _reflection.GeneratedProtocolMessageType('QueryAccountResponse', (_message.Message,), dict(
  DESCRIPTOR = _QUERYACCOUNTRESPONSE,
  __module__ = 'cosmos.auth.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.auth.v1beta1.QueryAccountResponse)
  ))
_sym_db.RegisterMessage(QueryAccountResponse)

QueryParamsRequest = _reflection.GeneratedProtocolMessageType('QueryParamsRequest', (_message.Message,), dict(
  DESCRIPTOR = _QUERYPARAMSREQUEST,
  __module__ = 'cosmos.auth.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.auth.v1beta1.QueryParamsRequest)
  ))
_sym_db.RegisterMessage(QueryParamsRequest)

QueryParamsResponse = _reflection.GeneratedProtocolMessageType('QueryParamsResponse', (_message.Message,), dict(
  DESCRIPTOR = _QUERYPARAMSRESPONSE,
  __module__ = 'cosmos.auth.v1beta1.query_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.auth.v1beta1.QueryParamsResponse)
  ))
_sym_db.RegisterMessage(QueryParamsResponse)


DESCRIPTOR._options = None
_QUERYACCOUNTREQUEST._options = None
_QUERYACCOUNTRESPONSE.fields_by_name['account']._options = None
_QUERYPARAMSRESPONSE.fields_by_name['params']._options = None

_QUERY = _descriptor.ServiceDescriptor(
  name='Query',
  full_name='cosmos.auth.v1beta1.Query',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=418,
  serialized_end=702,
  methods=[
  _descriptor.MethodDescriptor(
    name='Account',
    full_name='cosmos.auth.v1beta1.Query.Account',
    index=0,
    containing_service=None,
    input_type=_QUERYACCOUNTREQUEST,
    output_type=_QUERYACCOUNTRESPONSE,
    serialized_options=_b('\202\323\344\223\002)\022\'/cosmos/auth/v1beta1/accounts/{address}'),
  ),
  _descriptor.MethodDescriptor(
    name='Params',
    full_name='cosmos.auth.v1beta1.Query.Params',
    index=1,
    containing_service=None,
    input_type=_QUERYPARAMSREQUEST,
    output_type=_QUERYPARAMSRESPONSE,
    serialized_options=_b('\202\323\344\223\002\035\022\033/cosmos/auth/v1beta1/params'),
  ),
])
_sym_db.RegisterServiceDescriptor(_QUERY)

DESCRIPTOR.services_by_name['Query'] = _QUERY

# @@protoc_insertion_point(module_scope)
