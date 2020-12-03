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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries params of the distribution module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ValidatorOutstandingRewards queries rewards of a validator address.
	ValidatorOutstandingRewards(ctx context.Context, in *QueryValidatorOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryValidatorOutstandingRewardsResponse, error)
	// ValidatorCommission queries accumulated commission for a validator.
	ValidatorCommission(ctx context.Context, in *QueryValidatorCommissionRequest, opts ...grpc.CallOption) (*QueryValidatorCommissionResponse, error)
	// ValidatorSlashes queries slash events of a validator.
	ValidatorSlashes(ctx context.Context, in *QueryValidatorSlashesRequest, opts ...grpc.CallOption) (*QueryValidatorSlashesResponse, error)
	// DelegationRewards queries the total rewards accrued by a delegation.
	DelegationRewards(ctx context.Context, in *QueryDelegationRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationRewardsResponse, error)
	// DelegationTotalRewards queries the total rewards accrued by a each
	// validator.
	DelegationTotalRewards(ctx context.Context, in *QueryDelegationTotalRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationTotalRewardsResponse, error)
	// DelegatorValidators queries the validators of a delegator.
	DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorWithdrawAddress queries withdraw address of a delegator.
	DelegatorWithdrawAddress(ctx context.Context, in *QueryDelegatorWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryDelegatorWithdrawAddressResponse, error)
	// CommunityPool queries the community pool coins.
	CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorOutstandingRewards(ctx context.Context, in *QueryValidatorOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryValidatorOutstandingRewardsResponse, error) {
	out := new(QueryValidatorOutstandingRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorCommission(ctx context.Context, in *QueryValidatorCommissionRequest, opts ...grpc.CallOption) (*QueryValidatorCommissionResponse, error) {
	out := new(QueryValidatorCommissionResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorSlashes(ctx context.Context, in *QueryValidatorSlashesRequest, opts ...grpc.CallOption) (*QueryValidatorSlashesResponse, error) {
	out := new(QueryValidatorSlashesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorSlashes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationRewards(ctx context.Context, in *QueryDelegationRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationRewardsResponse, error) {
	out := new(QueryDelegationRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegationRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationTotalRewards(ctx context.Context, in *QueryDelegationTotalRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationTotalRewardsResponse, error) {
	out := new(QueryDelegationTotalRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error) {
	out := new(QueryDelegatorValidatorsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegatorValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorWithdrawAddress(ctx context.Context, in *QueryDelegatorWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryDelegatorWithdrawAddressResponse, error) {
	out := new(QueryDelegatorWithdrawAddressResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error) {
	out := new(QueryCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/CommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params queries params of the distribution module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ValidatorOutstandingRewards queries rewards of a validator address.
	ValidatorOutstandingRewards(context.Context, *QueryValidatorOutstandingRewardsRequest) (*QueryValidatorOutstandingRewardsResponse, error)
	// ValidatorCommission queries accumulated commission for a validator.
	ValidatorCommission(context.Context, *QueryValidatorCommissionRequest) (*QueryValidatorCommissionResponse, error)
	// ValidatorSlashes queries slash events of a validator.
	ValidatorSlashes(context.Context, *QueryValidatorSlashesRequest) (*QueryValidatorSlashesResponse, error)
	// DelegationRewards queries the total rewards accrued by a delegation.
	DelegationRewards(context.Context, *QueryDelegationRewardsRequest) (*QueryDelegationRewardsResponse, error)
	// DelegationTotalRewards queries the total rewards accrued by a each
	// validator.
	DelegationTotalRewards(context.Context, *QueryDelegationTotalRewardsRequest) (*QueryDelegationTotalRewardsResponse, error)
	// DelegatorValidators queries the validators of a delegator.
	DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorWithdrawAddress queries withdraw address of a delegator.
	DelegatorWithdrawAddress(context.Context, *QueryDelegatorWithdrawAddressRequest) (*QueryDelegatorWithdrawAddressResponse, error)
	// CommunityPool queries the community pool coins.
	CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ValidatorOutstandingRewards(context.Context, *QueryValidatorOutstandingRewardsRequest) (*QueryValidatorOutstandingRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorOutstandingRewards not implemented")
}
func (UnimplementedQueryServer) ValidatorCommission(context.Context, *QueryValidatorCommissionRequest) (*QueryValidatorCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorCommission not implemented")
}
func (UnimplementedQueryServer) ValidatorSlashes(context.Context, *QueryValidatorSlashesRequest) (*QueryValidatorSlashesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSlashes not implemented")
}
func (UnimplementedQueryServer) DelegationRewards(context.Context, *QueryDelegationRewardsRequest) (*QueryDelegationRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationRewards not implemented")
}
func (UnimplementedQueryServer) DelegationTotalRewards(context.Context, *QueryDelegationTotalRewardsRequest) (*QueryDelegationTotalRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationTotalRewards not implemented")
}
func (UnimplementedQueryServer) DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorValidators not implemented")
}
func (UnimplementedQueryServer) DelegatorWithdrawAddress(context.Context, *QueryDelegatorWithdrawAddressRequest) (*QueryDelegatorWithdrawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorWithdrawAddress not implemented")
}
func (UnimplementedQueryServer) CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPool not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorOutstandingRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorOutstandingRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, req.(*QueryValidatorOutstandingRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorCommission(ctx, req.(*QueryValidatorCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorSlashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorSlashesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSlashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorSlashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSlashes(ctx, req.(*QueryValidatorSlashesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationRewards(ctx, req.(*QueryDelegationRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationTotalRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationTotalRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationTotalRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationTotalRewards(ctx, req.(*QueryDelegationTotalRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(ctx, req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorWithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorWithdrawAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, req.(*QueryDelegatorWithdrawAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/CommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommunityPool(ctx, req.(*QueryCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.distribution.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ValidatorOutstandingRewards",
			Handler:    _Query_ValidatorOutstandingRewards_Handler,
		},
		{
			MethodName: "ValidatorCommission",
			Handler:    _Query_ValidatorCommission_Handler,
		},
		{
			MethodName: "ValidatorSlashes",
			Handler:    _Query_ValidatorSlashes_Handler,
		},
		{
			MethodName: "DelegationRewards",
			Handler:    _Query_DelegationRewards_Handler,
		},
		{
			MethodName: "DelegationTotalRewards",
			Handler:    _Query_DelegationTotalRewards_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorWithdrawAddress",
			Handler:    _Query_DelegatorWithdrawAddress_Handler,
		},
		{
			MethodName: "CommunityPool",
			Handler:    _Query_CommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/distribution/v1beta1/query.proto",
}
