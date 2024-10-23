package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
	}

	Tenant      uint64
	AuthRpc     zrpc.RpcClientConf
	MerchantRpc zrpc.RpcClientConf
	StoreRpc    zrpc.RpcClientConf
	ProductRpc  zrpc.RpcClientConf
}
