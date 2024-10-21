package logic

import (
	"context"

	"xlife/apps/model"
	"xlife/apps/product/rpc/internal/svc"
	"xlife/apps/product/rpc/product"
	"xlife/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.ProductInfo) (*product.ProductInfo, error) {
	entity, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find product err %v", err)
	}

	entity.Name = in.Name
	entity.Image = in.Image
	entity.StoreId = in.StoreId
	entity.CateId = in.CateId
	entity.Price = in.Price
	entity.Unit = in.Unit
	entity.Stock = in.Stock
	err = l.svcCtx.ProductModel.Update(l.ctx, entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update product err %v", err)
	}

	var productInfo product.ProductInfo
	copier.Copy(&productInfo, entity)
	return &productInfo, nil
}
