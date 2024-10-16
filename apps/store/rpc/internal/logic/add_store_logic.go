package logic

import (
	"context"

	"xlife/apps/model"
	"xlife/apps/store/rpc/internal/svc"
	"xlife/apps/store/rpc/store"
	"xlife/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStoreLogic {
	return &AddStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStoreLogic) AddStore(in *store.AddStoreReq) (*store.StoreInfo, error) {
	var storeEntity model.Store
	copier.Copy(&storeEntity, in)

	_, err := l.svcCtx.StoreModel.Insert(l.ctx, &storeEntity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert store err %v", err)
	}

	var info store.StoreInfo
	err = copier.Copy(&info, &storeEntity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "copy store err %v", err)
	}
	return &info, nil
}
