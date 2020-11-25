// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cosmos/mint/v1beta1/mint.proto

package types

import (
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Minter represents the minting state.
type Minter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// current annual inflation rate
	Inflation string `protobuf:"bytes,1,opt,name=inflation,proto3" json:"inflation,omitempty"`
	// current annual expected provisions
	AnnualProvisions string `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3" json:"annual_provisions,omitempty"`
}

func (x *Minter) Reset() {
	*x = Minter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_mint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Minter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Minter) ProtoMessage() {}

func (x *Minter) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_mint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Minter.ProtoReflect.Descriptor instead.
func (*Minter) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_mint_proto_rawDescGZIP(), []int{0}
}

func (x *Minter) GetInflation() string {
	if x != nil {
		return x.Inflation
	}
	return ""
}

func (x *Minter) GetAnnualProvisions() string {
	if x != nil {
		return x.AnnualProvisions
	}
	return ""
}

// Params holds parameters for the mint module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange string `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3" json:"inflation_rate_change,omitempty"`
	// maximum inflation rate
	InflationMax string `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3" json:"inflation_max,omitempty"`
	// minimum inflation rate
	InflationMin string `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3" json:"inflation_min,omitempty"`
	// goal of percent bonded atoms
	GoalBonded string `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3" json:"goal_bonded,omitempty"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_mint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_mint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_mint_proto_rawDescGZIP(), []int{1}
}

func (x *Params) GetMintDenom() string {
	if x != nil {
		return x.MintDenom
	}
	return ""
}

func (x *Params) GetInflationRateChange() string {
	if x != nil {
		return x.InflationRateChange
	}
	return ""
}

func (x *Params) GetInflationMax() string {
	if x != nil {
		return x.InflationMax
	}
	return ""
}

func (x *Params) GetInflationMin() string {
	if x != nil {
		return x.InflationMin
	}
	return ""
}

func (x *Params) GetGoalBonded() string {
	if x != nil {
		return x.GoalBonded
	}
	return ""
}

func (x *Params) GetBlocksPerYear() uint64 {
	if x != nil {
		return x.BlocksPerYear
	}
	return 0
}

var File_cosmos_mint_v1beta1_mint_proto protoreflect.FileDescriptor

var file_cosmos_mint_v1beta1_mint_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x06,
	0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x09, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xda, 0xde, 0x1f, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x66, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x4a, 0xf2, 0xde, 0x1f, 0x18, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x61, 0x6e, 0x6e, 0x75, 0x61,
	0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xda, 0xde, 0x1f,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x10, 0x61, 0x6e, 0x6e,
	0x75, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb7, 0x04,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x74,
	0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69,
	0x6e, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x69, 0x6e, 0x66, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4e, 0xf2, 0xde, 0x1f, 0x1c, 0x79, 0x61, 0x6d,
	0x6c, 0x3a, 0x22, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x13, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x6b, 0x0a, 0x0d,
	0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x46, 0xf2, 0xde, 0x1f, 0x14, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x69,
	0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x22, 0xda, 0xde, 0x1f,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x66,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x6b, 0x0a, 0x0d, 0x69, 0x6e, 0x66,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x46, 0xf2, 0xde, 0x1f, 0x14, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x69, 0x6e, 0x66, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x22, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x65, 0x0a, 0x0b, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x62,
	0x6f, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x44, 0xf2, 0xde, 0x1f,
	0x12, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x62, 0x6f, 0x6e, 0x64,
	0x65, 0x64, 0x22, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x42, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x42, 0x0a,
	0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a, 0xf2, 0xde, 0x1f, 0x16, 0x79, 0x61, 0x6d, 0x6c,
	0x3a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x79, 0x65, 0x61,
	0x72, 0x22, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x50, 0x65, 0x72, 0x59, 0x65, 0x61,
	0x72, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_mint_v1beta1_mint_proto_rawDescOnce sync.Once
	file_cosmos_mint_v1beta1_mint_proto_rawDescData = file_cosmos_mint_v1beta1_mint_proto_rawDesc
)

func file_cosmos_mint_v1beta1_mint_proto_rawDescGZIP() []byte {
	file_cosmos_mint_v1beta1_mint_proto_rawDescOnce.Do(func() {
		file_cosmos_mint_v1beta1_mint_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_mint_v1beta1_mint_proto_rawDescData)
	})
	return file_cosmos_mint_v1beta1_mint_proto_rawDescData
}

var file_cosmos_mint_v1beta1_mint_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_mint_v1beta1_mint_proto_goTypes = []interface{}{
	(*Minter)(nil), // 0: cosmos.mint.v1beta1.Minter
	(*Params)(nil), // 1: cosmos.mint.v1beta1.Params
}
var file_cosmos_mint_v1beta1_mint_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_mint_v1beta1_mint_proto_init() }
func file_cosmos_mint_v1beta1_mint_proto_init() {
	if File_cosmos_mint_v1beta1_mint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_mint_v1beta1_mint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Minter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_mint_v1beta1_mint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_mint_v1beta1_mint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_mint_v1beta1_mint_proto_goTypes,
		DependencyIndexes: file_cosmos_mint_v1beta1_mint_proto_depIdxs,
		MessageInfos:      file_cosmos_mint_v1beta1_mint_proto_msgTypes,
	}.Build()
	File_cosmos_mint_v1beta1_mint_proto = out.File
	file_cosmos_mint_v1beta1_mint_proto_rawDesc = nil
	file_cosmos_mint_v1beta1_mint_proto_goTypes = nil
	file_cosmos_mint_v1beta1_mint_proto_depIdxs = nil
}
