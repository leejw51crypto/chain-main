// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: tendermint/privval/types.proto

package privval

import (
	proto "github.com/golang/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
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

type Errors int32

const (
	Errors_ERRORS_UNKNOWN             Errors = 0
	Errors_ERRORS_UNEXPECTED_RESPONSE Errors = 1
	Errors_ERRORS_NO_CONNECTION       Errors = 2
	Errors_ERRORS_CONNECTION_TIMEOUT  Errors = 3
	Errors_ERRORS_READ_TIMEOUT        Errors = 4
	Errors_ERRORS_WRITE_TIMEOUT       Errors = 5
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0: "ERRORS_UNKNOWN",
		1: "ERRORS_UNEXPECTED_RESPONSE",
		2: "ERRORS_NO_CONNECTION",
		3: "ERRORS_CONNECTION_TIMEOUT",
		4: "ERRORS_READ_TIMEOUT",
		5: "ERRORS_WRITE_TIMEOUT",
	}
	Errors_value = map[string]int32{
		"ERRORS_UNKNOWN":             0,
		"ERRORS_UNEXPECTED_RESPONSE": 1,
		"ERRORS_NO_CONNECTION":       2,
		"ERRORS_CONNECTION_TIMEOUT":  3,
		"ERRORS_READ_TIMEOUT":        4,
		"ERRORS_WRITE_TIMEOUT":       5,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_tendermint_privval_types_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_tendermint_privval_types_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{0}
}

type RemoteSignerError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RemoteSignerError) Reset() {
	*x = RemoteSignerError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteSignerError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteSignerError) ProtoMessage() {}

func (x *RemoteSignerError) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteSignerError.ProtoReflect.Descriptor instead.
func (*RemoteSignerError) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteSignerError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RemoteSignerError) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// PubKeyRequest requests the consensus public key from the remote signer.
type PubKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *PubKeyRequest) Reset() {
	*x = PubKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeyRequest) ProtoMessage() {}

func (x *PubKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeyRequest.ProtoReflect.Descriptor instead.
func (*PubKeyRequest) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{1}
}

func (x *PubKeyRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

// PubKeyResponse is a response message containing the public key.
type PubKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey *crypto.PublicKey  `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Error  *RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PubKeyResponse) Reset() {
	*x = PubKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeyResponse) ProtoMessage() {}

func (x *PubKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeyResponse.ProtoReflect.Descriptor instead.
func (*PubKeyResponse) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{2}
}

func (x *PubKeyResponse) GetPubKey() *crypto.PublicKey {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *PubKeyResponse) GetError() *RemoteSignerError {
	if x != nil {
		return x.Error
	}
	return nil
}

// SignVoteRequest is a request to sign a vote
type SignVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote    *types.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	ChainId string      `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *SignVoteRequest) Reset() {
	*x = SignVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignVoteRequest) ProtoMessage() {}

func (x *SignVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignVoteRequest.ProtoReflect.Descriptor instead.
func (*SignVoteRequest) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{3}
}

func (x *SignVoteRequest) GetVote() *types.Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *SignVoteRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

// SignedVoteResponse is a response containing a signed vote or an error
type SignedVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote  *types.Vote        `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Error *RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SignedVoteResponse) Reset() {
	*x = SignedVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedVoteResponse) ProtoMessage() {}

func (x *SignedVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedVoteResponse.ProtoReflect.Descriptor instead.
func (*SignedVoteResponse) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{4}
}

func (x *SignedVoteResponse) GetVote() *types.Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *SignedVoteResponse) GetError() *RemoteSignerError {
	if x != nil {
		return x.Error
	}
	return nil
}

// SignProposalRequest is a request to sign a proposal
type SignProposalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposal *types.Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	ChainId  string          `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *SignProposalRequest) Reset() {
	*x = SignProposalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignProposalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignProposalRequest) ProtoMessage() {}

func (x *SignProposalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignProposalRequest.ProtoReflect.Descriptor instead.
func (*SignProposalRequest) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{5}
}

func (x *SignProposalRequest) GetProposal() *types.Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

func (x *SignProposalRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

// SignedProposalResponse is response containing a signed proposal or an error
type SignedProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposal *types.Proposal    `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	Error    *RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SignedProposalResponse) Reset() {
	*x = SignedProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedProposalResponse) ProtoMessage() {}

func (x *SignedProposalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedProposalResponse.ProtoReflect.Descriptor instead.
func (*SignedProposalResponse) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{6}
}

func (x *SignedProposalResponse) GetProposal() *types.Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

func (x *SignedProposalResponse) GetError() *RemoteSignerError {
	if x != nil {
		return x.Error
	}
	return nil
}

// PingRequest is a request to confirm that the connection is alive.
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{7}
}

// PingResponse is a response to confirm that the connection is alive.
type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{8}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sum:
	//	*Message_PubKeyRequest
	//	*Message_PubKeyResponse
	//	*Message_SignVoteRequest
	//	*Message_SignedVoteResponse
	//	*Message_SignProposalRequest
	//	*Message_SignedProposalResponse
	//	*Message_PingRequest
	//	*Message_PingResponse
	Sum isMessage_Sum `protobuf_oneof:"sum"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_privval_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_privval_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_tendermint_privval_types_proto_rawDescGZIP(), []int{9}
}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (x *Message) GetPubKeyRequest() *PubKeyRequest {
	if x, ok := x.GetSum().(*Message_PubKeyRequest); ok {
		return x.PubKeyRequest
	}
	return nil
}

