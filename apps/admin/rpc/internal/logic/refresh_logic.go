package logic

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/ctxdata"
	"zero-admin/pkg/xerr"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
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

func (l *RefreshLogic) Refresh(in *admin.RefreshReq) (*admin.LoginResp, error) {
	token, err := jwt.Parse(in.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(l.svcCtx.Config.JwtAuth.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims[ctxdata.Identify].(int64)
		userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "find user err %v", err)
		}
		return ctxdata.GetFullJwt(l.svcCtx.Config.JwtAuth.Secret, l.svcCtx.Config.JwtAuth.Expire,
			l.svcCtx.Config.JwtAuth.RefreshExpire, *userEntity)
	}

	return nil, errors.Wrapf(xerr.NewInternalErr(), "refresh token is invalid")
}
