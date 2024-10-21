// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: auth.proto

package server

import (
	"context"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/logic/casbin"
	"xlife/apps/auth/rpc/internal/svc"
)

type CasbinServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedCasbinServer
}

func NewCasbinServer(svcCtx *svc.ServiceContext) *CasbinServer {
	return &CasbinServer{
		svcCtx: svcCtx,
	}
}

func (s *CasbinServer) Check(ctx context.Context, in *auth.CasbinCheckReq) (*auth.CasbinCheckResp, error) {
	l := casbinlogic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
