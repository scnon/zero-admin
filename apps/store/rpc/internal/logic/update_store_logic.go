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

var (
	ErrStoreNotFound = xerr.NewMsg("店铺不存在")
)

type UpdateStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStoreLogic {
	return &UpdateStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStoreLogic) UpdateStore(in *store.StoreInfo) (*store.StoreInfo, error) {
	storeEntity, err := l.svcCtx.StoreModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrStoreNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find store err %v", err)
	}

	storeEntity.Name = in.Name
	storeEntity.Phone = in.Phone
	storeEntity.Address = in.Address
	storeEntity.BusinessId = in.BusinessId
	storeEntity.Status = int64(in.Status)
	storeEntity.StartTime = in.StartTime
	storeEntity.EndTime = in.EndTime

	err = l.svcCtx.StoreModel.Update(l.ctx, storeEntity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update store err %v", err)
	}

	var storeInfo store.StoreInfo
	err = copier.Copy(&storeInfo, storeEntity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "copy store err %v", err)
	}

	return &storeInfo, nil
}
