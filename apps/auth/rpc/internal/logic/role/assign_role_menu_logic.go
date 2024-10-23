package rolelogic

import (
	"context"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRoleMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleMenuLogic {
	return &AssignRoleMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssignRoleMenuLogic) AssignRoleMenu(in *auth.AssignRoleMenuReq) (*auth.AssignRoleMenuResp, error) {
	// todo: add your logic here and delete this line

	return &auth.AssignRoleMenuResp{}, nil
}
