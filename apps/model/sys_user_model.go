package model

import (
	"context"
	"fmt"
	"strings"
	"zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindWithTid(ctx context.Context, username string, tid int64) (*SysUser, error)
		FindAll(ctx context.Context, ids []int64, nickname, username string, status int32, tenantId int64, page, limit int) ([]*SysUser, int64, error)
		DeleteAll(ctx context.Context, ids []int64) error
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c, opts...),
	}
}

func (m *customSysUserModel) FindWithTid(ctx context.Context, username string, tid int64) (*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where username = ? and tenant_id = ? limit 1", sysUserRows, m.table)
	var resp SysUser
	if err := m.CachedConn.QueryRowNoCacheCtx(ctx, &resp, query, username, tid); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customSysUserModel) FindAll(ctx context.Context, ids []int64, nickname, username string, status int32,
	tenantId int64, page, limit int) ([]*SysUser, int64, error) {
	query := fmt.Sprintf("select %s from %s", sysUserRows, m.table)
	var args []interface{}
	whereCondition := []string{}

	if len(ids) > 0 {
		placeholders := utils.CreateDBPlaceholders(len(ids))
		whereCondition = append(whereCondition, fmt.Sprintf("id in (%s)", placeholders))
		for _, id := range ids {
			args = append(args, id)
		}
	}

	if nickname != "" {
		whereCondition = append(whereCondition, "nickname like ?")
		args = append(args, "%"+nickname+"%")
	}

	whereCondition = append(whereCondition, "status = ?")
	args = append(args, status)

	if tenantId != 0 {
		whereCondition = append(whereCondition, "tenant_id = ?")
		args = append(args, tenantId)
	}

	if username != "" {
		whereCondition = append(whereCondition, "username like ?")
		args = append(args, "%"+username+"%")
	}

	// 拼接条件
	if len(whereCondition) > 0 {
		query += " where " + strings.Join(whereCondition, " and ")
	}

	// 查询总数
	countQuery := "select count(*) from " + m.table
	if len(whereCondition) > 0 {
		countQuery += " where " + strings.Join(whereCondition, " and ")
	}

	var total int64
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	// 添加分页逻辑
	query += " limit ? offset ?"
	args = append(args, limit, (page-1)*limit)

	logc.Debug(ctx, "query: %s, args: %v", query, args)

	var list []*SysUser
	err = m.CachedConn.QueryRowsNoCacheCtx(ctx, &list, query, args...)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customSysUserModel) DeleteAll(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	query := fmt.Sprintf("delete from %s where id in %s", m.table, utils.CreateDBPlaceholders(len(ids)))
	args := make([]interface{}, len(ids))
	for i, v := range ids {
		args[i] = v
	}

	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, args...)
	return err
}
