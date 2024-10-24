// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: apps/merchant/rpc/merchant.proto

package merchant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MerchantClient is the client API for Merchant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantClient interface {
	AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...grpc.CallOption) (*MerchantInfo, error)
	UpdateMerchant(ctx context.Context, in *MerchantInfo, opts ...grpc.CallOption) (*MerchantInfo, error)
	DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, opts ...grpc.CallOption) (*DeleteMerchantResp, error)
	GetMerchant(ctx context.Context, in *MerchantListReq, opts ...grpc.CallOption) (*MerchantListResp, error)
}

type merchantClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantClient(cc grpc.ClientConnInterface) MerchantClient {
	return &merchantClient{cc}
}

func (c *merchantClient) AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...grpc.CallOption) (*MerchantInfo, error) {
	out := new(MerchantInfo)
	err := c.cc.Invoke(ctx, "/merchant.Merchant/AddMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) UpdateMerchant(ctx context.Context, in *MerchantInfo, opts ...grpc.CallOption) (*MerchantInfo, error) {
	out := new(MerchantInfo)
	err := c.cc.Invoke(ctx, "/merchant.Merchant/UpdateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, opts ...grpc.CallOption) (*DeleteMerchantResp, error) {
	out := new(DeleteMerchantResp)
	err := c.cc.Invoke(ctx, "/merchant.Merchant/DeleteMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) GetMerchant(ctx context.Context, in *MerchantListReq, opts ...grpc.CallOption) (*MerchantListResp, error) {
	out := new(MerchantListResp)
	err := c.cc.Invoke(ctx, "/merchant.Merchant/GetMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServer is the server API for Merchant service.
// All implementations must embed UnimplementedMerchantServer
// for forward compatibility
type MerchantServer interface {
	AddMerchant(context.Context, *AddMerchantReq) (*MerchantInfo, error)
	UpdateMerchant(context.Context, *MerchantInfo) (*MerchantInfo, error)
	DeleteMerchant(context.Context, *DeleteMerchantReq) (*DeleteMerchantResp, error)
	GetMerchant(context.Context, *MerchantListReq) (*MerchantListResp, error)
	mustEmbedUnimplementedMerchantServer()
}

// UnimplementedMerchantServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServer struct {
}

func (UnimplementedMerchantServer) AddMerchant(context.Context, *AddMerchantReq) (*MerchantInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMerchant not implemented")
}
func (UnimplementedMerchantServer) UpdateMerchant(context.Context, *MerchantInfo) (*MerchantInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMerchant not implemented")
}
func (UnimplementedMerchantServer) DeleteMerchant(context.Context, *DeleteMerchantReq) (*DeleteMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchant not implemented")
}
func (UnimplementedMerchantServer) GetMerchant(context.Context, *MerchantListReq) (*MerchantListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchant not implemented")
}
func (UnimplementedMerchantServer) mustEmbedUnimplementedMerchantServer() {}

// UnsafeMerchantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServer will
// result in compilation errors.
type UnsafeMerchantServer interface {
	mustEmbedUnimplementedMerchantServer()
}

func RegisterMerchantServer(s grpc.ServiceRegistrar, srv MerchantServer) {
	s.RegisterService(&Merchant_ServiceDesc, srv)
}

func _Merchant_AddMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).AddMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.Merchant/AddMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).AddMerchant(ctx, req.(*AddMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.Merchant/UpdateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).UpdateMerchant(ctx, req.(*MerchantInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_DeleteMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).DeleteMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.Merchant/DeleteMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).DeleteMerchant(ctx, req.(*DeleteMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_GetMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).GetMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.Merchant/GetMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).GetMerchant(ctx, req.(*MerchantListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Merchant_ServiceDesc is the grpc.ServiceDesc for Merchant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Merchant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "merchant.Merchant",
	HandlerType: (*MerchantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMerchant",
			Handler:    _Merchant_AddMerchant_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _Merchant_UpdateMerchant_Handler,
		},
		{
			MethodName: "DeleteMerchant",
			Handler:    _Merchant_DeleteMerchant_Handler,
		},
		{
			MethodName: "GetMerchant",
			Handler:    _Merchant_GetMerchant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/merchant/rpc/merchant.proto",
}
