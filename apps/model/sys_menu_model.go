package model

import (
	"context"
	"fmt"
	"zero-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		DeleteAll(ctx context.Context, ids []int64) error
		FindAll(ctx context.Context, tenantId int64, page, pageSize int64) (*[]SysMenu, int64, error)
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn, c, opts...),
	}
}

func (m *customSysMenuModel) DeleteAll(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf(`delete from sys_menu where id in (%s)`, utils.CreateDBPlaceholders(len(ids)))
	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, ids)
	if err != nil {
		return err
	}
	return nil
}

func (m *customSysMenuModel) FindAll(ctx context.Context, tenantId int64, page, pageSize int64) (*[]SysMenu, int64, error) {
	query := `select * from sys_menu where tenant_id = ? limit ? offset ?`
	offset := (page - 1) * pageSize
	args := []interface{}{tenantId, pageSize, offset}
	if tenantId == 0 {
		query = `select * from sys_menu limit ? offset ?`
		args = []interface{}{pageSize, offset}
	}

	list := &[]SysMenu{}
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, list, query, args...)
	if err != nil {
		return nil, 0, err
	}
	var total int64
	query = `select count(*) from sys_menu where tenant_id = ?`
	if tenantId == 0 {
		query = `select count(*) from sys_menu`
	}
	err = m.CachedConn.QueryRowNoCacheCtx(ctx, &total, query, tenantId)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
