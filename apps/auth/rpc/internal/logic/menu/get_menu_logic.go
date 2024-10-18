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
	// 1. 获取到用户
	user := models.SysUser{}
	res := l.svcCtx.DB.Where("id = ?", in.AdminId).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user error: %v", res.Error)
	}
	tenant := strconv.FormatUint(in.TenantId, 10)

	// 2. 查询到该用户的角色列表
	roles := l.svcCtx.Casbin.GetRolesForUserInDomain(strconv.FormatUint(uint64(user.ID), 10), tenant)
	// 3. 查询到该用户的资源列表
	var resources []string
	for _, role := range roles {
		policies, err := l.svcCtx.Casbin.GetFilteredPolicy(0, role, tenant, "", "read")
		if err != nil {
			return nil, err
		}
		for _, policy := range policies {
			resources = append(resources, policy[2])
		}
	}
	// 去重处理（如果角色重叠导致重复的资源）
	menuIds := make([]uint64, 0)
	for _, res := range resources {
		id, err := strconv.ParseUint(res, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "parse id error: %v", err)
		}
		menuIds = append(menuIds, id)
	}
	// 4. 查询到资源列表
	var menus []models.SysMenu
	res = l.svcCtx.DB.Where("id IN (?)", menuIds).Find(&menus)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find menu error: %v", res.Error)
	}
	var menuList []*auth.MenuData
	for _, menu := range menus {
		menuList = append(menuList, &auth.MenuData{
			Id:       uint64(menu.ID),
			Title:    menu.Title,
			ParentId: uint64(menu.ParentID),
			Path:     menu.Path,
		})
	}

	return &auth.GetMenuResp{
		Menu: menuList,
	}, nil
}
