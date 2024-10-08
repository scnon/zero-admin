// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	storeFieldNames          = builder.RawFieldNames(&Store{})
	storeRows                = strings.Join(storeFieldNames, ",")
	storeRowsExpectAutoSet   = strings.Join(stringx.Remove(storeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	storeRowsWithPlaceHolder = strings.Join(stringx.Remove(storeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheStoreIdPrefix = "cache:store:id:"
)

type (
	storeModel interface {
		Insert(ctx context.Context, data *Store) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Store, error)
		Update(ctx context.Context, data *Store) error
		Delete(ctx context.Context, id int64) error
	}

	defaultStoreModel struct {
		sqlc.CachedConn
		table string
	}

	Store struct {
		Id         int64     `db:"id"`          // 店铺ID
		BusinessId int64     `db:"business_id"` // 商家ID
		Name       string    `db:"name"`        // 店铺名称
		Phone      string    `db:"phone"`       // 店铺电话
		Status     int64     `db:"status"`      // 店铺状态 0:禁用 1:启用
		Address    string    `db:"address"`     // 店铺地址
		StartTime  string    `db:"start_time"`  // 营业开始时间
		EndTime    string    `db:"end_time"`    // 营业结束时间
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newStoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultStoreModel {
	return &defaultStoreModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`store`",
	}
}

func (m *defaultStoreModel) Delete(ctx context.Context, id int64) error {
	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, storeIdKey)
	return err
}

func (m *defaultStoreModel) FindOne(ctx context.Context, id int64) (*Store, error) {
	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, id)
	var resp Store
	err := m.QueryRowCtx(ctx, &resp, storeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", storeRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStoreModel) Insert(ctx context.Context, data *Store) (sql.Result, error) {
	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, storeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.BusinessId, data.Name, data.Phone, data.Status, data.Address, data.StartTime, data.EndTime)
	}, storeIdKey)
	return ret, err
}

func (m *defaultStoreModel) Update(ctx context.Context, data *Store) error {
	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, storeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.BusinessId, data.Name, data.Phone, data.Status, data.Address, data.StartTime, data.EndTime, data.Id)
	}, storeIdKey)
	return err
}

func (m *defaultStoreModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheStoreIdPrefix, primary)
}

func (m *defaultStoreModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", storeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultStoreModel) tableName() string {
	return m.table
}
