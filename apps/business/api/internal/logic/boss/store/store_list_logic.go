package store

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"
)

type StoreListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreListLogic {
	return &StoreListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreListLogic) StoreList(req *types.StoreListReq) (resp *types.StoreListResp, err error) {
	//uid := ctxdata.GetUId(l.ctx)
	//if uid == 0 {
	//	return nil, err
	//}
	//businesses, err := l.svcCtx.Business.GetBusiness(l.ctx, &business.GetBusinessReq{
	//	AdminIds: []int64{uid},
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if len(businesses.Business) == 0 {
	//	return nil, err
	//}
	//
	//log.Println("businesses:", req.Page, req.PageSize)
	//result, err := l.svcCtx.Store.GetStore(l.ctx, &store.GetStoreReq{
	//	BusinessIds: []int64{businesses.Business[0].Id},
	//	Page:        int32(req.Page),
	//	PageSize:    int32(req.PageSize),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//var storeList = make([]types.StoreInfo, 0)
	//for _, store := range result.Store {
	//	storeList = append(storeList, types.StoreInfo{
	//		ID:        store.Id,
	//		Name:      store.Name,
	//		Phone:     store.Phone,
	//		Address:   store.Address,
	//		Status:    int(store.Status),
	//		StartTime: store.StartTime,
	//		EndTime:   store.EndTime,
	//	})
	//}

	return &types.StoreListResp{
		//Total: result.Total,
		//List:  storeList,
	}, nil
}
