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
	// 1. 查询用户是否存在
	var user models.SysUser
	res := l.svcCtx.DB.Where("id = ?", in.UserId).Where("tenant_id = ?", in.TenantId).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user error: %v", res.Error)
	}
	// 2. 查询用户角色
	tenantId := strconv.FormatUint(in.TenantId, 10)
	roles := l.svcCtx.Casbin.GetRolesForUserInDomain(user.Username, tenantId)
	// 3. 查询角色的菜单列表
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
	// 4. 查询菜单
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
	// 5. 构造返回数据
	var list []*auth.MenuData
	for _, menu := range menus {
		data := &auth.MenuData{
			Id:       uint64(menu.ID),
			ParentId: uint64(menu.ParentID),
			Title:    menu.Title,
			Path:     menu.Path,
			Sort:     menu.Sort,
		}
		list = append(list, data)
	}
	return &auth.GetMenuResp{
		Menu: list,
	}, nil
}
