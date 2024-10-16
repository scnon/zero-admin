package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	result, err := l.svcCtx.User.Login(l.ctx, &auth.LoginReq{
		Username: req.Username,
		Password: req.Password,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Base: l.svcCtx.Success(),
		Data: types.LoginData{
			UserId:       result.UserId,
			NickName:     result.Nickname,
			Avatar:       result.Avatar,
			AccessToken:  result.Token,
			ExpireTime:   result.Expire,
			RefreshToken: result.RefreshToken,
		},
	}, nil
}
