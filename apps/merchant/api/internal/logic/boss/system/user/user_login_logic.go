package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
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
	// 1. 登录
	res, err := l.svcCtx.User.Login(l.ctx, &auth.LoginReq{
		Username: req.Username,
		Password: req.Password,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	// 2. 构造返回数据
	return &types.LoginResp{
		Base: l.svcCtx.Success(),
		Data: types.LoginData{
			UserId:       res.UserId,
			Username:     req.Username,
			NickName:     res.Nickname,
			Avatar:       res.Avatar,
			AccessToken:  res.Token,
			ExpireTime:   res.Expire,
			RefreshToken: res.RefreshToken,
			Roles:        res.Roles,
			Permissions: []string{
				"*:*:*",
			},
		},
	}, nil
}
