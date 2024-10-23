package svc

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/internal/config"
	"xlife/models"
)

type ServiceContext struct {
	Config config.Config

	*gorm.DB
	Casbin *casbin.SyncedCachedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initData(db)
	return &ServiceContext{
		Config: c,

		DB:     db,
		Casbin: initCasbin(c),
	}
}

func initData(db *gorm.DB) {
	err := db.AutoMigrate(&models.SysUser{}, &models.SysRole{}, &models.SysMenu{}, &models.SysAPI{})
	if err != nil {
		panic(err)
	}
	var tenant uint = 1
	// 初始化系统用户
	var count int64
	if res := db.Model(&models.SysUser{}).Count(&count); res.Error != nil {
		panic(res.Error)
	}
	if count == 0 {
		user := models.SysUser{}
		stmt := &gorm.Statement{DB: db}
		err := stmt.Parse(&user)
		if err != nil {
			panic(err)
		}
		query := fmt.Sprintf(`INSERT INTO %s (username, nickname, tenant_id) VALUES ('system', '系统', %d)`,
			stmt.Schema.Table, tenant)
		if res := db.Exec(query); res.Error != nil && !errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			panic(res.Error)
		}
	}
	// 初始化系统角色
	if res := db.Model(&models.SysRole{}).Count(&count); res.Error != nil {
		panic(res.Error)
	}
	if count == 0 {
		res := db.Create(&models.SysRole{Name: "system", Status: 1, ResModel: models.ResModel{CreatorID: 1, TenantID: tenant}})
		if res.Error != nil && !errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			panic(res.Error)
		}
	}
	// 初始化系统菜单
	if res := db.Model(&models.SysMenu{}).Count(&count); res.Error != nil {
		panic(res.Error)
	}
	if count == 0 {
		res := db.Create(&[]models.SysMenu{
			{
				Name:   "system",
				Title:  "系统管理",
				Path:   "/system",
				Status: 1,
				ResModel: models.ResModel{
					Sort:      99,
					CreatorID: 1,
					TenantID:  tenant,
				},
			},
			{
				Name:   "system_user",
				Title:  "用户管理",
				Path:   "/system/user/index",
				Status: 1,
				ResModel: models.ResModel{
					Sort:      1,
					CreatorID: 1,
					TenantID:  tenant,
				},
			},
			{
				Name:   "system_role",
				Title:  "角色管理",
				Path:   "/system/role/index",
				Status: 1,
				ResModel: models.ResModel{
					Sort:      2,
					CreatorID: 1,
					TenantID:  tenant,
				},
			},
			{
				Name:   "system_menu",
				Title:  "菜单管理",
				Path:   "/system/menu/index",
				Status: 1,
				ResModel: models.ResModel{
					Sort:      3,
					CreatorID: 1,
					TenantID:  tenant,
				},
			},
		})
		if res.Error != nil && !errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			panic(res.Error)
		}
	}
}

func initCasbin(c config.Config) *casbin.SyncedCachedEnforcer {
	m, err := model.NewModelFromString(c.Casbin.Model)
	if err != nil {
		panic(err)
	}

	adapter, err := entadapter.NewAdapter("mysql", c.Mysql.Dsn)
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewSyncedCachedEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// 初始化系统角色
	if roles, err := enforcer.GetUsersForRole("system"); len(roles) == 0 && err == nil {
		ok, err := enforcer.AddRoleForUser("system", "r_system", "1")
		if err != nil || !ok {
			panic(err)
		}
		var list [][]string
		for i := 1; i <= 3; i++ {
			list = append(list, []string{"r_system", "1", fmt.Sprint(i), "read"})
		}
		ok, err = enforcer.AddPolicies(list)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return enforcer
}
