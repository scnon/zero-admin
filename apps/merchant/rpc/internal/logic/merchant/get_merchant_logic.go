package merchantlogic

import (
	"context"

	"xlife/apps/merchant/rpc/internal/svc"
	"xlife/apps/merchant/rpc/merchant"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantLogic {
	return &GetMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantLogic) GetMerchant(in *merchant.MerchantListReq) (*merchant.MerchantListResp, error) {
	// todo: add your logic here and delete this line

	return &merchant.MerchantListResp{}, nil
}
