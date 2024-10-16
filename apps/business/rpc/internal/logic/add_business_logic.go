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

type AddBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBusinessLogic {
	return &AddBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBusinessLogic) AddBusiness(in *business.AddBusinessReq) (*business.BusinessInfo, error) {
	entity := model.Business{
		Phone:   in.Phone,
		TgId:    in.TgId,
		AdminId: in.AdminId,
	}
	_, err := l.svcCtx.BusinessModel.Insert(l.ctx, &entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert business err %v", err)
	}

	var businessInfo business.BusinessInfo
	copier.Copy(&businessInfo, entity)
	return &businessInfo, nil
}
