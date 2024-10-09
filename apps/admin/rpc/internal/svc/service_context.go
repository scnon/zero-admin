package svc

import (
	"zero-admin/apps/admin/rpc/internal/config"
	"zero-admin/apps/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.SysUserModel
	RoleModel model.SysRoleModel
	MenuModel model.SysMenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Dsn)

	return &ServiceContext{
		Config: c,

		UserModel: model.NewSysUserModel(sqlConn, c.Cache),
		RoleModel: model.NewSysRoleModel(sqlConn, c.Cache),
		MenuModel: model.NewSysMenuModel(sqlConn, c.Cache),
	}
}
