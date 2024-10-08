package logic

import (
	"context"

	"zero-admin/apps/store/rpc/internal/svc"
	"zero-admin/apps/store/rpc/store"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoreLogic {
	return &DeleteStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStoreLogic) DeleteStore(in *store.DeleteStoreReq) (*store.DeleteStoreResp, error) {
	err := l.svcCtx.StoreModel.DeleteAll(l.ctx, in.Ids, in.BusinessIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete store err %v", err)
	}

	return &store.DeleteStoreResp{}, nil
}
