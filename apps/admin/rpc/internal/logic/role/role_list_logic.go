package rolelogic

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *admin.RoleListReq) (*admin.RoleListResp, error) {
	entityList, total, err := l.svcCtx.RoleModel.FindAll(l.ctx, in.Page, in.PageSize)

	return &admin.RoleListResp{}, nil
}
