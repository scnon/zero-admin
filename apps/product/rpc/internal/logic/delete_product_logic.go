package logic

import (
	"context"

	"zero-admin/apps/model"
	"zero-admin/apps/product/rpc/internal/svc"
	"zero-admin/apps/product/rpc/product"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *product.ProductInfo) (*product.ProductInfo, error) {
	entity, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find product err %v", err)
	}

	err = l.svcCtx.ProductModel.Delete(l.ctx, entity.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete product err %v", err)
	}
	return &product.ProductInfo{}, nil
}
