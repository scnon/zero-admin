package logic

import (
	"context"

	"xlife/apps/business/rpc/business"
	"xlife/apps/business/rpc/internal/svc"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBusinessLogic {
	return &DeleteBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBusinessLogic) DeleteBusiness(in *business.DeleteBusinessReq) (*business.DeleteBusinessResp, error) {
	if err := l.svcCtx.BusinessModel.DeleteAll(l.ctx, in.Ids); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete business err %v", err)
	}
	return &business.DeleteBusinessResp{}, nil
}
