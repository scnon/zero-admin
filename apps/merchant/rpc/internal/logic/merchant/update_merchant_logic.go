package merchantlogic

import (
	"context"

	"xlife/apps/merchant/rpc/internal/svc"
	"xlife/apps/merchant/rpc/merchant"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantLogic {
	return &UpdateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMerchantLogic) UpdateMerchant(in *merchant.MerchantInfo) (*merchant.MerchantInfo, error) {
	// todo: add your logic here and delete this line

	return &merchant.MerchantInfo{}, nil
}
