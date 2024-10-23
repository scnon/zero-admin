package merchantlogic

import (
	"context"

	"xlife/apps/merchant/rpc/internal/svc"
	"xlife/apps/merchant/rpc/merchant"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMerchantLogic {
	return &AddMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMerchantLogic) AddMerchant(in *merchant.AddMerchantReq) (*merchant.MerchantInfo, error) {
	// todo: add your logic here and delete this line

	return &merchant.MerchantInfo{}, nil
}
