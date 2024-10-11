package model

import (
	"context"
	"database/sql"
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
		InsertWithRoles(ctx context.Context, data *SysUser, roleIds []int64) (int64, error)
		FindWithTid(ctx context.Context, username string, tid int64) (*SysUser, error)
		FindAll(ctx context.Context, ids []int64, nickname, username string, status int64, tenantId int64, page, limit int64) ([]*SysUserData, int64, error)
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

func (m *customSysUserModel) InsertWithRoles(ctx context.Context, data *SysUser, roles []int64) (int64, error) {
	var id int64 = 0
	err := m.CachedConn.TransactCtx(ctx, func(ctx context.Context, conn sqlx.Session) error {
		// 占位符数量减去 3 【id, create_time, update_time】
		rows := utils.CreateDBPlaceholders(len(sysUserFieldNames) - 3)
		query := fmt.Sprintf("insert into %s (%s) values (%s)", m.table, sysUserRowsExpectAutoSet, rows)
		res, err := conn.ExecCtx(ctx, query, data.Username, data.Password, data.Nickname, data.Avatar, data.Status,
			data.Sort, data.Remark, data.DepartmentId, data.TenantId, data.Creator, data.Updater, data.IsDeleted)
		if err != nil {
			return err
		}
		id, err = res.LastInsertId()
		if err != nil {
			return err
		}
		var args []interface{}
		holderList := make([]string, 0)
		for _, roleId := range roles {
			args = append(args, id, roleId, data.Creator)
			holderList = append(holderList, "(?, ?, ?)")
		}
		holders := strings.Join(holderList, ",")
		query = fmt.Sprintf("insert into sys_user_role (user_id, role_id, creator) values %s", holders)
		_, err = conn.ExecCtx(ctx, query, args...)
		return err
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *customSysUserModel) FindWithTid(ctx context.Context, username string, tid int64) (*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where username = ? and tenant_id = ? limit 1", sysUserRows, m.table)
	var resp SysUser
	if err := m.CachedConn.QueryRowNoCacheCtx(ctx, &resp, query, username, tid); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SysUserData struct {
	SysUser
	CreatorName sql.NullString `db:"creator_name"`
	UpdaterName sql.NullString `db:"updater_name"`
}

func (m *customSysUserModel) FindAll(ctx context.Context, ids []int64, nickname, username string, status int64,
	tenantId int64, page, pageSize int64) ([]*SysUserData, int64, error) {

	rows := utils.CreateJoinTableRows("sysUser", sysUserFieldNames)
	query := fmt.Sprintf(`select %s, creatorUser.username as creator_name, updaterUser.username as updater_name
        from %s as sysUser
        left join %s as creatorUser on sysUser.creator = creatorUser.id
        left join %s as updaterUser on sysUser.updater = updaterUser.id`,
		rows, m.table, m.table, m.table)
	var args []interface{}
	var whereCondition []string

	if len(ids) > 0 {
		placeholders := utils.CreateDBPlaceholders(len(ids))
		whereCondition = append(whereCondition, fmt.Sprintf("sysUser.id in (%s)", placeholders))
		for _, id := range ids {
			args = append(args, id)
		}
	}

	if nickname != "" {
		whereCondition = append(whereCondition, "sysUser.nickname like ?")
		args = append(args, "%"+nickname+"%")
	}

	whereCondition = append(whereCondition, "sysUser.status = ?")
	args = append(args, status)

	if tenantId != 0 {
		whereCondition = append(whereCondition, "sysUser.tenant_id = ?")
		args = append(args, tenantId)
	}

	if username != "" {
		whereCondition = append(whereCondition, "sysUser.username like ?")
		args = append(args, "%"+username+"%")
	}

	// 拼接条件
	if len(whereCondition) > 0 {
		query += " where " + strings.Join(whereCondition, " and ")
	}

	// 查询总数
	countQuery := "select count(*) from " + m.table
	if len(whereCondition) > 0 {
		condition := strings.Join(whereCondition, " and ")
		condition = strings.Replace(condition, "sysUser.", "", -1)
		countQuery += " where " + condition
	}

	var total int64
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	// 添加分页逻辑
	query += " limit ? offset ?"
	args = append(args, pageSize, (page-1)*pageSize)

	logc.Debug(ctx, "query: %s, args: %v", query, args)

	var list []*SysUserData
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

	query := fmt.Sprintf("delete from %s where id in (%s)", m.table, utils.CreateDBPlaceholders(len(ids)))
	args := make([]interface{}, len(ids))
	for i, v := range ids {
		args[i] = v
	}

	_, err := m.CachedConn.ExecNoCacheCtx(ctx, query, args...)
	return err
}
