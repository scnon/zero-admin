package menulogic

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *admin.DeleteMenuReq) (*admin.DeleteMenuResp, error) {
	err := l.svcCtx.MenuModel.DeleteAll(l.ctx, in.Ids)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete menu error: %v", err)
	}
	return &admin.DeleteMenuResp{}, nil
}
