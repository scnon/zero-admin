// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: auth.proto

package server

import (
	"context"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/logic/menu"
	"xlife/apps/auth/rpc/internal/svc"
)

type MenuServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedMenuServer
}

func NewMenuServer(svcCtx *svc.ServiceContext) *MenuServer {
	return &MenuServer{
		svcCtx: svcCtx,
	}
}

func (s *MenuServer) AddMenu(ctx context.Context, in *auth.AddMenuReq) (*auth.AddMenuResp, error) {
	l := menulogic.NewAddMenuLogic(ctx, s.svcCtx)
	return l.AddMenu(in)
}

func (s *MenuServer) DeleteMenu(ctx context.Context, in *auth.DeleteMenuReq) (*auth.DeleteMenuResp, error) {
	l := menulogic.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *MenuServer) UpdateMenu(ctx context.Context, in *auth.UpdateMenuReq) (*auth.UpdateMenuResp, error) {
	l := menulogic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *MenuServer) MenuList(ctx context.Context, in *auth.MenuListReq) (*auth.MenuListResp, error) {
	l := menulogic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

func (s *MenuServer) GetMenu(ctx context.Context, in *auth.GetMenuReq) (*auth.GetMenuResp, error) {
	l := menulogic.NewGetMenuLogic(ctx, s.svcCtx)
	return l.GetMenu(in)
}
