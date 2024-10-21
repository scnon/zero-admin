package logic

import (
	"context"

	"xlife/apps/customer/rpc/customer"
	"xlife/apps/customer/rpc/internal/svc"
	"xlife/apps/model"
	"xlife/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCustomerLogic {
	return &AddCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCustomerLogic) AddCustomer(in *customer.AddCustomerReq) (*customer.CustomerInfo, error) {
	var entity model.Customer
	copier.Copy(&entity, in)

	_, err := l.svcCtx.CustomerModel.Insert(l.ctx, &entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert customer err %v", err)
	}

	var customerInfo customer.CustomerInfo
	copier.Copy(&customerInfo, entity)
	return &customerInfo, nil
}
