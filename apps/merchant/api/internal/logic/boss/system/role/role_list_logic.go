package role

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	// 1. 获取角色列表
	res, err := l.svcCtx.Role.RoleList(l.ctx, &auth.RoleListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   &req.Status,
		TenantId: &l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	// 2. 数据转换
	var list []*types.RoleData
	for _, role := range res.List {
		list = append(list, &types.RoleData{
			Id:     role.Id,
			Name:   role.Name,
			Sort:   role.Sort,
			Status: role.Status,
			Remark: role.Remark,
		})
	}
	return &types.RoleListResp{
		Base: l.svcCtx.Success(),
		Data: types.RoleList{
			Total: res.Total,
			List:  list,
		},
	}, nil
}
