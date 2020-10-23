# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cosmos/base/simulate/v1beta1/simulate.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.base.abci.v1beta1 import abci_pb2 as cosmos_dot_base_dot_abci_dot_v1beta1_dot_abci__pb2
from cosmos.tx.v1beta1 import tx_pb2 as cosmos_dot_tx_dot_v1beta1_dot_tx__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='cosmos/base/simulate/v1beta1/simulate.proto',
  package='cosmos.base.simulate.v1beta1',
  syntax='proto3',
  serialized_options=b'Z1github.com/cosmos/cosmos-sdk/client/grpc/simulate',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n+cosmos/base/simulate/v1beta1/simulate.proto\x12\x1c\x63osmos.base.simulate.v1beta1\x1a\x1cgoogle/api/annotations.proto\x1a#cosmos/base/abci/v1beta1/abci.proto\x1a\x1a\x63osmos/tx/v1beta1/tx.proto\"4\n\x0fSimulateRequest\x12!\n\x02tx\x18\x01 \x01(\x0b\x32\x15.cosmos.tx.v1beta1.Tx\"y\n\x10SimulateResponse\x12\x33\n\x08gas_info\x18\x01 \x01(\x0b\x32!.cosmos.base.abci.v1beta1.GasInfo\x12\x30\n\x06result\x18\x02 \x01(\x0b\x32 .cosmos.base.abci.v1beta1.Result2\xad\x01\n\x0fSimulateService\x12\x99\x01\n\x08Simulate\x12-.cosmos.base.simulate.v1beta1.SimulateRequest\x1a..cosmos.base.simulate.v1beta1.SimulateResponse\".\x82\xd3\xe4\x93\x02(\"&/cosmos/base/simulate/v1beta1/simulateB3Z1github.com/cosmos/cosmos-sdk/client/grpc/simulateb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,cosmos_dot_base_dot_abci_dot_v1beta1_dot_abci__pb2.DESCRIPTOR,cosmos_dot_tx_dot_v1beta1_dot_tx__pb2.DESCRIPTOR,])




_SIMULATEREQUEST = _descriptor.Descriptor(
  name='SimulateRequest',
  full_name='cosmos.base.simulate.v1beta1.SimulateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='tx', full_name='cosmos.base.simulate.v1beta1.SimulateRequest.tx', index=0,
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
  serialized_start=172,
  serialized_end=224,
)


_SIMULATERESPONSE = _descriptor.Descriptor(
  name='SimulateResponse',
  full_name='cosmos.base.simulate.v1beta1.SimulateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='gas_info', full_name='cosmos.base.simulate.v1beta1.SimulateResponse.gas_info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result', full_name='cosmos.base.simulate.v1beta1.SimulateResponse.result', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=226,
  serialized_end=347,
)

_SIMULATEREQUEST.fields_by_name['tx'].message_type = cosmos_dot_tx_dot_v1beta1_dot_tx__pb2._TX
_SIMULATERESPONSE.fields_by_name['gas_info'].message_type = cosmos_dot_base_dot_abci_dot_v1beta1_dot_abci__pb2._GASINFO
_SIMULATERESPONSE.fields_by_name['result'].message_type = cosmos_dot_base_dot_abci_dot_v1beta1_dot_abci__pb2._RESULT
DESCRIPTOR.message_types_by_name['SimulateRequest'] = _SIMULATEREQUEST
DESCRIPTOR.message_types_by_name['SimulateResponse'] = _SIMULATERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SimulateRequest = _reflection.GeneratedProtocolMessageType('SimulateRequest', (_message.Message,), {
  'DESCRIPTOR' : _SIMULATEREQUEST,
  '__module__' : 'cosmos.base.simulate.v1beta1.simulate_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.base.simulate.v1beta1.SimulateRequest)
  })
_sym_db.RegisterMessage(SimulateRequest)

SimulateResponse = _reflection.GeneratedProtocolMessageType('SimulateResponse', (_message.Message,), {
  'DESCRIPTOR' : _SIMULATERESPONSE,
  '__module__' : 'cosmos.base.simulate.v1beta1.simulate_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.base.simulate.v1beta1.SimulateResponse)
  })
_sym_db.RegisterMessage(SimulateResponse)


DESCRIPTOR._options = None

_SIMULATESERVICE = _descriptor.ServiceDescriptor(
  name='SimulateService',
  full_name='cosmos.base.simulate.v1beta1.SimulateService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=350,
  serialized_end=523,
  methods=[
  _descriptor.MethodDescriptor(
    name='Simulate',
    full_name='cosmos.base.simulate.v1beta1.SimulateService.Simulate',
    index=0,
    containing_service=None,
    input_type=_SIMULATEREQUEST,
    output_type=_SIMULATERESPONSE,
    serialized_options=b'\202\323\344\223\002(\"&/cosmos/base/simulate/v1beta1/simulate',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_SIMULATESERVICE)

DESCRIPTOR.services_by_name['SimulateService'] = _SIMULATESERVICE

# @@protoc_insertion_point(module_scope)