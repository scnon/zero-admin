package userlogic

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *admin.DeleteUserReq) (*admin.DeleteUserReq, error) {
	if err := l.svcCtx.UserModel.DeleteAll(l.ctx, in.Ids); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete user err %v", err)
	}
	return &admin.DeleteUserReq{}, nil
}
