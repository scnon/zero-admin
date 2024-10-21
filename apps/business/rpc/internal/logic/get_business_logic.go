package logic

import (
	"context"

	"xlife/apps/business/rpc/business"
	"xlife/apps/business/rpc/internal/svc"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusinessLogic {
	return &GetBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBusinessLogic) GetBusiness(in *business.GetBusinessReq) (*business.GetBusinessResp, error) {
	entities, err := l.svcCtx.BusinessModel.FindAll(l.ctx, in.Ids, in.AdminIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find business err %v", err)
	}

	businessList := make([]*business.BusinessInfo, 0)
	for _, entity := range entities {
		businessList = append(businessList, &business.BusinessInfo{
			Id:      entity.Id,
			Phone:   entity.Phone,
			TgId:    entity.TgId,
			AdminId: entity.AdminId,
		})
	}

	return &business.GetBusinessResp{
		Business: businessList,
	}, nil
}
