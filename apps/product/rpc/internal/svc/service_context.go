package svc

import (
	"xlife/apps/model"
	"xlife/apps/product/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	model.SkuModel
	model.CateModel
	model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Dsn)
	return &ServiceContext{
		Config: c,

		SkuModel:     model.NewSkuModel(sqlConn, c.Cache),
		CateModel:    model.NewCateModel(sqlConn, c.Cache),
		ProductModel: model.NewProductModel(sqlConn, c.Cache),
	}
}
