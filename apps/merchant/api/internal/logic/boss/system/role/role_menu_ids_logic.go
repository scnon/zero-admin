package role

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuIdsLogic {
	return &RoleMenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleMenuIdsLogic) RoleMenuIds(req *types.RoleMenuIdsReq) (resp *types.RoleMenuIdsResp, err error) {
	res, err := l.svcCtx.Role.RoleMenuIds(l.ctx, &auth.RoleMenuIdsReq{
		RoleId:   req.RoleId,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	return &types.RoleMenuIdsResp{
		Base: l.svcCtx.Success(),
		Data: res.MenuIds,
	}, nil
}
