package logic

import (
	"context"

	"xlife/apps/product/rpc/internal/svc"
	"xlife/apps/product/rpc/product"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.GetProductResp, error) {
	entities, err := l.svcCtx.ProductModel.FindAll(l.ctx, in.Ids, in.BusinessIds, in.StoreIds, in.CateIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find product err %v", err)
	}

	prodocts := make([]*product.ProductInfo, 0)
	for _, entity := range entities {
		prodocts = append(prodocts, &product.ProductInfo{
			Id:      entity.Id,
			Name:    entity.Name,
			StoreId: entity.StoreId,
			CateId:  entity.CateId,
			Price:   entity.Price,
			Unit:    entity.Unit,
			Stock:   entity.Stock,
		})
	}

	return &product.GetProductResp{
		Product: prodocts,
	}, nil
}
