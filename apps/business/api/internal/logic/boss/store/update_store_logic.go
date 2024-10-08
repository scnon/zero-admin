package store

import (
	"context"

	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStoreLogic {
	return &UpdateStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStoreLogic) UpdateStore(req *types.StoreInfo) (resp *types.StoreInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
