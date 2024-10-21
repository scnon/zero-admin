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

type AddCateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCateLogic {
	return &AddCateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCateLogic) AddCate(in *product.AddCateReq) (*product.CateInfo, error) {
	var entity model.Cate
	copier.Copy(&entity, in)

	_, err := l.svcCtx.CateModel.Insert(l.ctx, &entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert cate err %v", err)
	}

	var cateInfo product.CateInfo
	copier.Copy(&cateInfo, entity)
	return &cateInfo, nil
}
