package ctxdata

import (
	"time"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"
	"xlife/pkg/xerr"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const Identify = "zero-auth"

func GetJwtToken(secretKey string, iat, seconds int64, uid uint) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func GetFullJwt(secretKey string, expire, refreshExpire int64, entity models.SysUser) (*auth.LoginResp, error) {
	// 生成token
	now := time.Now().Unix()
	token, err := GetJwtToken(secretKey, now, expire,
		entity.ID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "ctxdata get jwt token err %v", err)
	}

	refresh, err := GetJwtToken(secretKey, now, refreshExpire, entity.ID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "ctxdata get jwt token err %v", err)
	}

	return &auth.LoginResp{
		UserId:       uint64(entity.ID),
		Nickname:     entity.Nickname,
		Avatar:       entity.Avatar,
		Token:        token,
		Expire:       uint64(now + expire),
		RefreshToken: refresh,
	}, nil
}
