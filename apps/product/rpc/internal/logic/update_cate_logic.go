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

type UpdateCateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCateLogic {
	return &UpdateCateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCateLogic) UpdateCate(in *product.CateInfo) (*product.CateInfo, error) {
	entity, err := l.svcCtx.CateModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find cate err %v", err)
	}

	entity.Name = in.Name
	entity.StoreId = in.StoreId
	entity.Sort = int64(in.Sort)
	err = l.svcCtx.CateModel.Update(l.ctx, entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update cate err %v", err)
	}

	var cateInfo product.CateInfo
	copier.Copy(&cateInfo, entity)
	return &cateInfo, nil
}
