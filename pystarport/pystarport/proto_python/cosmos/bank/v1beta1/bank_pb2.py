# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cosmos/bank/v1beta1/bank.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from cosmos.base.v1beta1 import coin_pb2 as cosmos_dot_base_dot_v1beta1_dot_coin__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
    name="cosmos/bank/v1beta1/bank.proto",
    package="cosmos.bank.v1beta1",
    syntax="proto3",
    serialized_options=b"Z)github.com/cosmos/cosmos-sdk/x/bank/types",
    create_key=_descriptor._internal_create_key,
    serialized_pb=b'\n\x1e\x63osmos/bank/v1beta1/bank.proto\x12\x13\x63osmos.bank.v1beta1\x1a\x14gogoproto/gogo.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x1e\x63osmos/base/v1beta1/coin.proto"\xb2\x01\n\x06Params\x12Y\n\x0csend_enabled\x18\x01 \x03(\x0b\x32 .cosmos.bank.v1beta1.SendEnabledB!\xf2\xde\x1f\x1dyaml:"send_enabled,omitempty"\x12G\n\x14\x64\x65\x66\x61ult_send_enabled\x18\x02 \x01(\x08\x42)\xf2\xde\x1f%yaml:"default_send_enabled,omitempty":\x04\x98\xa0\x1f\x00"7\n\x0bSendEnabled\x12\r\n\x05\x64\x65nom\x18\x01 \x01(\t\x12\x0f\n\x07\x65nabled\x18\x02 \x01(\x08:\x08\xe8\xa0\x1f\x01\x98\xa0\x1f\x00"~\n\x05Input\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12Z\n\x05\x63oins\x18\x02 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.Coins:\x08\xe8\xa0\x1f\x00\x88\xa0\x1f\x00"\x7f\n\x06Output\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12Z\n\x05\x63oins\x18\x02 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.Coins:\x08\xe8\xa0\x1f\x00\x88\xa0\x1f\x00"\xab\x01\n\x06Supply\x12Z\n\x05total\x18\x01 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.Coins:E\xe8\xa0\x1f\x01\x88\xa0\x1f\x00\x98\xa0\x1f\x00\xd2\xb4-5*github.com/cosmos/cosmos-sdk/x/bank/exported.SupplyI"=\n\tDenomUnit\x12\r\n\x05\x64\x65nom\x18\x01 \x01(\t\x12\x10\n\x08\x65xponent\x18\x02 \x01(\r\x12\x0f\n\x07\x61liases\x18\x03 \x03(\t"s\n\x08Metadata\x12\x13\n\x0b\x64\x65scription\x18\x01 \x01(\t\x12\x33\n\x0b\x64\x65nom_units\x18\x02 \x03(\x0b\x32\x1e.cosmos.bank.v1beta1.DenomUnit\x12\x0c\n\x04\x62\x61se\x18\x03 \x01(\t\x12\x0f\n\x07\x64isplay\x18\x04 \x01(\tB+Z)github.com/cosmos/cosmos-sdk/x/bank/typesb\x06proto3',
    dependencies=[
        gogoproto_dot_gogo__pb2.DESCRIPTOR,
        cosmos__proto_dot_cosmos__pb2.DESCRIPTOR,
        cosmos_dot_base_dot_v1beta1_dot_coin__pb2.DESCRIPTOR,
    ],
)


_PARAMS = _descriptor.Descriptor(
    name="Params",
    full_name="cosmos.bank.v1beta1.Params",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="send_enabled",
            full_name="cosmos.bank.v1beta1.Params.send_enabled",
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b'\362\336\037\035yaml:"send_enabled,omitempty"',
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="default_send_enabled",
            full_name="cosmos.bank.v1beta1.Params.default_send_enabled",
            index=1,
            number=2,
            type=8,
            cpp_type=7,
            label=1,
            has_default_value=False,
            default_value=False,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b'\362\336\037%yaml:"default_send_enabled,omitempty"',
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\230\240\037\000",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=137,
    serialized_end=315,
)


_SENDENABLED = _descriptor.Descriptor(
    name="SendEnabled",
    full_name="cosmos.bank.v1beta1.SendEnabled",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="denom",
            full_name="cosmos.bank.v1beta1.SendEnabled.denom",
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="enabled",
            full_name="cosmos.bank.v1beta1.SendEnabled.enabled",
            index=1,
            number=2,
            type=8,
            cpp_type=7,
            label=1,
            has_default_value=False,
            default_value=False,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\350\240\037\001\230\240\037\000",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=317,
    serialized_end=372,
)


_INPUT = _descriptor.Descriptor(
    name="Input",
    full_name="cosmos.bank.v1beta1.Input",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="address",
            full_name="cosmos.bank.v1beta1.Input.address",
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="coins",
            full_name="cosmos.bank.v1beta1.Input.coins",
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\350\240\037\000\210\240\037\000",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=374,
    serialized_end=500,
)


_OUTPUT = _descriptor.Descriptor(
    name="Output",
    full_name="cosmos.bank.v1beta1.Output",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="address",
            full_name="cosmos.bank.v1beta1.Output.address",
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="coins",
            full_name="cosmos.bank.v1beta1.Output.coins",
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\350\240\037\000\210\240\037\000",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=502,
    serialized_end=629,
)


