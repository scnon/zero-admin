package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		DeleteAll(ctx context.Context, roleIds []int64) error
		FindAll(ctx context.Context, page, pageSize int64) ([]SysRole, int64, error)
	}

	customSysRoleModel struct {
		*defaultSysRoleModel
	}
)

// NewSysRoleModel returns a model for the database table.
func NewSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRoleModel {
	return &customSysRoleModel{
		defaultSysRoleModel: newSysRoleModel(conn, c, opts...),
	}
}

func (m *customSysRoleModel) DeleteAll(ctx context.Context, roleIds []int64) error {
	query := fmt.Sprintf("delete from %s where id in ?", m.table)
	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, roleIds)
	return err
}

func (m *customSysRoleModel) FindAll(ctx context.Context, page, pageSize int64) (roles []SysRole, total int64, err error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s limit %d offset %d", sysRoleRows, m.table, pageSize, offset)
	err = m.CachedConn.QueryRowsNoCacheCtx(ctx, &roles, query)
	if err != nil {
		return
	}

	query = fmt.Sprintf("select count(*) from %s", m.table)
	err = m.CachedConn.QueryRowNoCacheCtx(ctx, &total, query)
	return
}
