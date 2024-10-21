package logic

import (
	"context"

	"xlife/apps/product/rpc/internal/svc"
	"xlife/apps/product/rpc/product"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCateLogic {
	return &GetCateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCateLogic) GetCate(in *product.GetCateReq) (*product.GetCateResp, error) {
	cates, err := l.svcCtx.CateModel.FindAll(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find cates err %v", err)
	}

	cateInfos := make([]*product.CateInfo, 0)
	for _, cate := range cates {
		cateInfos = append(cateInfos, &product.CateInfo{
			Id:   cate.Id,
			Name: cate.Name,
		})
	}
	return &product.GetCateResp{
		Cate: cateInfos,
	}, nil
}
