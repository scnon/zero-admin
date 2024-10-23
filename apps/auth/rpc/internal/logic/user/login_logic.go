package userlogic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
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
	ErrUserNotFound = xerr.NewMsg("用户不存在")
	ErrUserPwdError = xerr.NewMsg("密码不正确")
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
	// 1. 查询用户是否存在
	var entity models.SysUser
	res := l.svcCtx.DB.Where("username = ?", in.Username).Where("tenant_id = ?", in.TenantId).First(&entity)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrUserNotFound)
		}
		return nil, perr.Wrapf(xerr.NewDBErr(), "查询用户失败: %v", res.Error)
	}
	// 2. 密码验证
	if !encrypt.ValidatePasswordHash(in.Password, entity.Password) {
		return nil, perr.WithStack(ErrUserPwdError)
	}
	// 3. 生成token
	tokenRes, err := ctxdata.GetFullJwt(l.svcCtx.Config.JwtAuth.Secret,
		l.svcCtx.Config.JwtAuth.Expire, l.svcCtx.Config.JwtAuth.RefreshExpire, entity)
	if err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "生成token失败: %v", err)
	}
	// 4. 查询用户角色
	tenantId := strconv.FormatUint(in.TenantId, 10)
	roles := l.svcCtx.Casbin.GetRolesForUserInDomain(tokenRes.Username, tenantId)
	tokenRes.Roles = roles
	return tokenRes, nil
}
