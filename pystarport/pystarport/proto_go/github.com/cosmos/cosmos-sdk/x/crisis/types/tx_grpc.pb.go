// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// VerifyInvariant defines a method to verify a particular invariance.
	VerifyInvariant(ctx context.Context, in *MsgVerifyInvariant, opts ...grpc.CallOption) (*MsgVerifyInvariantResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) VerifyInvariant(ctx context.Context, in *MsgVerifyInvariant, opts ...grpc.CallOption) (*MsgVerifyInvariantResponse, error) {
	out := new(MsgVerifyInvariantResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1beta1.Msg/VerifyInvariant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// VerifyInvariant defines a method to verify a particular invariance.
	VerifyInvariant(context.Context, *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) VerifyInvariant(context.Context, *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyInvariant not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_VerifyInvariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVerifyInvariant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VerifyInvariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1beta1.Msg/VerifyInvariant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VerifyInvariant(ctx, req.(*MsgVerifyInvariant))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.crisis.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyInvariant",
			Handler:    _Msg_VerifyInvariant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/crisis/v1beta1/tx.proto",
}
