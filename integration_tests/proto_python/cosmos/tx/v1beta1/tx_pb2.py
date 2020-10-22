# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cosmos/tx/v1beta1/tx.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from cosmos.crypto.multisig.v1beta1 import multisig_pb2 as cosmos_dot_crypto_dot_multisig_dot_v1beta1_dot_multisig__pb2
from cosmos.base.v1beta1 import coin_pb2 as cosmos_dot_base_dot_v1beta1_dot_coin__pb2
from cosmos.tx.signing.v1beta1 import signing_pb2 as cosmos_dot_tx_dot_signing_dot_v1beta1_dot_signing__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='cosmos/tx/v1beta1/tx.proto',
  package='cosmos.tx.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z%github.com/cosmos/cosmos-sdk/types/tx'),
  serialized_pb=_b('\n\x1a\x63osmos/tx/v1beta1/tx.proto\x12\x11\x63osmos.tx.v1beta1\x1a\x14gogoproto/gogo.proto\x1a-cosmos/crypto/multisig/v1beta1/multisig.proto\x1a\x1e\x63osmos/base/v1beta1/coin.proto\x1a\'cosmos/tx/signing/v1beta1/signing.proto\x1a\x19google/protobuf/any.proto\"q\n\x02Tx\x12\'\n\x04\x62ody\x18\x01 \x01(\x0b\x32\x19.cosmos.tx.v1beta1.TxBody\x12.\n\tauth_info\x18\x02 \x01(\x0b\x32\x1b.cosmos.tx.v1beta1.AuthInfo\x12\x12\n\nsignatures\x18\x03 \x03(\x0c\"H\n\x05TxRaw\x12\x12\n\nbody_bytes\x18\x01 \x01(\x0c\x12\x17\n\x0f\x61uth_info_bytes\x18\x02 \x01(\x0c\x12\x12\n\nsignatures\x18\x03 \x03(\x0c\"`\n\x07SignDoc\x12\x12\n\nbody_bytes\x18\x01 \x01(\x0c\x12\x17\n\x0f\x61uth_info_bytes\x18\x02 \x01(\x0c\x12\x10\n\x08\x63hain_id\x18\x03 \x01(\t\x12\x16\n\x0e\x61\x63\x63ount_number\x18\x04 \x01(\x04\"\xc7\x01\n\x06TxBody\x12&\n\x08messages\x18\x01 \x03(\x0b\x32\x14.google.protobuf.Any\x12\x0c\n\x04memo\x18\x02 \x01(\t\x12\x16\n\x0etimeout_height\x18\x03 \x01(\x04\x12\x30\n\x11\x65xtension_options\x18\xff\x07 \x03(\x0b\x32\x14.google.protobuf.Any\x12=\n\x1enon_critical_extension_options\x18\xff\x0f \x03(\x0b\x32\x14.google.protobuf.Any\"d\n\x08\x41uthInfo\x12\x33\n\x0csigner_infos\x18\x01 \x03(\x0b\x32\x1d.cosmos.tx.v1beta1.SignerInfo\x12#\n\x03\x66\x65\x65\x18\x02 \x01(\x0b\x32\x16.cosmos.tx.v1beta1.Fee\"x\n\nSignerInfo\x12(\n\npublic_key\x18\x01 \x01(\x0b\x32\x14.google.protobuf.Any\x12.\n\tmode_info\x18\x02 \x01(\x0b\x32\x1b.cosmos.tx.v1beta1.ModeInfo\x12\x10\n\x08sequence\x18\x03 \x01(\x04\"\xb5\x02\n\x08ModeInfo\x12\x34\n\x06single\x18\x01 \x01(\x0b\x32\".cosmos.tx.v1beta1.ModeInfo.SingleH\x00\x12\x32\n\x05multi\x18\x02 \x01(\x0b\x32!.cosmos.tx.v1beta1.ModeInfo.MultiH\x00\x1a;\n\x06Single\x12\x31\n\x04mode\x18\x01 \x01(\x0e\x32#.cosmos.tx.signing.v1beta1.SignMode\x1a{\n\x05Multi\x12\x41\n\x08\x62itarray\x18\x01 \x01(\x0b\x32/.cosmos.crypto.multisig.v1beta1.CompactBitArray\x12/\n\nmode_infos\x18\x02 \x03(\x0b\x32\x1b.cosmos.tx.v1beta1.ModeInfoB\x05\n\x03sum\"\x95\x01\n\x03\x46\x65\x65\x12[\n\x06\x61mount\x18\x01 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.Coins\x12\x11\n\tgas_limit\x18\x02 \x01(\x04\x12\r\n\x05payer\x18\x03 \x01(\t\x12\x0f\n\x07granter\x18\x04 \x01(\tB\'Z%github.com/cosmos/cosmos-sdk/types/txb\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,cosmos_dot_crypto_dot_multisig_dot_v1beta1_dot_multisig__pb2.DESCRIPTOR,cosmos_dot_base_dot_v1beta1_dot_coin__pb2.DESCRIPTOR,cosmos_dot_tx_dot_signing_dot_v1beta1_dot_signing__pb2.DESCRIPTOR,google_dot_protobuf_dot_any__pb2.DESCRIPTOR,])




_TX = _descriptor.Descriptor(
  name='Tx',
  full_name='cosmos.tx.v1beta1.Tx',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='body', full_name='cosmos.tx.v1beta1.Tx.body', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_info', full_name='cosmos.tx.v1beta1.Tx.auth_info', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signatures', full_name='cosmos.tx.v1beta1.Tx.signatures', index=2,
      number=3, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=218,
  serialized_end=331,
)


_TXRAW = _descriptor.Descriptor(
  name='TxRaw',
  full_name='cosmos.tx.v1beta1.TxRaw',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='body_bytes', full_name='cosmos.tx.v1beta1.TxRaw.body_bytes', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_info_bytes', full_name='cosmos.tx.v1beta1.TxRaw.auth_info_bytes', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signatures', full_name='cosmos.tx.v1beta1.TxRaw.signatures', index=2,
      number=3, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=333,
  serialized_end=405,
)


_SIGNDOC = _descriptor.Descriptor(
  name='SignDoc',
  full_name='cosmos.tx.v1beta1.SignDoc',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='body_bytes', full_name='cosmos.tx.v1beta1.SignDoc.body_bytes', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_info_bytes', full_name='cosmos.tx.v1beta1.SignDoc.auth_info_bytes', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='chain_id', full_name='cosmos.tx.v1beta1.SignDoc.chain_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='account_number', full_name='cosmos.tx.v1beta1.SignDoc.account_number', index=3,
      number=4, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=407,
  serialized_end=503,
)


_TXBODY = _descriptor.Descriptor(
  name='TxBody',
  full_name='cosmos.tx.v1beta1.TxBody',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='messages', full_name='cosmos.tx.v1beta1.TxBody.messages', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='memo', full_name='cosmos.tx.v1beta1.TxBody.memo', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timeout_height', full_name='cosmos.tx.v1beta1.TxBody.timeout_height', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='extension_options', full_name='cosmos.tx.v1beta1.TxBody.extension_options', index=3,
      number=1023, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='non_critical_extension_options', full_name='cosmos.tx.v1beta1.TxBody.non_critical_extension_options', index=4,
      number=2047, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=506,
  serialized_end=705,
)


_AUTHINFO = _descriptor.Descriptor(
  name='AuthInfo',
  full_name='cosmos.tx.v1beta1.AuthInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='signer_infos', full_name='cosmos.tx.v1beta1.AuthInfo.signer_infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fee', full_name='cosmos.tx.v1beta1.AuthInfo.fee', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=707,
  serialized_end=807,
)


_SIGNERINFO = _descriptor.Descriptor(
  name='SignerInfo',
  full_name='cosmos.tx.v1beta1.SignerInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='public_key', full_name='cosmos.tx.v1beta1.SignerInfo.public_key', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mode_info', full_name='cosmos.tx.v1beta1.SignerInfo.mode_info', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sequence', full_name='cosmos.tx.v1beta1.SignerInfo.sequence', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=809,
  serialized_end=929,
)


_MODEINFO_SINGLE = _descriptor.Descriptor(
  name='Single',
  full_name='cosmos.tx.v1beta1.ModeInfo.Single',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='cosmos.tx.v1beta1.ModeInfo.Single.mode', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1050,
  serialized_end=1109,
)

_MODEINFO_MULTI = _descriptor.Descriptor(
  name='Multi',
  full_name='cosmos.tx.v1beta1.ModeInfo.Multi',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bitarray', full_name='cosmos.tx.v1beta1.ModeInfo.Multi.bitarray', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mode_infos', full_name='cosmos.tx.v1beta1.ModeInfo.Multi.mode_infos', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=1111,
  serialized_end=1234,
)

_MODEINFO = _descriptor.Descriptor(
  name='ModeInfo',
  full_name='cosmos.tx.v1beta1.ModeInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='single', full_name='cosmos.tx.v1beta1.ModeInfo.single', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='multi', full_name='cosmos.tx.v1beta1.ModeInfo.multi', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_MODEINFO_SINGLE, _MODEINFO_MULTI, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='sum', full_name='cosmos.tx.v1beta1.ModeInfo.sum',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=932,
  serialized_end=1241,
)


_FEE = _descriptor.Descriptor(
  name='Fee',
  full_name='cosmos.tx.v1beta1.Fee',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='amount', full_name='cosmos.tx.v1beta1.Fee.amount', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gas_limit', full_name='cosmos.tx.v1beta1.Fee.gas_limit', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='payer', full_name='cosmos.tx.v1beta1.Fee.payer', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='granter', full_name='cosmos.tx.v1beta1.Fee.granter', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1244,
  serialized_end=1393,
)

_TX.fields_by_name['body'].message_type = _TXBODY
_TX.fields_by_name['auth_info'].message_type = _AUTHINFO
_TXBODY.fields_by_name['messages'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_TXBODY.fields_by_name['extension_options'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_TXBODY.fields_by_name['non_critical_extension_options'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_AUTHINFO.fields_by_name['signer_infos'].message_type = _SIGNERINFO
_AUTHINFO.fields_by_name['fee'].message_type = _FEE
_SIGNERINFO.fields_by_name['public_key'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_SIGNERINFO.fields_by_name['mode_info'].message_type = _MODEINFO
_MODEINFO_SINGLE.fields_by_name['mode'].enum_type = cosmos_dot_tx_dot_signing_dot_v1beta1_dot_signing__pb2._SIGNMODE
_MODEINFO_SINGLE.containing_type = _MODEINFO
_MODEINFO_MULTI.fields_by_name['bitarray'].message_type = cosmos_dot_crypto_dot_multisig_dot_v1beta1_dot_multisig__pb2._COMPACTBITARRAY
_MODEINFO_MULTI.fields_by_name['mode_infos'].message_type = _MODEINFO
_MODEINFO_MULTI.containing_type = _MODEINFO
_MODEINFO.fields_by_name['single'].message_type = _MODEINFO_SINGLE
_MODEINFO.fields_by_name['multi'].message_type = _MODEINFO_MULTI
_MODEINFO.oneofs_by_name['sum'].fields.append(
  _MODEINFO.fields_by_name['single'])
_MODEINFO.fields_by_name['single'].containing_oneof = _MODEINFO.oneofs_by_name['sum']
_MODEINFO.oneofs_by_name['sum'].fields.append(
  _MODEINFO.fields_by_name['multi'])
_MODEINFO.fields_by_name['multi'].containing_oneof = _MODEINFO.oneofs_by_name['sum']
_FEE.fields_by_name['amount'].message_type = cosmos_dot_base_dot_v1beta1_dot_coin__pb2._COIN
DESCRIPTOR.message_types_by_name['Tx'] = _TX
DESCRIPTOR.message_types_by_name['TxRaw'] = _TXRAW
DESCRIPTOR.message_types_by_name['SignDoc'] = _SIGNDOC
DESCRIPTOR.message_types_by_name['TxBody'] = _TXBODY
DESCRIPTOR.message_types_by_name['AuthInfo'] = _AUTHINFO
DESCRIPTOR.message_types_by_name['SignerInfo'] = _SIGNERINFO
DESCRIPTOR.message_types_by_name['ModeInfo'] = _MODEINFO
DESCRIPTOR.message_types_by_name['Fee'] = _FEE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Tx = _reflection.GeneratedProtocolMessageType('Tx', (_message.Message,), dict(
  DESCRIPTOR = _TX,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.Tx)
  ))
_sym_db.RegisterMessage(Tx)

TxRaw = _reflection.GeneratedProtocolMessageType('TxRaw', (_message.Message,), dict(
  DESCRIPTOR = _TXRAW,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.TxRaw)
  ))
_sym_db.RegisterMessage(TxRaw)

SignDoc = _reflection.GeneratedProtocolMessageType('SignDoc', (_message.Message,), dict(
  DESCRIPTOR = _SIGNDOC,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.SignDoc)
  ))
_sym_db.RegisterMessage(SignDoc)

TxBody = _reflection.GeneratedProtocolMessageType('TxBody', (_message.Message,), dict(
  DESCRIPTOR = _TXBODY,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.TxBody)
  ))
_sym_db.RegisterMessage(TxBody)

AuthInfo = _reflection.GeneratedProtocolMessageType('AuthInfo', (_message.Message,), dict(
  DESCRIPTOR = _AUTHINFO,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.AuthInfo)
  ))
