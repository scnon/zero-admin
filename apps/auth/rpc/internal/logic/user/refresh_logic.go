package userlogic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"

	"xlife/apps/auth/rpc/internal/svc"
	"xlife/pkg/ctxdata"
	"xlife/pkg/xerr"

	"github.com/golang-jwt/jwt/v4"
	perr "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *auth.RefreshReq) (*auth.LoginResp, error) {
	// 1. 解析 refresh token
	token, err := jwt.Parse(in.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(l.svcCtx.Config.JwtAuth.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	// 2. 判断 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims[ctxdata.Identify].(int64)
		var entity models.SysUser
		res := l.svcCtx.DB.Where("id = ?", userId).First(&entity)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return nil, perr.WithStack(ErrUserNotFound)
			}
			return nil, perr.Wrapf(xerr.NewDBErr(), "查询用户失败: %v", res.Error)
		}
		// 3. 生成新的 token
		return ctxdata.GetFullJwt(l.svcCtx.Config.JwtAuth.Secret, l.svcCtx.Config.JwtAuth.Expire,
			l.svcCtx.Config.JwtAuth.RefreshExpire, entity)
	}

	return nil, perr.Wrapf(xerr.NewInternalErr(), "refresh token is invalid")
}
