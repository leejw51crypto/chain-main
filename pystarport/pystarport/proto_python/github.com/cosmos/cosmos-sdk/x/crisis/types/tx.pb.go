// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cosmos/crisis/v1beta1/tx.proto

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

// MsgVerifyInvariant represents a message to verify a particular invariance.
type MsgVerifyInvariant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender              string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	InvariantModuleName string `protobuf:"bytes,2,opt,name=invariant_module_name,json=invariantModuleName,proto3" json:"invariant_module_name,omitempty"`
	InvariantRoute      string `protobuf:"bytes,3,opt,name=invariant_route,json=invariantRoute,proto3" json:"invariant_route,omitempty"`
}

func (x *MsgVerifyInvariant) Reset() {
	*x = MsgVerifyInvariant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVerifyInvariant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVerifyInvariant) ProtoMessage() {}

func (x *MsgVerifyInvariant) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgVerifyInvariant.ProtoReflect.Descriptor instead.
func (*MsgVerifyInvariant) Descriptor() ([]byte, []int) {
	return file_cosmos_crisis_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgVerifyInvariant) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgVerifyInvariant) GetInvariantModuleName() string {
	if x != nil {
		return x.InvariantModuleName
	}
	return ""
}

func (x *MsgVerifyInvariant) GetInvariantRoute() string {
	if x != nil {
		return x.InvariantRoute
	}
	return ""
}

// MsgVerifyInvariantResponse defines the Msg/VerifyInvariant response type.
type MsgVerifyInvariantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgVerifyInvariantResponse) Reset() {
	*x = MsgVerifyInvariantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVerifyInvariantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVerifyInvariantResponse) ProtoMessage() {}

func (x *MsgVerifyInvariantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crisis_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgVerifyInvariantResponse.ProtoReflect.Descriptor instead.
func (*MsgVerifyInvariantResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_crisis_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

var File_cosmos_crisis_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_crisis_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01,
	0x0a, 0x12, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x15,
	0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xf2, 0xde, 0x1f,
	0x1c, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x13, 0x69,
	0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xf2, 0xde, 0x1f,
	0x16, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x3a, 0x08, 0xe8, 0xa0, 0x1f, 0x00, 0x88, 0xa0, 0x1f,
	0x00, 0x22, 0x1c, 0x0a, 0x1a, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x76, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x6f, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x49, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72,
	0x69, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x69, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crisis_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_crisis_v1beta1_tx_proto_rawDescData = file_cosmos_crisis_v1beta1_tx_proto_rawDesc
)

func file_cosmos_crisis_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_crisis_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_crisis_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crisis_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_crisis_v1beta1_tx_proto_rawDescData
}

var file_cosmos_crisis_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_crisis_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgVerifyInvariant)(nil),         // 0: cosmos.crisis.v1beta1.MsgVerifyInvariant
	(*MsgVerifyInvariantResponse)(nil), // 1: cosmos.crisis.v1beta1.MsgVerifyInvariantResponse
}
var file_cosmos_crisis_v1beta1_tx_proto_depIdxs = []int32{
	0, // 0: cosmos.crisis.v1beta1.Msg.VerifyInvariant:input_type -> cosmos.crisis.v1beta1.MsgVerifyInvariant
	1, // 1: cosmos.crisis.v1beta1.Msg.VerifyInvariant:output_type -> cosmos.crisis.v1beta1.MsgVerifyInvariantResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_crisis_v1beta1_tx_proto_init() }
func file_cosmos_crisis_v1beta1_tx_proto_init() {
	if File_cosmos_crisis_v1beta1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crisis_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVerifyInvariant); i {
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
		file_cosmos_crisis_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVerifyInvariantResponse); i {
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
			RawDescriptor: file_cosmos_crisis_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_crisis_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_crisis_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_crisis_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_crisis_v1beta1_tx_proto = out.File
	file_cosmos_crisis_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_crisis_v1beta1_tx_proto_goTypes = nil
	file_cosmos_crisis_v1beta1_tx_proto_depIdxs = nil
}
