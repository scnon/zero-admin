// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: auth.proto

package role

import (
	"context"

	"xlife/apps/auth/rpc/auth"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddDeptReq         = auth.AddDeptReq
	AddDeptResp        = auth.AddDeptResp
	AddMenuReq         = auth.AddMenuReq
	AddMenuResp        = auth.AddMenuResp
	AddRoleReq         = auth.AddRoleReq
	AddRoleResp        = auth.AddRoleResp
	AddUserReq         = auth.AddUserReq
	AddUserResp        = auth.AddUserResp
	AssignRoleMenuReq  = auth.AssignRoleMenuReq
	AssignRoleMenuResp = auth.AssignRoleMenuResp
	AssignUserRoleReq  = auth.AssignUserRoleReq
	AssignUserRoleResp = auth.AssignUserRoleResp
	CasbinCheckReq     = auth.CasbinCheckReq
	CasbinCheckResp    = auth.CasbinCheckResp
	DeleteDeptReq      = auth.DeleteDeptReq
	DeleteDeptResp     = auth.DeleteDeptResp
	DeleteMenuReq      = auth.DeleteMenuReq
	DeleteMenuResp     = auth.DeleteMenuResp
	DeleteRoleReq      = auth.DeleteRoleReq
	DeleteRoleResp     = auth.DeleteRoleResp
	DeleteUserReq      = auth.DeleteUserReq
	DeleteUserResp     = auth.DeleteUserResp
	DeptData           = auth.DeptData
	DeptListReq        = auth.DeptListReq
	DeptListResp       = auth.DeptListResp
	GetMenuReq         = auth.GetMenuReq
	GetMenuResp        = auth.GetMenuResp
	LoginReq           = auth.LoginReq
	LoginResp          = auth.LoginResp
	MenuData           = auth.MenuData
	MenuListReq        = auth.MenuListReq
	MenuListResp       = auth.MenuListResp
	RefreshReq         = auth.RefreshReq
	ResetPasswordReq   = auth.ResetPasswordReq
	ResetPasswordResp  = auth.ResetPasswordResp
	RoleData           = auth.RoleData
	RoleListReq        = auth.RoleListReq
	RoleListResp       = auth.RoleListResp
	RoleMenuIdsReq     = auth.RoleMenuIdsReq
	RoleMenuIdsResp    = auth.RoleMenuIdsResp
	UpdateDeptReq      = auth.UpdateDeptReq
	UpdateDeptResp     = auth.UpdateDeptResp
	UpdateMenuReq      = auth.UpdateMenuReq
	UpdateMenuResp     = auth.UpdateMenuResp
	UpdateRoleReq      = auth.UpdateRoleReq
	UpdateRoleResp     = auth.UpdateRoleResp
	UpdateUserReq      = auth.UpdateUserReq
	UpdateUserResp     = auth.UpdateUserResp
	UserData           = auth.UserData
	UserListReq        = auth.UserListReq
	UserListResp       = auth.UserListResp
	UserRoleIdsReq     = auth.UserRoleIdsReq
	UserRoleIdsResp    = auth.UserRoleIdsResp

	Role interface {
		AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error)
		DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error)
		UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		AssignRoleMenu(ctx context.Context, in *AssignRoleMenuReq, opts ...grpc.CallOption) (*AssignRoleMenuResp, error)
		RoleMenuIds(ctx context.Context, in *RoleMenuIdsReq, opts ...grpc.CallOption) (*RoleMenuIdsResp, error)
	}

	defaultRole struct {
		cli zrpc.Client
	}
)

func NewRole(cli zrpc.Client) Role {
	return &defaultRole{
		cli: cli,
	}
}

func (m *defaultRole) AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error) {
	client := auth.NewRoleClient(m.cli.Conn())
	return client.AddRole(ctx, in, opts...)
}

func (m *defaultRole) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error) {
	client := auth.NewRoleClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultRole) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error) {
	client := auth.NewRoleClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultRole) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := auth.NewRoleClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultRole) AssignRoleMenu(ctx context.Context, in *AssignRoleMenuReq, opts ...grpc.CallOption) (*AssignRoleMenuResp, error) {
	client := auth.NewRoleClient(m.cli.Conn())
	return client.AssignRoleMenu(ctx, in, opts...)
}

func (m *defaultRole) RoleMenuIds(ctx context.Context, in *RoleMenuIdsReq, opts ...grpc.CallOption) (*RoleMenuIdsResp, error) {
	client := auth.NewRoleClient(m.cli.Conn())
	return client.RoleMenuIds(ctx, in, opts...)
}
