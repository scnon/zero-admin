package role

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.RoleDeleteReq) (resp *types.RoleDeleteResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 删除角色
	if _, err = l.svcCtx.Role.DeleteRole(l.ctx, &auth.DeleteRoleReq{
		Ids:      req.Ids,
		Op:       uid,
		TenantId: l.svcCtx.Config.Tenant,
	}); err != nil {
		return nil, err
	}
	return &types.RoleDeleteResp{
		Base: l.svcCtx.Success(),
	}, nil
}
