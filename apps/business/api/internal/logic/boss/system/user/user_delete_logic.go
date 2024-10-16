package user

import (
	"context"
	"xlife/apps/auth/rpc/admin"

	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"

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
	_, err = l.svcCtx.User.DeleteUser(l.ctx, &admin.DeleteUserReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserDeleteResp{
		Base: l.svcCtx.Success(),
	}, nil

}
