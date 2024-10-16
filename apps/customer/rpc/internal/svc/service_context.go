package svc

import (
	"xlife/apps/customer/rpc/internal/config"
	"xlife/apps/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	model.CustomerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Dsn)
	return &ServiceContext{
		Config: c,

		CustomerModel: model.NewCustomerModel(sqlConn, c.Cache),
	}
}
