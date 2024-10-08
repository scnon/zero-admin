package menu

import (
	"context"
	"zero-admin/apps/admin/rpc/admin"

	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.MenuUpdateReq) (resp *types.MenuUpdateResp, err error) {
	_, err = l.svcCtx.Admin.UpdateMenu(l.ctx, &admin.UpdateMenuReq{
		Id:       req.ID,
		ParentId: req.ParentID,
		Path:     req.Path,
		Title:    req.Title,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	return &types.MenuUpdateResp{
		Base: l.svcCtx.Success(),
	}, nil
}
