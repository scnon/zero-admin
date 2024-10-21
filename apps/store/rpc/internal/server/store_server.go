// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: store.proto

package server

import (
	"context"

	"xlife/apps/store/rpc/internal/logic"
	"xlife/apps/store/rpc/internal/svc"
	"xlife/apps/store/rpc/store"
)

type StoreServer struct {
	svcCtx *svc.ServiceContext
	store.UnimplementedStoreServer
}

func NewStoreServer(svcCtx *svc.ServiceContext) *StoreServer {
	return &StoreServer{
		svcCtx: svcCtx,
	}
}

func (s *StoreServer) AddStore(ctx context.Context, in *store.AddStoreReq) (*store.StoreInfo, error) {
	l := logic.NewAddStoreLogic(ctx, s.svcCtx)
	return l.AddStore(in)
}

func (s *StoreServer) UpdateStore(ctx context.Context, in *store.StoreInfo) (*store.StoreInfo, error) {
	l := logic.NewUpdateStoreLogic(ctx, s.svcCtx)
	return l.UpdateStore(in)
}

func (s *StoreServer) DeleteStore(ctx context.Context, in *store.DeleteStoreReq) (*store.DeleteStoreResp, error) {
	l := logic.NewDeleteStoreLogic(ctx, s.svcCtx)
	return l.DeleteStore(in)
}

func (s *StoreServer) GetStore(ctx context.Context, in *store.GetStoreReq) (*store.GetStoreResp, error) {
	l := logic.NewGetStoreLogic(ctx, s.svcCtx)
	return l.GetStore(in)
}
