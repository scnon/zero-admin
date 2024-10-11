package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserRoleModel = (*customSysUserRoleModel)(nil)

type (
	// SysUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserRoleModel.
	SysUserRoleModel interface {
		sysUserRoleModel
		FindAllByUserId(ctx context.Context, userId int64) ([]*SysUserRole, error)
	}

	customSysUserRoleModel struct {
		*defaultSysUserRoleModel
	}
)

// NewSysUserRoleModel returns a model for the database table.
func NewSysUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysUserRoleModel {
	return &customSysUserRoleModel{
		defaultSysUserRoleModel: newSysUserRoleModel(conn, c, opts...),
	}
}

func (m *customSysUserRoleModel) FindAllByUserId(ctx context.Context, userId int64) ([]*SysUserRole, error) {
	query := fmt.Sprintf("select * from %s where user_id = ?", m.table)
	var resp []*SysUserRole
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