_sym_db.RegisterMessage(AuthInfo)

SignerInfo = _reflection.GeneratedProtocolMessageType('SignerInfo', (_message.Message,), dict(
  DESCRIPTOR = _SIGNERINFO,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.SignerInfo)
  ))
_sym_db.RegisterMessage(SignerInfo)

ModeInfo = _reflection.GeneratedProtocolMessageType('ModeInfo', (_message.Message,), dict(

  Single = _reflection.GeneratedProtocolMessageType('Single', (_message.Message,), dict(
    DESCRIPTOR = _MODEINFO_SINGLE,
    __module__ = 'cosmos.tx.v1beta1.tx_pb2'
    # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.ModeInfo.Single)
    ))
  ,

  Multi = _reflection.GeneratedProtocolMessageType('Multi', (_message.Message,), dict(
    DESCRIPTOR = _MODEINFO_MULTI,
    __module__ = 'cosmos.tx.v1beta1.tx_pb2'
    # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.ModeInfo.Multi)
    ))
  ,
  DESCRIPTOR = _MODEINFO,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.ModeInfo)
  ))
_sym_db.RegisterMessage(ModeInfo)
_sym_db.RegisterMessage(ModeInfo.Single)
_sym_db.RegisterMessage(ModeInfo.Multi)

Fee = _reflection.GeneratedProtocolMessageType('Fee', (_message.Message,), dict(
  DESCRIPTOR = _FEE,
  __module__ = 'cosmos.tx.v1beta1.tx_pb2'
  # @@protoc_insertion_point(class_scope:cosmos.tx.v1beta1.Fee)
  ))
_sym_db.RegisterMessage(Fee)


DESCRIPTOR._options = None
_FEE.fields_by_name['amount']._options = None
# @@protoc_insertion_point(module_scope)
