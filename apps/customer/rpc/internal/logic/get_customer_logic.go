package logic

import (
	"context"

	"xlife/apps/customer/rpc/customer"
	"xlife/apps/customer/rpc/internal/svc"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerLogic {
	return &GetCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCustomerLogic) GetCustomer(in *customer.GetCustomerReq) (*customer.GetCustomerResp, error) {
	entities, err := l.svcCtx.CustomerModel.FindAll(l.ctx, in.Ids)
	if err != nil {
		return nil, errors.Wrapf(err, "find customer err %v", err)
	}

	customerList := make([]*customer.CustomerInfo, 0)
	for _, entity := range entities {
		var customerInfo customer.CustomerInfo
		copier.Copy(&customerInfo, entity)
		customerList = append(customerList, &customerInfo)
	}

	return &customer.GetCustomerResp{
		Customer: customerList,
	}, nil
}
