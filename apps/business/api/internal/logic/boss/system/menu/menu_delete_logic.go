package menu

import (
	"context"
	"xlife/apps/auth/rpc/admin"

	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req *types.MenuDeleteReq) (resp *types.MenuDeleteResp, err error) {
	_, err = l.svcCtx.Menu.DeleteMenu(l.ctx, &admin.DeleteMenuReq{
		Ids:      req.Ids,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	return &types.MenuDeleteResp{
		Base: l.svcCtx.Success(),
	}, nil
}
