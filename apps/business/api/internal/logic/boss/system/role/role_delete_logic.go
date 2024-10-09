package role

import (
	"context"

	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.RoleDeleteReq) (resp *types.RoleDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
