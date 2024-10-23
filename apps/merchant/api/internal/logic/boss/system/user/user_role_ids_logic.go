package user

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleIdsLogic {
	return &UserRoleIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleIdsLogic) UserRoleIds(req *types.UserRoleIdsReq) (resp *types.UserRoleIdsResp, err error) {
	// 1. 获取当前用户
	//uid := ctxdata.GetUId(l.ctx)
	// 2. 获取用户角色
	res, err := l.svcCtx.User.UserRoleIds(l.ctx, &auth.UserRoleIdsReq{
		UserId:   req.UserId,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserRoleIdsResp{
		Base: l.svcCtx.Success(),
		Data: res.RoleIds,
	}, nil
}
