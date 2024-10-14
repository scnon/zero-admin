package userlogic

import (
	"context"
	"zero-admin/ent/sysuser"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/ctxdata"
	"zero-admin/pkg/encrypt"
	"zero-admin/pkg/xerr"

	perr "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotFound = xerr.New(xerr.SERVER_COMMON_ERROR, "用户不存在")
	ErrUserPwdError = xerr.New(xerr.SERVER_COMMON_ERROR, "密码不正确")
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
	// 查询用户
	entity, err := l.svcCtx.Ent.SysUser.Query().Where(sysuser.UsernameEQ(in.Username)).Only(l.ctx)
	if err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}
	if entity == nil {
		return nil, perr.WithStack(ErrUserNotFound)
	}
	// 密码验证
	if !encrypt.ValidatePasswordHash(in.Password, entity.Password) {
		return nil, perr.WithStack(ErrUserPwdError)
	}

	// 生成token
	return ctxdata.GetFullJwt(l.svcCtx.Config.JwtAuth.Secret,
		l.svcCtx.Config.JwtAuth.Expire, l.svcCtx.Config.JwtAuth.RefreshExpire, *entity)
}
