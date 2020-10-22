# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cosmos/evidence/v1beta1/tx.proto

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
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='cosmos/evidence/v1beta1/tx.proto',
  package='cosmos.evidence.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z-github.com/cosmos/cosmos-sdk/x/evidence/types\250\342\036\001'),
  serialized_pb=_b('\n cosmos/evidence/v1beta1/tx.proto\x12\x17\x63osmos.evidence.v1beta1\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\x1a\x19\x63osmos_proto/cosmos.proto\"f\n\x11MsgSubmitEvidence\x12\x11\n\tsubmitter\x18\x01 \x01(\t\x12\x34\n\x08\x65vidence\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyB\x0c\xca\xb4-\x08\x45vidence:\x08\xe8\xa0\x1f\x00\x88\xa0\x1f\x00\")\n\x19MsgSubmitEvidenceResponse\x12\x0c\n\x04hash\x18\x04 \x01(\x0c\x32w\n\x03Msg\x12p\n\x0eSubmitEvidence\x12*.cosmos.evidence.v1beta1.MsgSubmitEvidence\x1a\x32.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponseB3Z-github.com/cosmos/cosmos-sdk/x/evidence/types\xa8\xe2\x1e\x01\x62\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,google_dot_protobuf_dot_any__pb2.DESCRIPTOR,cosmos__proto_dot_cosmos__pb2.DESCRIPTOR,])




_MSGSUBMITEVIDENCE = _descriptor.Descriptor(
  name='MsgSubmitEvidence',
  full_name='cosmos.evidence.v1beta1.MsgSubmitEvidence',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='submitter', full_name='cosmos.evidence.v1beta1.MsgSubmitEvidence.submitter', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='evidence', full_name='cosmos.evidence.v1beta1.MsgSubmitEvidence.evidence', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\312\264-\010Evidence'), file=DESCRIPTOR),
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
  serialized_start=137,
  serialized_end=239,
)


_MSGSUBMITEVIDENCERESPONSE = _descriptor.Descriptor(
  name='MsgSubmitEvidenceResponse',
  full_name='cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hash', full_name='cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.hash', index=0,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=241,
  serialized_end=282,
)

_MSGSUBMITEVIDENCE.fields_by_name['evidence'].message_type = google_dot_protobuf_dot_any__pb2._ANY
DESCRIPTOR.message_types_by_name['MsgSubmitEvidence'] = _MSGSUBMITEVIDENCE
DESCRIPTOR.message_types_by_name['MsgSubmitEvidenceResponse'] = _MSGSUBMITEVIDENCERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MsgSubmitEvidence = _reflection.GeneratedProtocolMessageType('MsgSubmitEvidence', (_message.Message,), dict(
  DESCRIPTOR = _MSGSUBMITEVIDENCE,
  __module__ = 'cosmos.evidence.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.evidence.v1beta1.MsgSubmitEvidence)
  ))
_sym_db.RegisterMessage(MsgSubmitEvidence)

MsgSubmitEvidenceResponse = _reflection.GeneratedProtocolMessageType('MsgSubmitEvidenceResponse', (_message.Message,), dict(
  DESCRIPTOR = _MSGSUBMITEVIDENCERESPONSE,
  __module__ = 'cosmos.evidence.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse)
  ))
_sym_db.RegisterMessage(MsgSubmitEvidenceResponse)


DESCRIPTOR._options = None
_MSGSUBMITEVIDENCE.fields_by_name['evidence']._options = None
_MSGSUBMITEVIDENCE._options = None

_MSG = _descriptor.ServiceDescriptor(
  name='Msg',
  full_name='cosmos.evidence.v1beta1.Msg',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=284,
  serialized_end=403,
  methods=[
  _descriptor.MethodDescriptor(
    name='SubmitEvidence',
    full_name='cosmos.evidence.v1beta1.Msg.SubmitEvidence',
    index=0,
    containing_service=None,
    input_type=_MSGSUBMITEVIDENCE,
    output_type=_MSGSUBMITEVIDENCERESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_MSG)

DESCRIPTOR.services_by_name['Msg'] = _MSG

# @@protoc_insertion_point(module_scope)
