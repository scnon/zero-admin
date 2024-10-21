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

type UpdateCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCustomerLogic) UpdateCustomer(in *customer.CustomerInfo) (*customer.CustomerInfo, error) {
	entity, err := l.svcCtx.CustomerModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find customer err %v", err)
	}

	entity.Phone = in.Phone
	entity.TgId = in.TgId
	entity.TgUsername = in.TgUsername
	entity.TgFirstName = in.TgFirstName
	entity.TgLastName = in.TgLastName
	entity.TgLanguageCode = in.TgLanguageCode
	if err := l.svcCtx.CustomerModel.Update(l.ctx, entity); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update customer err %v", err)
	}

	var customerInfo customer.CustomerInfo
	copier.Copy(&customerInfo, entity)
	return &customerInfo, nil
}