func (x *Message) GetPubKeyResponse() *PubKeyResponse {
	if x, ok := x.GetSum().(*Message_PubKeyResponse); ok {
		return x.PubKeyResponse
	}
	return nil
}

func (x *Message) GetSignVoteRequest() *SignVoteRequest {
	if x, ok := x.GetSum().(*Message_SignVoteRequest); ok {
		return x.SignVoteRequest
	}
	return nil
}

func (x *Message) GetSignedVoteResponse() *SignedVoteResponse {
	if x, ok := x.GetSum().(*Message_SignedVoteResponse); ok {
		return x.SignedVoteResponse
	}
	return nil
}

func (x *Message) GetSignProposalRequest() *SignProposalRequest {
	if x, ok := x.GetSum().(*Message_SignProposalRequest); ok {
		return x.SignProposalRequest
	}
	return nil
}

func (x *Message) GetSignedProposalResponse() *SignedProposalResponse {
	if x, ok := x.GetSum().(*Message_SignedProposalResponse); ok {
		return x.SignedProposalResponse
	}
	return nil
}

func (x *Message) GetPingRequest() *PingRequest {
	if x, ok := x.GetSum().(*Message_PingRequest); ok {
		return x.PingRequest
	}
	return nil
}

func (x *Message) GetPingResponse() *PingResponse {
	if x, ok := x.GetSum().(*Message_PingResponse); ok {
		return x.PingResponse
	}
	return nil
}

type isMessage_Sum interface {
	isMessage_Sum()
}

type Message_PubKeyRequest struct {
	PubKeyRequest *PubKeyRequest `protobuf:"bytes,1,opt,name=pub_key_request,json=pubKeyRequest,proto3,oneof"`
}

type Message_PubKeyResponse struct {
	PubKeyResponse *PubKeyResponse `protobuf:"bytes,2,opt,name=pub_key_response,json=pubKeyResponse,proto3,oneof"`
}

type Message_SignVoteRequest struct {
	SignVoteRequest *SignVoteRequest `protobuf:"bytes,3,opt,name=sign_vote_request,json=signVoteRequest,proto3,oneof"`
}

type Message_SignedVoteResponse struct {
	SignedVoteResponse *SignedVoteResponse `protobuf:"bytes,4,opt,name=signed_vote_response,json=signedVoteResponse,proto3,oneof"`
}

type Message_SignProposalRequest struct {
	SignProposalRequest *SignProposalRequest `protobuf:"bytes,5,opt,name=sign_proposal_request,json=signProposalRequest,proto3,oneof"`
}

type Message_SignedProposalResponse struct {
	SignedProposalResponse *SignedProposalResponse `protobuf:"bytes,6,opt,name=signed_proposal_response,json=signedProposalResponse,proto3,oneof"`
}

type Message_PingRequest struct {
	PingRequest *PingRequest `protobuf:"bytes,7,opt,name=ping_request,json=pingRequest,proto3,oneof"`
}

type Message_PingResponse struct {
	PingResponse *PingResponse `protobuf:"bytes,8,opt,name=ping_response,json=pingResponse,proto3,oneof"`
}

func (*Message_PubKeyRequest) isMessage_Sum() {}

func (*Message_PubKeyResponse) isMessage_Sum() {}

func (*Message_SignVoteRequest) isMessage_Sum() {}

func (*Message_SignedVoteResponse) isMessage_Sum() {}

func (*Message_SignProposalRequest) isMessage_Sum() {}

func (*Message_SignedProposalResponse) isMessage_Sum() {}

func (*Message_PingRequest) isMessage_Sum() {}

func (*Message_PingResponse) isMessage_Sum() {}

var File_tendermint_privval_types_proto protoreflect.FileDescriptor

var file_tendermint_privval_types_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x76, 0x61, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x76, 0x61, 0x6c, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x0d, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x75,
	0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x58,
	0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x68, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x22, 0x8d, 0x01, 0x0a, 0x16, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xb2, 0x05, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x10, 0x70, 0x75, 0x62,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x69, 0x67,
	0x6e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x14,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x15, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x18, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x16, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0d, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61,
	0x6c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x05,
	0x0a, 0x03, 0x73, 0x75, 0x6d, 0x2a, 0xa8, 0x01, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x55,
	0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x4e,
	0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x1d,
	0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a,
	0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53,
	0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x05,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x76, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_privval_types_proto_rawDescOnce sync.Once
	file_tendermint_privval_types_proto_rawDescData = file_tendermint_privval_types_proto_rawDesc
)

