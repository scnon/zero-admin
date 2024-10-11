package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/pkg/utils"
)

var _ SysRoleMenuModel = (*customSysRoleMenuModel)(nil)

type (
	// SysRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleMenuModel.
	SysRoleMenuModel interface {
		sysRoleMenuModel
		FindAllByRoleIds(ctx context.Context, roleIds []int64) ([]*SysRoleMenu, error)
	}

	customSysRoleMenuModel struct {
		*defaultSysRoleMenuModel
	}
)

// NewSysRoleMenuModel returns a model for the database table.
func NewSysRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRoleMenuModel {
	return &customSysRoleMenuModel{
		defaultSysRoleMenuModel: newSysRoleMenuModel(conn, c, opts...),
	}
}

func (m *customSysRoleMenuModel) FindAllByRoleIds(ctx context.Context, roleIds []int64) ([]*SysRoleMenu, error) {
	rows := utils.CreateDBPlaceholders(len(roleIds))
	query := fmt.Sprintf(`select %s from %s where role_id in (%s)`, sysRoleMenuRows, m.table, rows)
	var list []*SysRoleMenu
	var args []interface{}
	for _, v := range roleIds {
		args = append(args, v)
	}
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &list, query, args...)
	return list, err
}
