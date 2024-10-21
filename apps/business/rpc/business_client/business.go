// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: business.proto

package business_client

import (
	"context"

	"xlife/apps/business/rpc/business"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddBusinessReq     = business.AddBusinessReq
	BusinessInfo       = business.BusinessInfo
	DeleteBusinessReq  = business.DeleteBusinessReq
	DeleteBusinessResp = business.DeleteBusinessResp
	GetBusinessReq     = business.GetBusinessReq
	GetBusinessResp    = business.GetBusinessResp

	Business interface {
		AddBusiness(ctx context.Context, in *AddBusinessReq, opts ...grpc.CallOption) (*BusinessInfo, error)
		UpdateBusiness(ctx context.Context, in *BusinessInfo, opts ...grpc.CallOption) (*BusinessInfo, error)
		DeleteBusiness(ctx context.Context, in *DeleteBusinessReq, opts ...grpc.CallOption) (*DeleteBusinessResp, error)
		GetBusiness(ctx context.Context, in *GetBusinessReq, opts ...grpc.CallOption) (*GetBusinessResp, error)
	}

	defaultBusiness struct {
		cli zrpc.Client
	}
)

func NewBusiness(cli zrpc.Client) Business {
	return &defaultBusiness{
		cli: cli,
	}
}

func (m *defaultBusiness) AddBusiness(ctx context.Context, in *AddBusinessReq, opts ...grpc.CallOption) (*BusinessInfo, error) {
	client := business.NewBusinessClient(m.cli.Conn())
	return client.AddBusiness(ctx, in, opts...)
}

func (m *defaultBusiness) UpdateBusiness(ctx context.Context, in *BusinessInfo, opts ...grpc.CallOption) (*BusinessInfo, error) {
	client := business.NewBusinessClient(m.cli.Conn())
	return client.UpdateBusiness(ctx, in, opts...)
}

func (m *defaultBusiness) DeleteBusiness(ctx context.Context, in *DeleteBusinessReq, opts ...grpc.CallOption) (*DeleteBusinessResp, error) {
	client := business.NewBusinessClient(m.cli.Conn())
	return client.DeleteBusiness(ctx, in, opts...)
}

func (m *defaultBusiness) GetBusiness(ctx context.Context, in *GetBusinessReq, opts ...grpc.CallOption) (*GetBusinessResp, error) {
	client := business.NewBusinessClient(m.cli.Conn())
	return client.GetBusiness(ctx, in, opts...)
}