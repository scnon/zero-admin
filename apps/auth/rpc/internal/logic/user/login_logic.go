package userlogic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"

	"xlife/apps/auth/rpc/internal/svc"
	"xlife/pkg/ctxdata"
	"xlife/pkg/encrypt"
	"xlife/pkg/xerr"

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

func (l *LoginLogic) Login(in *auth.LoginReq) (*auth.LoginResp, error) {
	// 查询用户
	var entity models.SysUser
	res := l.svcCtx.DB.Where("username = ?", in.Username).First(&entity)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrUserNotFound)
		}
		return nil, perr.Wrapf(xerr.NewDBErr(), "查询用户失败: %v", res.Error)
	}
	// 密码验证
	if !encrypt.ValidatePasswordHash(in.Password, entity.Password) {
		return nil, perr.WithStack(ErrUserPwdError)
	}

	// 生成token
	return ctxdata.GetFullJwt(l.svcCtx.Config.JwtAuth.Secret,
		l.svcCtx.Config.JwtAuth.Expire, l.svcCtx.Config.JwtAuth.RefreshExpire, entity)
}
