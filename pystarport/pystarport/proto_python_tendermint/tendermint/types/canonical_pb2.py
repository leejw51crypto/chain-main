# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tendermint/types/canonical.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from tendermint.types import types_pb2 as tendermint_dot_types_dot_types__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='tendermint/types/canonical.proto',
  package='tendermint.types',
  syntax='proto3',
  serialized_options=b'Z7github.com/tendermint/tendermint/proto/tendermint/types',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n tendermint/types/canonical.proto\x12\x10tendermint.types\x1a\x14gogoproto/gogo.proto\x1a\x1ctendermint/types/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"i\n\x10\x43\x61nonicalBlockID\x12\x0c\n\x04hash\x18\x01 \x01(\x0c\x12G\n\x0fpart_set_header\x18\x02 \x01(\x0b\x32(.tendermint.types.CanonicalPartSetHeaderB\x04\xc8\xde\x1f\x00\"5\n\x16\x43\x61nonicalPartSetHeader\x12\r\n\x05total\x18\x01 \x01(\r\x12\x0c\n\x04hash\x18\x02 \x01(\x0c\"\x9d\x02\n\x11\x43\x61nonicalProposal\x12-\n\x04type\x18\x01 \x01(\x0e\x32\x1f.tendermint.types.SignedMsgType\x12\x0e\n\x06height\x18\x02 \x01(\x10\x12\r\n\x05round\x18\x03 \x01(\x10\x12\x1f\n\tpol_round\x18\x04 \x01(\x03\x42\x0c\xe2\xde\x1f\x08POLRound\x12\x41\n\x08\x62lock_id\x18\x05 \x01(\x0b\x32\".tendermint.types.CanonicalBlockIDB\x0b\xe2\xde\x1f\x07\x42lockID\x12\x37\n\ttimestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01\x12\x1d\n\x08\x63hain_id\x18\x07 \x01(\tB\x0b\xe2\xde\x1f\x07\x43hainID\"\xf8\x01\n\rCanonicalVote\x12-\n\x04type\x18\x01 \x01(\x0e\x32\x1f.tendermint.types.SignedMsgType\x12\x0e\n\x06height\x18\x02 \x01(\x10\x12\r\n\x05round\x18\x03 \x01(\x10\x12\x41\n\x08\x62lock_id\x18\x04 \x01(\x0b\x32\".tendermint.types.CanonicalBlockIDB\x0b\xe2\xde\x1f\x07\x42lockID\x12\x37\n\ttimestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01\x12\x1d\n\x08\x63hain_id\x18\x06 \x01(\tB\x0b\xe2\xde\x1f\x07\x43hainIDB9Z7github.com/tendermint/tendermint/proto/tendermint/typesb\x06proto3'
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,tendermint_dot_types_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_CANONICALBLOCKID = _descriptor.Descriptor(
  name='CanonicalBlockID',
  full_name='tendermint.types.CanonicalBlockID',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='hash', full_name='tendermint.types.CanonicalBlockID.hash', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='part_set_header', full_name='tendermint.types.CanonicalBlockID.part_set_header', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\310\336\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=139,
  serialized_end=244,
)


_CANONICALPARTSETHEADER = _descriptor.Descriptor(
  name='CanonicalPartSetHeader',
  full_name='tendermint.types.CanonicalPartSetHeader',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='total', full_name='tendermint.types.CanonicalPartSetHeader.total', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='hash', full_name='tendermint.types.CanonicalPartSetHeader.hash', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
  serialized_start=246,
  serialized_end=299,
)


_CANONICALPROPOSAL = _descriptor.Descriptor(
  name='CanonicalProposal',
  full_name='tendermint.types.CanonicalProposal',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='tendermint.types.CanonicalProposal.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='height', full_name='tendermint.types.CanonicalProposal.height', index=1,
      number=2, type=16, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='round', full_name='tendermint.types.CanonicalProposal.round', index=2,
      number=3, type=16, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pol_round', full_name='tendermint.types.CanonicalProposal.pol_round', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\336\037\010POLRound', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='block_id', full_name='tendermint.types.CanonicalProposal.block_id', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\336\037\007BlockID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='tendermint.types.CanonicalProposal.timestamp', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\310\336\037\000\220\337\037\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='chain_id', full_name='tendermint.types.CanonicalProposal.chain_id', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\336\037\007ChainID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=302,
  serialized_end=587,
)


