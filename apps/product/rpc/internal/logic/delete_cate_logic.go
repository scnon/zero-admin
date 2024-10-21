package logic

import (
	"context"

	"xlife/apps/product/rpc/internal/svc"
	"xlife/apps/product/rpc/product"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrNotFound = errors.New("记录不存在")
)

type DeleteCateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCateLogic {
	return &DeleteCateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCateLogic) DeleteCate(in *product.DeleteCateReq) (*product.DeleCateResp, error) {
	if err := l.svcCtx.CateModel.DeleteAll(l.ctx, in.Ids); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete cate err %v", err)
	}
	return &product.DeleCateResp{}, nil
}
