package logic

import (
	"context"

	"xlife/apps/business/rpc/business"
	"xlife/apps/business/rpc/internal/svc"
	"xlife/apps/model"
	"xlife/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrNotFound = xerr.NewMsg("商家不存在")
)

type UpdateBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBusinessLogic {
	return &UpdateBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBusinessLogic) UpdateBusiness(in *business.BusinessInfo) (*business.BusinessInfo, error) {
	entity, err := l.svcCtx.BusinessModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.WithStack(ErrNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find business err %v", err)
	}

	entity.Phone = in.Phone
	entity.TgId = in.TgId
	entity.AdminId = in.AdminId
	err = l.svcCtx.BusinessModel.Update(l.ctx, entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update business err %v", err)
	}

	var businessInfo business.BusinessInfo
	copier.Copy(&businessInfo, entity)
	return &businessInfo, nil
}
