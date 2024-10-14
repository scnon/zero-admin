package ctxdata

import (
	"time"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/ent"
	"zero-admin/pkg/xerr"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const Identify = "zero-admin"

func GetJwtToken(secretKey string, iat, seconds int64, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func GetFullJwt(secretKey string, expire, refreshExpire int64, entity ent.SysUser) (*admin.LoginResp, error) {
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

	return &admin.LoginResp{
		UserId:       entity.ID,
		Nickname:     entity.Nickname,
		Token:        token,
		Expire:       now + expire,
		RefreshToken: refresh,
	}, nil
}
