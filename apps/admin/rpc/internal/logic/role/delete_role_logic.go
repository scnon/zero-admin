package rolelogic

import (
	"context"
	"github.com/pkg/errors"
	"zero-admin/pkg/xerr"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *admin.DeleteRoleReq) (*admin.DeleteRoleResp, error) {
	err := l.svcCtx.RoleModel.DeleteAll(l.ctx, in.Ids)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "delete role error: %v", err)
	}
	return &admin.DeleteRoleResp{}, nil
}