_CANONICALVOTE = _descriptor.Descriptor(
  name='CanonicalVote',
  full_name='tendermint.types.CanonicalVote',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='tendermint.types.CanonicalVote.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='height', full_name='tendermint.types.CanonicalVote.height', index=1,
      number=2, type=16, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='round', full_name='tendermint.types.CanonicalVote.round', index=2,
      number=3, type=16, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='block_id', full_name='tendermint.types.CanonicalVote.block_id', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\336\037\007BlockID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='tendermint.types.CanonicalVote.timestamp', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\310\336\037\000\220\337\037\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='chain_id', full_name='tendermint.types.CanonicalVote.chain_id', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\336\037\007ChainID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=590,
  serialized_end=838,
)

_CANONICALBLOCKID.fields_by_name['part_set_header'].message_type = _CANONICALPARTSETHEADER
_CANONICALPROPOSAL.fields_by_name['type'].enum_type = tendermint_dot_types_dot_types__pb2._SIGNEDMSGTYPE
_CANONICALPROPOSAL.fields_by_name['block_id'].message_type = _CANONICALBLOCKID
_CANONICALPROPOSAL.fields_by_name['timestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CANONICALVOTE.fields_by_name['type'].enum_type = tendermint_dot_types_dot_types__pb2._SIGNEDMSGTYPE
_CANONICALVOTE.fields_by_name['block_id'].message_type = _CANONICALBLOCKID
_CANONICALVOTE.fields_by_name['timestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['CanonicalBlockID'] = _CANONICALBLOCKID
DESCRIPTOR.message_types_by_name['CanonicalPartSetHeader'] = _CANONICALPARTSETHEADER
DESCRIPTOR.message_types_by_name['CanonicalProposal'] = _CANONICALPROPOSAL
DESCRIPTOR.message_types_by_name['CanonicalVote'] = _CANONICALVOTE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CanonicalBlockID = _reflection.GeneratedProtocolMessageType('CanonicalBlockID', (_message.Message,), {
  'DESCRIPTOR' : _CANONICALBLOCKID,
  '__module__' : 'tendermint.types.canonical_pb2'
  # @@protoc_insertion_point(class_scope:tendermint.types.CanonicalBlockID)
  })
_sym_db.RegisterMessage(CanonicalBlockID)

CanonicalPartSetHeader = _reflection.GeneratedProtocolMessageType('CanonicalPartSetHeader', (_message.Message,), {
  'DESCRIPTOR' : _CANONICALPARTSETHEADER,
  '__module__' : 'tendermint.types.canonical_pb2'
  # @@protoc_insertion_point(class_scope:tendermint.types.CanonicalPartSetHeader)
  })
_sym_db.RegisterMessage(CanonicalPartSetHeader)

CanonicalProposal = _reflection.GeneratedProtocolMessageType('CanonicalProposal', (_message.Message,), {
  'DESCRIPTOR' : _CANONICALPROPOSAL,
  '__module__' : 'tendermint.types.canonical_pb2'
  # @@protoc_insertion_point(class_scope:tendermint.types.CanonicalProposal)
  })
_sym_db.RegisterMessage(CanonicalProposal)

CanonicalVote = _reflection.GeneratedProtocolMessageType('CanonicalVote', (_message.Message,), {
  'DESCRIPTOR' : _CANONICALVOTE,
  '__module__' : 'tendermint.types.canonical_pb2'
  # @@protoc_insertion_point(class_scope:tendermint.types.CanonicalVote)
  })
_sym_db.RegisterMessage(CanonicalVote)


DESCRIPTOR._options = None
_CANONICALBLOCKID.fields_by_name['part_set_header']._options = None
_CANONICALPROPOSAL.fields_by_name['pol_round']._options = None
_CANONICALPROPOSAL.fields_by_name['block_id']._options = None
_CANONICALPROPOSAL.fields_by_name['timestamp']._options = None
_CANONICALPROPOSAL.fields_by_name['chain_id']._options = None
_CANONICALVOTE.fields_by_name['block_id']._options = None
_CANONICALVOTE.fields_by_name['timestamp']._options = None
_CANONICALVOTE.fields_by_name['chain_id']._options = None
# @@protoc_insertion_point(module_scope)
