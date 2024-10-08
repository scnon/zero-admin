package logic

import (
	"context"
	"errors"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"
	"zero-admin/pkg/ctxdata"
	"zero-admin/pkg/encrypt"
	"zero-admin/pkg/xerr"

	perr "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotRegister = xerr.New(xerr.SERVER_COMMON_ERROR, "用户不存在")
	ErrUserPwdError    = xerr.New(xerr.SERVER_COMMON_ERROR, "密码不正确")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *admin.LoginReq) (*admin.LoginResp, error) {
	userEntity, err := l.svcCtx.UserModel.FindWithTid(l.ctx, in.Username, in.TenantId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, perr.WithStack(ErrUserNotRegister)
		}

		return nil, perr.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}

	// 密码验证
	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password) {
		return nil, perr.WithStack(ErrUserPwdError)
	}

	// 生成token
	return ctxdata.GetFullJwt(l.svcCtx.Config.JwtAuth.Secret,
		l.svcCtx.Config.JwtAuth.Expire, l.svcCtx.Config.JwtAuth.RefreshExpire, *userEntity)
}
