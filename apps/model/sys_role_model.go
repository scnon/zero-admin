package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/pkg/utils"
)

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		DeleteAll(ctx context.Context, roleIds []int64) error
		FindAll(ctx context.Context, tenantId, page, pageSize int64) ([]*SysRoleData, int64, error)
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

type SysRoleData struct {
	SysRole
	CreatorName sql.NullString `db:"creator_name"`
	UpdaterName sql.NullString `db:"updater_name"`
}

func (m *customSysRoleModel) FindAll(ctx context.Context, tenantId, page, pageSize int64) (roles []*SysRoleData, total int64, err error) {
	offset := (page - 1) * pageSize
	rows := utils.CreateJoinTableRows("sysRole", sysRoleFieldNames)
	query := fmt.Sprintf(`select %s ,creatorUser.username as creator_name, updaterUser.username as updater_name
		from %s as sysRole
		left join %s as creatorUser on sysRole.creator = creatorUser.id
        left join %s as updaterUser on sysRole.updater = updaterUser.id`,
		rows, m.table, "sys_user", "sys_user")

	if tenantId != 0 {
		query = fmt.Sprintf("%s where sysRole.tenant_id = %d", query, tenantId)
	}
	query = fmt.Sprintf("%s limit %d offset %d", query, pageSize, offset)
	err = m.CachedConn.QueryRowsNoCacheCtx(ctx, &roles, query)
	if err != nil {
		return
	}

	query = fmt.Sprintf("select count(*) from %s", m.table)
	if tenantId != 0 {
		query = fmt.Sprintf("%s where tenant_id = %d", query, tenantId)
	}
	err = m.CachedConn.QueryRowNoCacheCtx(ctx, &total, query)
	return
}
