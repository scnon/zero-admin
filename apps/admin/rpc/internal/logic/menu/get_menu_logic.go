package menulogic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotFound = xerr.NewMsg("用户不存在")
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuLogic) GetMenu(in *admin.GetMenuReq) (*admin.GetMenuResp, error) {
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.AdminId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user error: %v", err)
	}

	roles, err := l.svcCtx.UserRoleModel.FindAllByUserId(l.ctx, in.AdminId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find all role error: %v", err)
	}
	roleIds := make([]int64, 0)
	for _, role := range roles {
		roleIds = append(roleIds, role.RoleId)
	}
	menus, err := l.svcCtx.RoleMenuModel.FindAllByRoleIds(l.ctx, roleIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find all role menu error: %v", err)
	}
	menuIds := make([]int64, 0)
	for _, menu := range menus {
		menuIds = append(menuIds, menu.MenuId)
	}
	menuList, err := l.svcCtx.MenuModel.FindAllByIds(l.ctx, menuIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find all menu error: %v", err)
	}

	var list []*admin.MenuData
	for _, menu := range menuList {
		data := &admin.MenuData{}
		err := copier.Copy(data, menu)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		data.Creator = menu.CreatorName.String
		if menu.Creator == 0 {
			data.Creator = "系统"
		}
		data.Updater = menu.UpdaterName.String
		data.CreateTime = menu.CreateTime.Unix()
		data.UpdateTime = menu.UpdateTime.Time.Unix()
		list = append(list, data)
	}

	return &admin.GetMenuResp{
		Menu: list,
	}, nil
}
