// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: tendermint/crypto/proof.proto

package crypto

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

type Proof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Index    int64    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	LeafHash []byte   `protobuf:"bytes,3,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
	Aunts    [][]byte `protobuf:"bytes,4,rep,name=aunts,proto3" json:"aunts,omitempty"`
}

func (x *Proof) Reset() {
	*x = Proof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_crypto_proof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_crypto_proof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_tendermint_crypto_proof_proto_rawDescGZIP(), []int{0}
}

func (x *Proof) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Proof) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Proof) GetLeafHash() []byte {
	if x != nil {
		return x.LeafHash
	}
	return nil
}

func (x *Proof) GetAunts() [][]byte {
	if x != nil {
		return x.Aunts
	}
	return nil
}

type ValueOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encoded in ProofOp.Key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// To encode in ProofOp.Data
	Proof *Proof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *ValueOp) Reset() {
	*x = ValueOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_crypto_proof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueOp) ProtoMessage() {}

func (x *ValueOp) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_crypto_proof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueOp.ProtoReflect.Descriptor instead.
func (*ValueOp) Descriptor() ([]byte, []int) {
	return file_tendermint_crypto_proof_proto_rawDescGZIP(), []int{1}
}

func (x *ValueOp) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ValueOp) GetProof() *Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

type DominoOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Input  string `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Output string `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *DominoOp) Reset() {
	*x = DominoOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_crypto_proof_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DominoOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DominoOp) ProtoMessage() {}

func (x *DominoOp) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_crypto_proof_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DominoOp.ProtoReflect.Descriptor instead.
func (*DominoOp) Descriptor() ([]byte, []int) {
	return file_tendermint_crypto_proof_proto_rawDescGZIP(), []int{2}
}

func (x *DominoOp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DominoOp) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *DominoOp) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

// ProofOp defines an operation used for calculating Merkle root
// The data could be arbitrary format, providing nessecary data
// for example neighbouring node hash
type ProofOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key  []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProofOp) Reset() {
	*x = ProofOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_crypto_proof_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofOp) ProtoMessage() {}

func (x *ProofOp) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_crypto_proof_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofOp.ProtoReflect.Descriptor instead.
func (*ProofOp) Descriptor() ([]byte, []int) {
	return file_tendermint_crypto_proof_proto_rawDescGZIP(), []int{3}
}

func (x *ProofOp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProofOp) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ProofOp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ProofOps is Merkle proof defined by the list of ProofOps
type ProofOps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ops []*ProofOp `protobuf:"bytes,1,rep,name=ops,proto3" json:"ops,omitempty"`
}

func (x *ProofOps) Reset() {
	*x = ProofOps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_crypto_proof_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofOps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofOps) ProtoMessage() {}

func (x *ProofOps) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_crypto_proof_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofOps.ProtoReflect.Descriptor instead.
func (*ProofOps) Descriptor() ([]byte, []int) {
	return file_tendermint_crypto_proof_proto_rawDescGZIP(), []int{4}
}

func (x *ProofOps) GetOps() []*ProofOp {
	if x != nil {
		return x.Ops
	}
	return nil
}

var File_tendermint_crypto_proof_proto protoreflect.FileDescriptor

var file_tendermint_crypto_proof_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x65, 0x61, 0x66, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x6c, 0x65, 0x61, 0x66, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x75, 0x6e, 0x74, 0x73,
	0x22, 0x4b, 0x0a, 0x07, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a,
	0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x4a, 0x0a,
	0x08, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x6f, 0x4f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x43, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x4f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4f, 0x70, 0x73, 0x12, 0x32, 0x0a, 0x03, 0x6f, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x4f, 0x70, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x03, 0x6f, 0x70, 0x73, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tendermint_crypto_proof_proto_rawDescOnce sync.Once
	file_tendermint_crypto_proof_proto_rawDescData = file_tendermint_crypto_proof_proto_rawDesc
)

func file_tendermint_crypto_proof_proto_rawDescGZIP() []byte {
	file_tendermint_crypto_proof_proto_rawDescOnce.Do(func() {
		file_tendermint_crypto_proof_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_crypto_proof_proto_rawDescData)
	})
	return file_tendermint_crypto_proof_proto_rawDescData
}

var file_tendermint_crypto_proof_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tendermint_crypto_proof_proto_goTypes = []interface{}{
	(*Proof)(nil),    // 0: tendermint.crypto.Proof
	(*ValueOp)(nil),  // 1: tendermint.crypto.ValueOp
	(*DominoOp)(nil), // 2: tendermint.crypto.DominoOp
	(*ProofOp)(nil),  // 3: tendermint.crypto.ProofOp
	(*ProofOps)(nil), // 4: tendermint.crypto.ProofOps
}
var file_tendermint_crypto_proof_proto_depIdxs = []int32{
	0, // 0: tendermint.crypto.ValueOp.proof:type_name -> tendermint.crypto.Proof
	3, // 1: tendermint.crypto.ProofOps.ops:type_name -> tendermint.crypto.ProofOp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tendermint_crypto_proof_proto_init() }
func file_tendermint_crypto_proof_proto_init() {
	if File_tendermint_crypto_proof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_crypto_proof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proof); i {
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
		file_tendermint_crypto_proof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueOp); i {
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
		file_tendermint_crypto_proof_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DominoOp); i {
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
		file_tendermint_crypto_proof_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofOp); i {
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
		file_tendermint_crypto_proof_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofOps); i {
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
			RawDescriptor: file_tendermint_crypto_proof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_crypto_proof_proto_goTypes,
		DependencyIndexes: file_tendermint_crypto_proof_proto_depIdxs,
		MessageInfos:      file_tendermint_crypto_proof_proto_msgTypes,
	}.Build()
	File_tendermint_crypto_proof_proto = out.File
	file_tendermint_crypto_proof_proto_rawDesc = nil
	file_tendermint_crypto_proof_proto_goTypes = nil
	file_tendermint_crypto_proof_proto_depIdxs = nil
}
