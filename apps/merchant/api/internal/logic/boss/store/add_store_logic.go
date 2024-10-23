package store

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
)

type AddStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStoreLogic {
	return &AddStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStoreLogic) AddStore(req *types.AddStoreReq) (resp *types.StoreInfo, err error) {
	//uid := ctxdata.GetUId(l.ctx)
	//if uid == 0 {
	//	return nil, xerr.NewMsg("用户未登录")
	//}
	//businesses, err := l.svcCtx.Business.GetBusiness(l.ctx, &merchant.GetBusinessReq{
	//	AdminIds: []int64{uid},
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if len(businesses.Business) == 0 {
	//	return nil, xerr.NewMsg("商户不存在")
	//}
	//merchant := businesses.Business[0]

	//result, err := l.svcCtx.Store.AddStore(l.ctx, &store.AddStoreReq{
	//	BusinessId: merchant.Id,
	//	Name:       req.Name,
	//	Phone:      req.Phone,
	//	Status:     int32(req.Status),
	//	StartTime:  req.StartTime,
	//	EndTime:    req.EndTime,
	//	Address:    req.Address,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//log.Println(result.Id)
	var storeInfo types.StoreInfo
	//copier.Copy(&storeInfo, &result)
	return &storeInfo, nil
}
