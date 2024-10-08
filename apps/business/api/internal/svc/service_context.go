package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-admin/apps/admin/rpc/admin_client"
	"zero-admin/apps/business/api/internal/config"
	"zero-admin/apps/business/api/internal/types"
	"zero-admin/apps/business/rpc/business_client"
)

type ServiceContext struct {
	Config config.Config

	admin_client.Admin
	business_client.Business
	//store_client.Store
	// product_client.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Admin:    admin_client.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
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
