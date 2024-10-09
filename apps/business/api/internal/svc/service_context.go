package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-admin/apps/admin/rpc/client/menu"
	"zero-admin/apps/admin/rpc/client/user"
	"zero-admin/apps/business/api/internal/config"
	"zero-admin/apps/business/api/internal/types"
	"zero-admin/apps/business/rpc/business_client"
)

type ServiceContext struct {
	Config config.Config

	user.User
	menu.Menu
	business_client.Business
	//store_client.Store
	// product_client.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User:     user.NewUser(zrpc.MustNewClient(c.AdminRpc)),
		Menu:     menu.NewMenu(zrpc.MustNewClient(c.AdminRpc)),
		Business: business_client.NewBusiness(zrpc.MustNewClient(c.BusinessRpc)),
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
