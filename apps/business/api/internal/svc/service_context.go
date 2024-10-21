package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"xlife/apps/auth/rpc/client/dept"
	"xlife/apps/auth/rpc/client/menu"
	"xlife/apps/auth/rpc/client/role"
	"xlife/apps/auth/rpc/client/user"
	"xlife/apps/business/api/internal/config"
	"xlife/apps/business/api/internal/types"
)

type ServiceContext struct {
	Config config.Config

	user.User
	role.Role
	menu.Menu
	dept.Dept
	//business_client.Business
	//store_client.Store
	// product_client.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User: user.NewUser(zrpc.MustNewClient(c.AdminRpc)),
		Role: role.NewRole(zrpc.MustNewClient(c.AdminRpc)),
		Menu: menu.NewMenu(zrpc.MustNewClient(c.AdminRpc)),
		Dept: dept.NewDept(zrpc.MustNewClient(c.AdminRpc)),
		//Business: business_client.NewBusiness(zrpc.MustNewClient(c.BusinessRpc)),
		//Store:    store_client.NewStore(zrpc.MustNewClient(c.StoreRpc)),
		// Product:  product_client.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

func (svcCtx *ServiceContext) Success() types.Base {
	return types.Base{
		Success: true,
		Msg:     "ok",
	}
}
