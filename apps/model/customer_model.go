package model

import (
	"context"
	"fmt"
	"zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CustomerModel = (*customCustomerModel)(nil)

type (
	// CustomerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerModel.
	CustomerModel interface {
		customerModel
		FindAll(ctx context.Context, ids []int64) ([]*Customer, error)
		DeleteAll(ctx context.Context, ids []int64) error
	}

	customCustomerModel struct {
		*defaultCustomerModel
	}
)

// NewCustomerModel returns a model for the database table.
func NewCustomerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CustomerModel {
	return &customCustomerModel{
		defaultCustomerModel: newCustomerModel(conn, c, opts...),
	}
}

func (m *customCustomerModel) DeleteAll(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	var query string
	var args []interface{}
	if len(ids) > 0 {
		query = fmt.Sprintf("delete from %s where `id` in (%s)", m.table,
			utils.CreateDBPlaceholders(len(ids)),
		)
		for _, id := range ids {
			args = append(args, id)
		}
	}

	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, args...)
	return err
}

func (m *customCustomerModel) FindAll(ctx context.Context, ids []int64) ([]*Customer, error) {
	var resp []*Customer
	args := []interface{}{}

	query := fmt.Sprintf("select %s from %s ", businessRows, m.table)
	if len(ids) > 0 {
		query += fmt.Sprintf("where `id` in (%s)", utils.CreateDBPlaceholders(len(ids)))
		for _, id := range ids {
			args = append(args, id)
		}
	}
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, args)
	return resp, err
}
