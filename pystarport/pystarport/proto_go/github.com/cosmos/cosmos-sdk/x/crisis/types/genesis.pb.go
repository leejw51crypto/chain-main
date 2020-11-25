// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cosmos/crisis/v1beta1/genesis.proto

package types

import (
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the crisis module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// constant_fee is the fee used to verify the invariant in the crisis
	// module.
	ConstantFee *types.Coin `protobuf:"bytes,3,opt,name=constant_fee,json=constantFee,proto3" json:"constant_fee,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crisis_v1beta1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crisis_v1beta1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_cosmos_crisis_v1beta1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetConstantFee() *types.Coin {
	if x != nil {
		return x.ConstantFee
	}
	return nil
}

var File_cosmos_crisis_v1beta1_genesis_proto protoreflect.FileDescriptor

var file_cosmos_crisis_v1beta1_genesis_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72,
	0x69, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x42, 0x1b, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde, 0x1f, 0x13, 0x79, 0x61, 0x6d,
	0x6c, 0x3a, 0x22, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x22,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78, 0x2f,
	0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crisis_v1beta1_genesis_proto_rawDescOnce sync.Once
	file_cosmos_crisis_v1beta1_genesis_proto_rawDescData = file_cosmos_crisis_v1beta1_genesis_proto_rawDesc
)

func file_cosmos_crisis_v1beta1_genesis_proto_rawDescGZIP() []byte {
	file_cosmos_crisis_v1beta1_genesis_proto_rawDescOnce.Do(func() {
		file_cosmos_crisis_v1beta1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crisis_v1beta1_genesis_proto_rawDescData)
	})
	return file_cosmos_crisis_v1beta1_genesis_proto_rawDescData
}

var file_cosmos_crisis_v1beta1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_crisis_v1beta1_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil), // 0: cosmos.crisis.v1beta1.GenesisState
	(*types.Coin)(nil),   // 1: cosmos.base.v1beta1.Coin
}
var file_cosmos_crisis_v1beta1_genesis_proto_depIdxs = []int32{
	1, // 0: cosmos.crisis.v1beta1.GenesisState.constant_fee:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_crisis_v1beta1_genesis_proto_init() }
func file_cosmos_crisis_v1beta1_genesis_proto_init() {
	if File_cosmos_crisis_v1beta1_genesis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crisis_v1beta1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_cosmos_crisis_v1beta1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_crisis_v1beta1_genesis_proto_goTypes,
		DependencyIndexes: file_cosmos_crisis_v1beta1_genesis_proto_depIdxs,
		MessageInfos:      file_cosmos_crisis_v1beta1_genesis_proto_msgTypes,
	}.Build()
	File_cosmos_crisis_v1beta1_genesis_proto = out.File
	file_cosmos_crisis_v1beta1_genesis_proto_rawDesc = nil
	file_cosmos_crisis_v1beta1_genesis_proto_goTypes = nil
	file_cosmos_crisis_v1beta1_genesis_proto_depIdxs = nil
}
