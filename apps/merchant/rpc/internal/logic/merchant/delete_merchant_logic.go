package merchantlogic

import (
	"context"

	"xlife/apps/merchant/rpc/internal/svc"
	"xlife/apps/merchant/rpc/merchant"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantLogic {
	return &DeleteMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMerchantLogic) DeleteMerchant(in *merchant.DeleteMerchantReq) (*merchant.DeleteMerchantResp, error) {
	// todo: add your logic here and delete this line

	return &merchant.DeleteMerchantResp{}, nil
}
