package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"
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
	result, err := l.svcCtx.User.Login(l.ctx, &admin.LoginReq{
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
			AccessToken:  result.Token,
			ExpireTime:   result.Expire,
			RefreshToken: result.RefreshToken,
		},
	}, nil
}
