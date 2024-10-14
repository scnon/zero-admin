package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
	"zero-admin/apps/admin/rpc/internal/config"
	"zero-admin/ent"
)

type ServiceContext struct {
	Config config.Config

	Ent *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := ent.Open("mysql", c.Mysql.Dsn)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,

		Ent: client,
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
