package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SkuModel = (*customSkuModel)(nil)

type (
	// SkuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSkuModel.
	SkuModel interface {
		skuModel
	}

	customSkuModel struct {
		*defaultSkuModel
	}
)

// NewSkuModel returns a model for the database table.
func NewSkuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SkuModel {
	return &customSkuModel{
		defaultSkuModel: newSkuModel(conn, c, opts...),
	}
}
