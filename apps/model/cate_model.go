package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CateModel = (*customCateModel)(nil)

type (
	// CateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCateModel.
	CateModel interface {
		cateModel
		FindAll(context.Context) ([]*Cate, error)
		DeleteAll(ctx context.Context, ids []int64) error
	}

	customCateModel struct {
		*defaultCateModel
	}
)

// NewCateModel returns a model for the database table.
func NewCateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CateModel {
	return &customCateModel{
		defaultCateModel: newCateModel(conn, c, opts...),
	}
}

func (m *customCateModel) FindAll(ctx context.Context) ([]*Cate, error) {
	query := fmt.Sprintf("select %s from %s", cateRows, m.table)
	var cateList []*Cate
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &cateList, query)
	if err != nil {
		return nil, err
	}
	return cateList, nil
}

func (m *customCateModel) DeleteAll(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, ids)
	return err
}
