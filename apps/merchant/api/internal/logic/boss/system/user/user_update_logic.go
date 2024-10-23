package user

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	_, err = l.svcCtx.User.UpdateUser(l.ctx, &auth.UpdateUserReq{
		Id:       req.Id,
		Username: req.Username,
		Nickname: req.NickName,
		Sort:     req.Sort,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserUpdateResp{
		Base: l.svcCtx.Success(),
	}, nil
}