func file_tendermint_privval_types_proto_rawDescGZIP() []byte {
	file_tendermint_privval_types_proto_rawDescOnce.Do(func() {
		file_tendermint_privval_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_privval_types_proto_rawDescData)
	})
	return file_tendermint_privval_types_proto_rawDescData
}

var file_tendermint_privval_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tendermint_privval_types_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tendermint_privval_types_proto_goTypes = []interface{}{
	(Errors)(0),                    // 0: tendermint.privval.Errors
	(*RemoteSignerError)(nil),      // 1: tendermint.privval.RemoteSignerError
	(*PubKeyRequest)(nil),          // 2: tendermint.privval.PubKeyRequest
	(*PubKeyResponse)(nil),         // 3: tendermint.privval.PubKeyResponse
	(*SignVoteRequest)(nil),        // 4: tendermint.privval.SignVoteRequest
	(*SignedVoteResponse)(nil),     // 5: tendermint.privval.SignedVoteResponse
	(*SignProposalRequest)(nil),    // 6: tendermint.privval.SignProposalRequest
	(*SignedProposalResponse)(nil), // 7: tendermint.privval.SignedProposalResponse
	(*PingRequest)(nil),            // 8: tendermint.privval.PingRequest
	(*PingResponse)(nil),           // 9: tendermint.privval.PingResponse
	(*Message)(nil),                // 10: tendermint.privval.Message
	(*crypto.PublicKey)(nil),       // 11: tendermint.crypto.PublicKey
	(*types.Vote)(nil),             // 12: tendermint.types.Vote
	(*types.Proposal)(nil),         // 13: tendermint.types.Proposal
}
var file_tendermint_privval_types_proto_depIdxs = []int32{
	11, // 0: tendermint.privval.PubKeyResponse.pub_key:type_name -> tendermint.crypto.PublicKey
	1,  // 1: tendermint.privval.PubKeyResponse.error:type_name -> tendermint.privval.RemoteSignerError
	12, // 2: tendermint.privval.SignVoteRequest.vote:type_name -> tendermint.types.Vote
	12, // 3: tendermint.privval.SignedVoteResponse.vote:type_name -> tendermint.types.Vote
	1,  // 4: tendermint.privval.SignedVoteResponse.error:type_name -> tendermint.privval.RemoteSignerError
	13, // 5: tendermint.privval.SignProposalRequest.proposal:type_name -> tendermint.types.Proposal
	13, // 6: tendermint.privval.SignedProposalResponse.proposal:type_name -> tendermint.types.Proposal
	1,  // 7: tendermint.privval.SignedProposalResponse.error:type_name -> tendermint.privval.RemoteSignerError
	2,  // 8: tendermint.privval.Message.pub_key_request:type_name -> tendermint.privval.PubKeyRequest
	3,  // 9: tendermint.privval.Message.pub_key_response:type_name -> tendermint.privval.PubKeyResponse
	4,  // 10: tendermint.privval.Message.sign_vote_request:type_name -> tendermint.privval.SignVoteRequest
	5,  // 11: tendermint.privval.Message.signed_vote_response:type_name -> tendermint.privval.SignedVoteResponse
	6,  // 12: tendermint.privval.Message.sign_proposal_request:type_name -> tendermint.privval.SignProposalRequest
	7,  // 13: tendermint.privval.Message.signed_proposal_response:type_name -> tendermint.privval.SignedProposalResponse
	8,  // 14: tendermint.privval.Message.ping_request:type_name -> tendermint.privval.PingRequest
	9,  // 15: tendermint.privval.Message.ping_response:type_name -> tendermint.privval.PingResponse
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_tendermint_privval_types_proto_init() }
func file_tendermint_privval_types_proto_init() {
	if File_tendermint_privval_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_privval_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteSignerError); i {
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
		file_tendermint_privval_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeyRequest); i {
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
		file_tendermint_privval_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeyResponse); i {
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
		file_tendermint_privval_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignVoteRequest); i {
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
		file_tendermint_privval_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedVoteResponse); i {
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
		file_tendermint_privval_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignProposalRequest); i {
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
		file_tendermint_privval_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedProposalResponse); i {
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
		file_tendermint_privval_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_tendermint_privval_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_tendermint_privval_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_tendermint_privval_types_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Message_PubKeyRequest)(nil),
		(*Message_PubKeyResponse)(nil),
		(*Message_SignVoteRequest)(nil),
		(*Message_SignedVoteResponse)(nil),
		(*Message_SignProposalRequest)(nil),
		(*Message_SignedProposalResponse)(nil),
		(*Message_PingRequest)(nil),
		(*Message_PingResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_privval_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_privval_types_proto_goTypes,
		DependencyIndexes: file_tendermint_privval_types_proto_depIdxs,
		EnumInfos:         file_tendermint_privval_types_proto_enumTypes,
		MessageInfos:      file_tendermint_privval_types_proto_msgTypes,
	}.Build()
	File_tendermint_privval_types_proto = out.File
	file_tendermint_privval_types_proto_rawDesc = nil
	file_tendermint_privval_types_proto_goTypes = nil
	file_tendermint_privval_types_proto_depIdxs = nil
}
