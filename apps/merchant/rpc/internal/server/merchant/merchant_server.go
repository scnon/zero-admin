// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: merchant.proto

package server

import (
	"context"

	"xlife/apps/merchant/rpc/internal/logic/merchant"
	"xlife/apps/merchant/rpc/internal/svc"
	"xlife/apps/merchant/rpc/merchant"
)

type MerchantServer struct {
	svcCtx *svc.ServiceContext
	merchant.UnimplementedMerchantServer
}

func NewMerchantServer(svcCtx *svc.ServiceContext) *MerchantServer {
	return &MerchantServer{
		svcCtx: svcCtx,
	}
}

func (s *MerchantServer) AddMerchant(ctx context.Context, in *merchant.AddMerchantReq) (*merchant.MerchantInfo, error) {
	l := merchantlogic.NewAddMerchantLogic(ctx, s.svcCtx)
	return l.AddMerchant(in)
}

func (s *MerchantServer) UpdateMerchant(ctx context.Context, in *merchant.MerchantInfo) (*merchant.MerchantInfo, error) {
	l := merchantlogic.NewUpdateMerchantLogic(ctx, s.svcCtx)
	return l.UpdateMerchant(in)
}

func (s *MerchantServer) DeleteMerchant(ctx context.Context, in *merchant.DeleteMerchantReq) (*merchant.DeleteMerchantResp, error) {
	l := merchantlogic.NewDeleteMerchantLogic(ctx, s.svcCtx)
	return l.DeleteMerchant(in)
}

func (s *MerchantServer) GetMerchant(ctx context.Context, in *merchant.MerchantListReq) (*merchant.MerchantListResp, error) {
	l := merchantlogic.NewGetMerchantLogic(ctx, s.svcCtx)
	return l.GetMerchant(in)
}
