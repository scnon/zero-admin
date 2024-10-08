package model

import (
	"context"
	"fmt"
	"strings"
	"zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		FindAll(ctx context.Context, ids []int64, businessIds []int64, storeIds []int64, cateIds []int64) ([]*Product, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

func (m *customProductModel) FindAll(ctx context.Context, ids []int64, businessIds []int64, storeIds []int64, cateIds []int64) ([]*Product, error) {
	var resp []*Product
	query := fmt.Sprintf("select %s from %s ", productRows, m.table)
	args := []interface{}{}
	whereCondition := []string{}

	if len(ids) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("id in (%s)", utils.CreateDBPlaceholders(len(ids))))
		for _, id := range ids {
			args = append(args, id)
		}
	}
	if len(businessIds) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("business_id in (%s)", utils.CreateDBPlaceholders(len(businessIds))))
		for _, businessId := range businessIds {
			args = append(args, businessId)
		}
	}
	if len(storeIds) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("store_id in (%s)", utils.CreateDBPlaceholders(len(storeIds))))
		for _, storeId := range storeIds {
			args = append(args, storeId)
		}
	}
	if len(cateIds) > 0 {
		whereCondition = append(whereCondition, fmt.Sprintf("cate_id in (%s)", utils.CreateDBPlaceholders(len(cateIds))))
		for _, cateId := range cateIds {
			args = append(args, cateId)
		}
	}
	if len(whereCondition) > 0 {
		query += " where " + strings.Join(whereCondition, " and ")
	}

	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
