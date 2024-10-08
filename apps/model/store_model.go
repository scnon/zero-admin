package model

import (
	"context"
	"fmt"
	"strings"
	"zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StoreModel = (*customStoreModel)(nil)

type (
	// StoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStoreModel.
	StoreModel interface {
		storeModel
		FindAll(ctx context.Context, ids []int64, businessIds []int64, page int, pageSize int) ([]*Store, int64, error)
		DeleteAll(ctx context.Context, ids []int64, businessIds []int64) error
	}

	customStoreModel struct {
		*defaultStoreModel
	}
)

// NewStoreModel returns a model for the database table.
func NewStoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StoreModel {
	return &customStoreModel{
		defaultStoreModel: newStoreModel(conn, c, opts...),
	}
}

func (m *customStoreModel) FindAll(ctx context.Context, ids []int64, businessIds []int64,
	page int, pageSize int) ([]*Store, int64, error) {
	query := fmt.Sprintf("select %s from %s", storeRows, m.table)
	var storeList []*Store
	args := []interface{}{}
	whereCondition := []string{}

	if len(ids) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("`id` in (%s)",
			utils.CreateDBPlaceholders(len(ids))))
		for _, id := range ids {
			args = append(args, id)
		}
	}
	if len(businessIds) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("`business_id` in (%s)",
			utils.CreateDBPlaceholders(len(businessIds))))
		for _, businessId := range businessIds {
			args = append(args, businessId)
		}
	}

	// 计算总记录数的查询
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)
	if len(whereCondition) > 0 {
		countQuery += " WHERE " + strings.Join(whereCondition, " AND ")
	}
	// 执行总数查询
	var total int64
	if err := m.CachedConn.QueryRowNoCacheCtx(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	if len(whereCondition) > 0 {
		query += " WHERE " + strings.Join(whereCondition, " AND ")
	}
	offset := (page - 1) * pageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	if err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &storeList, query, args...); err != nil {
		return nil, 0, err
	}
	return storeList, total, nil
}

func (m *customStoreModel) DeleteAll(ctx context.Context, ids []int64, businessIds []int64) error {
	if len(ids) == 0 && len(businessIds) == 0 {
		return nil
	}

	query := fmt.Sprintf("delete from %s", m.table)
	args := []interface{}{}
	whereCondition := []string{}

	if len(ids) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("where `id` in (%s)",
			utils.CreateDBPlaceholders(len(ids))))
		args = append(args, ids)
	}
	if len(businessIds) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("where `business_id` in (%s)",
			utils.CreateDBPlaceholders(len(businessIds))))
		args = append(args, businessIds)
	}
	if len(whereCondition) > 0 {
		query += " where " + strings.Join(whereCondition, " and ")
	}
	query = fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, args)
	return err
}
