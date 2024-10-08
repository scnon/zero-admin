package menu

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList(req *types.MenuListReq) (resp *types.MenuListResp, err error) {
	result, err := l.svcCtx.Admin.GetMenu(l.ctx, &admin.GetMenuReq{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		TenantId: l.svcCtx.Config.Tenant,
	})

	menus := make([]types.MenuInfo, 0)
	for _, menu := range result.Menu {
		menus = append(menus, types.MenuInfo{
			ID:       menu.Id,
			Title:    menu.Title,
			ParentID: menu.ParentId,
			Path:     menu.Path,
		})
	}

	resp = &types.MenuListResp{
		Base: l.svcCtx.Success(),
		Data: types.MenuListData{
			Total: result.Total,
			List:  menus,
		},
	}
	return
}
