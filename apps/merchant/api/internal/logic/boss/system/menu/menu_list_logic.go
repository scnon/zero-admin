package menu

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

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
	// 1. 获取菜单列表
	result, err := l.svcCtx.Menu.GetMenu(l.ctx, &auth.GetMenuReq{
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	// 2. 返回菜单列表
	menus := make([]types.MenuData, 0)
	for _, menu := range result.Menu {
		menus = append(menus, types.MenuData{
			ID:       menu.Id,
			Title:    menu.Title,
			ParentID: menu.ParentId,
			Path:     menu.Path,
		})
	}

	return &types.MenuListResp{
		Base: l.svcCtx.Success(),
		Data: menus,
	}, nil
}
