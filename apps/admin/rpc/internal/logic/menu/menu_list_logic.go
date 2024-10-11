package menulogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"zero-admin/pkg/xerr"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *admin.MenuListReq) (*admin.MenuListResp, error) {
	entityList, total, err := l.svcCtx.MenuModel.FindAll(l.ctx, in.TenantId, in.Page, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find menu err %v", err)
	}

	menuList := make([]*admin.MenuData, 0)
	for _, entity := range entityList {
		var menu admin.MenuData
		if err := copier.Copy(&menu, &entity); err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		menu.CreateTime = entity.CreateTime.Unix()
		menu.UpdateTime = entity.UpdateTime.Time.Unix()
		menu.Creator = entity.CreatorName.String
		if entity.Creator == 0 {
			menu.Creator = "系统"
		}
		menu.Updater = entity.UpdaterName.String
		menuList = append(menuList, &menu)
	}

	return &admin.MenuListResp{
		List:  menuList,
		Total: total,
	}, nil
}
