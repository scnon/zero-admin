package svc

import (
	"zero-admin/apps/model"
	"zero-admin/apps/store/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	model.StoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Dsn)

	return &ServiceContext{
		Config: c,

		StoreModel: model.NewStoreModel(sqlConn, c.Cache),
	}
}
