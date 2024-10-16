package logic

import (
	"context"
	"log"

	"xlife/apps/model"
	"xlife/apps/store/rpc/internal/svc"
	"xlife/apps/store/rpc/store"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreLogic {
	return &GetStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoreLogic) GetStore(in *store.GetStoreReq) (*store.GetStoreResp, error) {
	storeList := make([]*store.StoreInfo, 0)
	page := 1
	if in.Page > 0 {
		page = int(in.Page)
	}

	pageSize := 10
	if in.PageSize > 0 {
		pageSize = int(in.PageSize)
	}

	log.Println("page:", page, "pageSize:", pageSize)
	result, total, err := l.svcCtx.StoreModel.FindAll(l.ctx, in.Ids, in.BusinessIds, page, pageSize)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrStoreNotFound)
		}

		return nil, errors.Wrapf(xerr.NewDBErr(), "find store err %v", err)
	}

	for _, storeEntity := range result {
		storeList = append(storeList, &store.StoreInfo{
			Id:         storeEntity.Id,
			Name:       storeEntity.Name,
			Phone:      storeEntity.Phone,
			Address:    storeEntity.Address,
			BusinessId: storeEntity.BusinessId,
			Status:     int32(storeEntity.Status),
			StartTime:  storeEntity.StartTime,
			EndTime:    storeEntity.EndTime,
		})
	}

	return &store.GetStoreResp{
		Store: storeList,
		Total: total,
	}, nil
}
