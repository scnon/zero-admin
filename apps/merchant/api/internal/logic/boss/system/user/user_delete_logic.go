package user

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.UserDeleteReq) (resp *types.UserDeleteResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 删除用户(逻辑删除)
	if _, err = l.svcCtx.User.DeleteUser(l.ctx, &auth.DeleteUserReq{
		Ids:      req.Ids,
		TenantId: l.svcCtx.Config.Tenant,
		Op:       uid,
	}); err != nil {
		return nil, err
	}
	return &types.UserDeleteResp{
		Base: l.svcCtx.Success(),
	}, nil
}
