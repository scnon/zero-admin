package model

import (
	"context"
	"fmt"
	"strings"
	"zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BusinessModel = (*customBusinessModel)(nil)

type (
	// BusinessModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBusinessModel.
	BusinessModel interface {
		businessModel
		FindAll(ctx context.Context, ids []int64, adminIds []int64) ([]*Business, error)
		DeleteAll(ctx context.Context, ids []int64) error
	}

	customBusinessModel struct {
		*defaultBusinessModel
	}
)

// NewBusinessModel returns a model for the database table.
func NewBusinessModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BusinessModel {
	return &customBusinessModel{
		defaultBusinessModel: newBusinessModel(conn, c, opts...),
	}
}

func (m *customBusinessModel) DeleteAll(ctx context.Context, ids []int64) error {
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

func (m *customBusinessModel) FindAll(ctx context.Context, ids []int64, adminIds []int64) ([]*Business, error) {
	var resp []*Business
	args := []interface{}{}

	query := fmt.Sprintf("select %s from %s ", businessRows, m.table)
	whereCondition := []string{}
	if len(ids) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("id in (%s)", utils.CreateDBPlaceholders(len(ids))))
		// args = append(args, ids)
		for _, id := range ids {
			args = append(args, id)
		}
	}
	if len(adminIds) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("admin_id in (%s)", utils.CreateDBPlaceholders(len(adminIds))))
		// args = append(args, adminIds)
		for _, adminId := range adminIds {
			args = append(args, adminId)
		}
	}
	if len(whereCondition) > 0 {
		query += " where " + strings.Join(whereCondition, " and ")
	}
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	return resp, err
}
