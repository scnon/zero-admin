package svc

import (
	"xlife/apps/business/rpc/internal/config"
	"xlife/apps/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	model.BusinessModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Dsn)

	return &ServiceContext{
		Config: c,

		BusinessModel: model.NewBusinessModel(sqlConn, c.Cache),
	}
}
