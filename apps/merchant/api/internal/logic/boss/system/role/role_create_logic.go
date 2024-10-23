package role

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleCreateLogic {
	return &RoleCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleCreateLogic) RoleCreate(req *types.RoleCreateReq) (resp *types.RoleCreateResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 1. 创建角色
	res, err := l.svcCtx.Role.AddRole(l.ctx, &auth.AddRoleReq{
		Name:     req.Name,
		Sort:     req.Sort,
		Status:   req.Status,
		Remark:   req.Remark,
		Op:       uid,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	return &types.RoleCreateResp{
		Base: l.svcCtx.Success(),
		Data: res.Id,
	}, nil
}
