package store

import (
	"context"

	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoreLogic {
	return &DeleteStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStoreLogic) DeleteStore(req *types.DeleteStoreReq) error {
	// todo: add your logic here and delete this line

	return nil
}
