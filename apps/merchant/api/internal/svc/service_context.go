package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"xlife/apps/auth/rpc/client/dept"
	"xlife/apps/auth/rpc/client/menu"
	"xlife/apps/auth/rpc/client/role"
	"xlife/apps/auth/rpc/client/user"
	"xlife/apps/merchant/api/internal/config"
	"xlife/apps/merchant/api/internal/types"
)

type ServiceContext struct {
	Config config.Config

	user.User
	role.Role
	menu.Menu
	dept.Dept
	//merchant_client.Business
	//store_client.Store
	// product_client.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User: user.NewUser(zrpc.MustNewClient(c.AuthRpc)),
		Role: role.NewRole(zrpc.MustNewClient(c.AuthRpc)),
		Menu: menu.NewMenu(zrpc.MustNewClient(c.AuthRpc)),
		Dept: dept.NewDept(zrpc.MustNewClient(c.AuthRpc)),
		//Business: merchant_client.NewBusiness(zrpc.MustNewClient(c.BusinessRpc)),
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