_SUPPLY = _descriptor.Descriptor(
    name="Supply",
    full_name="cosmos.bank.v1beta1.Supply",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="total",
            full_name="cosmos.bank.v1beta1.Supply.total",
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\350\240\037\001\210\240\037\000\230\240\037\000\322\264-5*github.com/cosmos/cosmos-sdk/x/bank/exported.SupplyI",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=632,
    serialized_end=803,
)


_DENOMUNIT = _descriptor.Descriptor(
    name="DenomUnit",
    full_name="cosmos.bank.v1beta1.DenomUnit",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="denom",
            full_name="cosmos.bank.v1beta1.DenomUnit.denom",
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="exponent",
            full_name="cosmos.bank.v1beta1.DenomUnit.exponent",
            index=1,
            number=2,
            type=13,
            cpp_type=3,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="aliases",
            full_name="cosmos.bank.v1beta1.DenomUnit.aliases",
            index=2,
            number=3,
            type=9,
            cpp_type=9,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=805,
    serialized_end=866,
)


_METADATA = _descriptor.Descriptor(
    name="Metadata",
    full_name="cosmos.bank.v1beta1.Metadata",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="description",
            full_name="cosmos.bank.v1beta1.Metadata.description",
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="denom_units",
            full_name="cosmos.bank.v1beta1.Metadata.denom_units",
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="base",
            full_name="cosmos.bank.v1beta1.Metadata.base",
            index=2,
            number=3,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="display",
            full_name="cosmos.bank.v1beta1.Metadata.display",
            index=3,
            number=4,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode("utf-8"),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=868,
    serialized_end=983,
)

_PARAMS.fields_by_name["send_enabled"].message_type = _SENDENABLED
_INPUT.fields_by_name[
    "coins"
].message_type = cosmos_dot_base_dot_v1beta1_dot_coin__pb2._COIN
_OUTPUT.fields_by_name[
    "coins"
].message_type = cosmos_dot_base_dot_v1beta1_dot_coin__pb2._COIN
_SUPPLY.fields_by_name[
    "total"
].message_type = cosmos_dot_base_dot_v1beta1_dot_coin__pb2._COIN
_METADATA.fields_by_name["denom_units"].message_type = _DENOMUNIT
DESCRIPTOR.message_types_by_name["Params"] = _PARAMS
DESCRIPTOR.message_types_by_name["SendEnabled"] = _SENDENABLED
DESCRIPTOR.message_types_by_name["Input"] = _INPUT
DESCRIPTOR.message_types_by_name["Output"] = _OUTPUT
DESCRIPTOR.message_types_by_name["Supply"] = _SUPPLY
DESCRIPTOR.message_types_by_name["DenomUnit"] = _DENOMUNIT
DESCRIPTOR.message_types_by_name["Metadata"] = _METADATA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Params = _reflection.GeneratedProtocolMessageType(
    "Params",
    (_message.Message,),
    {
        "DESCRIPTOR": _PARAMS,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.Params)
    },
)
_sym_db.RegisterMessage(Params)

SendEnabled = _reflection.GeneratedProtocolMessageType(
    "SendEnabled",
    (_message.Message,),
    {
        "DESCRIPTOR": _SENDENABLED,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.SendEnabled)
    },
)
_sym_db.RegisterMessage(SendEnabled)

Input = _reflection.GeneratedProtocolMessageType(
    "Input",
    (_message.Message,),
    {
        "DESCRIPTOR": _INPUT,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.Input)
    },
)
_sym_db.RegisterMessage(Input)

Output = _reflection.GeneratedProtocolMessageType(
    "Output",
    (_message.Message,),
    {
        "DESCRIPTOR": _OUTPUT,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.Output)
    },
)
_sym_db.RegisterMessage(Output)

Supply = _reflection.GeneratedProtocolMessageType(
    "Supply",
    (_message.Message,),
    {
        "DESCRIPTOR": _SUPPLY,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.Supply)
    },
)
_sym_db.RegisterMessage(Supply)

DenomUnit = _reflection.GeneratedProtocolMessageType(
    "DenomUnit",
    (_message.Message,),
    {
        "DESCRIPTOR": _DENOMUNIT,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.DenomUnit)
    },
)
_sym_db.RegisterMessage(DenomUnit)

Metadata = _reflection.GeneratedProtocolMessageType(
    "Metadata",
    (_message.Message,),
    {
        "DESCRIPTOR": _METADATA,
        "__module__": "cosmos.bank.v1beta1.bank_pb2"
        # @@protoc_insertion_point(class_scope:cosmos.bank.v1beta1.Metadata)
    },
)
_sym_db.RegisterMessage(Metadata)


DESCRIPTOR._options = None
_PARAMS.fields_by_name["send_enabled"]._options = None
_PARAMS.fields_by_name["default_send_enabled"]._options = None
_PARAMS._options = None
_SENDENABLED._options = None
_INPUT.fields_by_name["coins"]._options = None
_INPUT._options = None
_OUTPUT.fields_by_name["coins"]._options = None
_OUTPUT._options = None
_SUPPLY.fields_by_name["total"]._options = None
_SUPPLY._options = None
# @@protoc_insertion_point(module_scope)
