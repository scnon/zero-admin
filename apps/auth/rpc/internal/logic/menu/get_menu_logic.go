package menulogic

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
	"xlife/pkg/xerr"

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

func (l *GetMenuLogic) GetMenu(in *auth.GetMenuReq) (*auth.GetMenuResp, error) {
	res := l.svcCtx.DB.Where("id = ?", in.AdminId).First(&models.SysUser{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user error: %v", res.Error)
	}
	userId := strconv.FormatUint(in.AdminId, 10)
	tenantId := strconv.FormatUint(in.TenantId, 10)
	roles := l.svcCtx.Casbin.GetRolesForUserInDomain(userId, tenantId)

	resourceMap := make(map[string]bool)
	for _, role := range roles {
		policies, err := l.svcCtx.Casbin.GetFilteredPolicy(0, role, tenantId, "", "read")
		if err == nil {
			for _, policy := range policies {
				if len(policy) >= 4 {
					// 资源位于策略的第三列 (obj)
					resource := policy[2]
					resourceMap[resource] = true
				}
			}
		}
	}
	var resources []string
	for resource := range resourceMap {
		resources = append(resources, resource)
	}

	var menuIds []uint64
	for _, menuStr := range resources {
		menuId, err := strconv.ParseUint(menuStr, 10, 64)
		if err == nil {
			menuIds = append(menuIds, menuId)
		}
	}
	var menus []models.SysMenu
	res = l.svcCtx.DB.Where("id in (?)", menuIds).Find(&menus)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find all menu error: %v", res.Error)
	}

	var list []*auth.MenuData
	for _, menu := range menus {
		data := &auth.MenuData{
			Id:       uint64(menu.ID),
			Title:    menu.Title,
			Path:     menu.Path,
			ParentId: uint64(menu.ParentID),
			Sort:     int32(menu.Sort),
		}
		if menu.Creator != nil {
			data.Creator = menu.Creator.Username
		}
		if menu.Updater != nil {
			data.Updater = menu.Updater.Username
		}
		list = append(list, data)
	}

	return &auth.GetMenuResp{
		Menu: list,
	}, nil
}
