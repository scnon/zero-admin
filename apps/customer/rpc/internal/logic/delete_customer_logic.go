package logic

import (
	"context"

	"xlife/apps/customer/rpc/customer"
	"xlife/apps/customer/rpc/internal/svc"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrNotFound = xerr.NewMsg("记录不存在")
)

type DeleteCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomerLogic {
	return &DeleteCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCustomerLogic) DeleteCustomer(in *customer.DeleteCustomerReq) (*customer.DeleteCustomerResp, error) {
	if err := l.svcCtx.CustomerModel.DeleteAll(l.ctx, in.Ids); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete customer err %v", err)
	}

	return &customer.DeleteCustomerResp{}, nil
}
