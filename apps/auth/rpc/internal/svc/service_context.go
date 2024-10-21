package svc

import (
	"errors"
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

	res := db.Where("id = ?", 1).First(&models.SysUser{})
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Exec(`INSERT INTO sys_users (id, username, nickname, created_at) VALUES (1, 'system', '系统', NOW())`)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
				// 重复键错误，忽略
			} else {
				panic(res.Error)
			}
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

	return enforcer
}
