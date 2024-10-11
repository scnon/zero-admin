package model

import (
	"context"
	"database/sql"
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
		FindAll(ctx context.Context, tenantId int64, page, pageSize int64) ([]*SysMenuData, int64, error)
		FindAllByIds(ctx context.Context, ids []int64) ([]*SysMenuData, error)
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

type SysMenuData struct {
	SysMenu
	CreatorName sql.NullString `db:"creator_name"`
	UpdaterName sql.NullString `db:"updater_name"`
}

func (m *customSysMenuModel) FindAll(ctx context.Context, tenantId int64, page, pageSize int64) ([]*SysMenuData, int64, error) {
	rows := utils.CreateJoinTableRows("sysMenu", sysMenuFieldNames)
	query := fmt.Sprintf(`select %s,creatorUser.username as creator_name, updaterUser.username as updater_name
		from %s as sysMenu
		left join %s as creatorUser on sysMenu.creator = creatorUser.id
        left join %s as updaterUser on sysMenu.updater = updaterUser.id`,
		rows, m.table, "sys_user", "sys_user")
	offset := (page - 1) * pageSize
	args := []interface{}{pageSize, offset}

	if tenantId != 0 {
		query = fmt.Sprintf("%s where sysMenu.tenant_id = ? limit ? offset ?", query)
		args = []interface{}{tenantId, pageSize, offset}
	} else {
		query = fmt.Sprintf("%s limit ? offset ?", query)
		args = []interface{}{pageSize, offset}
	}

	var list []*SysMenuData
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &list, query, args...)
	if err != nil {
		return nil, 0, err
	}
	// 获取总数
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

func (m *customSysMenuModel) FindAllByIds(ctx context.Context, ids []int64) ([]*SysMenuData, error) {
	rows := utils.CreateJoinTableRows("sysMenu", sysMenuFieldNames)
	query := fmt.Sprintf(`select %s,creatorUser.username as creator_name, updaterUser.username as updater_name
		from %s as sysMenu
		left join %s as creatorUser on sysMenu.creator = creatorUser.id
		left join %s as updaterUser on sysMenu.updater = updaterUser.id
		where sysMenu.id in (%s)`,
		rows, m.table, "sys_user", "sys_user", utils.CreateDBPlaceholders(len(ids)))
	var list []*SysMenuData
	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &list, query, args...)
	if err != nil {
		return nil, err
	}
	return list, nil
}
