package user

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRestPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRestPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRestPwdLogic {
	return &UserRestPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRestPwdLogic) UserRestPwd(req *types.UserRestPwdReq) (resp *types.UserRestPwdResp, err error) {
	uid := ctxdata.GetUId(l.ctx)
	_, err = l.svcCtx.User.ResetPassword(l.ctx, &auth.ResetPasswordReq{
		UserId:   req.UserId,
		Password: req.Password,
		Op:       uid,
	})
	if err != nil {
		return
	}

	return &types.UserRestPwdResp{
		Base: l.svcCtx.Success(),
	}, nil
}
