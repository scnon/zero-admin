package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Dsn string
	}

	JwtAuth struct {
		Secret        string
		Expire        int64
		RefreshExpire int64
	}

	Cache cache.CacheConf
}
