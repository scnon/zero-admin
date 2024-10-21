package logic

import (
	"context"

	"xlife/apps/model"
	"xlife/apps/product/rpc/internal/svc"
	"xlife/apps/product/rpc/product"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductLogic) AddProduct(in *product.AddProductReq) (*product.ProductInfo, error) {
	var entity model.Product
	copier.Copy(&entity, in)

	_, err := l.svcCtx.ProductModel.Insert(l.ctx, &entity)
	if err != nil {
		return nil, errors.Wrapf(err, "insert product err %v", err)
	}

	var productInfo product.ProductInfo
	copier.Copy(&productInfo, entity)
	return &productInfo, nil
}
