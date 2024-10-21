// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: apps/product/rpc/product.proto

package product

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

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	AddProduct(ctx context.Context, in *AddProductReq, opts ...grpc.CallOption) (*ProductInfo, error)
	UpdateProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*ProductInfo, error)
	DeleteProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*ProductInfo, error)
	GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error)
	AddCate(ctx context.Context, in *AddCateReq, opts ...grpc.CallOption) (*CateInfo, error)
	UpdateCate(ctx context.Context, in *CateInfo, opts ...grpc.CallOption) (*CateInfo, error)
	DeleteCate(ctx context.Context, in *DeleteCateReq, opts ...grpc.CallOption) (*DeleCateResp, error)
	GetCate(ctx context.Context, in *GetCateReq, opts ...grpc.CallOption) (*GetCateResp, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) AddProduct(ctx context.Context, in *AddProductReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/product.Product/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/product.Product/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/product.Product/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error) {
	out := new(GetProductResp)
	err := c.cc.Invoke(ctx, "/product.Product/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) AddCate(ctx context.Context, in *AddCateReq, opts ...grpc.CallOption) (*CateInfo, error) {
	out := new(CateInfo)
	err := c.cc.Invoke(ctx, "/product.Product/AddCate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateCate(ctx context.Context, in *CateInfo, opts ...grpc.CallOption) (*CateInfo, error) {
	out := new(CateInfo)
	err := c.cc.Invoke(ctx, "/product.Product/UpdateCate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteCate(ctx context.Context, in *DeleteCateReq, opts ...grpc.CallOption) (*DeleCateResp, error) {
	out := new(DeleCateResp)
	err := c.cc.Invoke(ctx, "/product.Product/DeleteCate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetCate(ctx context.Context, in *GetCateReq, opts ...grpc.CallOption) (*GetCateResp, error) {
	out := new(GetCateResp)
	err := c.cc.Invoke(ctx, "/product.Product/GetCate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	AddProduct(context.Context, *AddProductReq) (*ProductInfo, error)
	UpdateProduct(context.Context, *ProductInfo) (*ProductInfo, error)
	DeleteProduct(context.Context, *ProductInfo) (*ProductInfo, error)
	GetProduct(context.Context, *GetProductReq) (*GetProductResp, error)
	AddCate(context.Context, *AddCateReq) (*CateInfo, error)
	UpdateCate(context.Context, *CateInfo) (*CateInfo, error)
	DeleteCate(context.Context, *DeleteCateReq) (*DeleCateResp, error)
	GetCate(context.Context, *GetCateReq) (*GetCateResp, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) AddProduct(context.Context, *AddProductReq) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductServer) UpdateProduct(context.Context, *ProductInfo) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *ProductInfo) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) GetProduct(context.Context, *GetProductReq) (*GetProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServer) AddCate(context.Context, *AddCateReq) (*CateInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCate not implemented")
}
func (UnimplementedProductServer) UpdateCate(context.Context, *CateInfo) (*CateInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCate not implemented")
}
func (UnimplementedProductServer) DeleteCate(context.Context, *DeleteCateReq) (*DeleCateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCate not implemented")
}
func (UnimplementedProductServer) GetCate(context.Context, *GetCateReq) (*GetCateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCate not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).AddProduct(ctx, req.(*AddProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProduct(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct(ctx, req.(*GetProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_AddCate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).AddCate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/AddCate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).AddCate(ctx, req.(*AddCateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateCate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateCate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/UpdateCate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateCate(ctx, req.(*CateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteCate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteCate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/DeleteCate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteCate(ctx, req.(*DeleteCateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetCate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetCate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/GetCate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetCate(ctx, req.(*GetCateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _Product_AddProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Product_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Product_GetProduct_Handler,
		},
		{
			MethodName: "AddCate",
			Handler:    _Product_AddCate_Handler,
		},
		{
			MethodName: "UpdateCate",
			Handler:    _Product_UpdateCate_Handler,
		},
		{
			MethodName: "DeleteCate",
			Handler:    _Product_DeleteCate_Handler,
		},
		{
			MethodName: "GetCate",
			Handler:    _Product_GetCate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/product/rpc/product.proto",
}
