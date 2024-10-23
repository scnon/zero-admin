package rolelogic

import (
	"context"
	"errors"
	perr "github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"xlife/models"

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
	// 1. 检查角色是否存在
	var roleEntity models.SysRole
	res := l.svcCtx.DB.Where("id = ?", in.RoleId).Where("tenant_id = ?", in.TenantId).First(&roleEntity)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrRoleNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询角色失败: %v", res.Error)
	}
	// 2. 查询菜单
	var menus []models.SysMenu
	res = l.svcCtx.DB.Where("id IN ?", in.MenuIds).Where("tenant_id = ?", in.TenantId).Find(&menus)
	if res.Error != nil {
		return nil, perr.Wrapf(res.Error, "查询菜单失败: %v", res.Error)
	}
	// 2.1 如果菜单为空 直接返回
	if len(menus) == 0 {
		return &auth.AssignRoleMenuResp{}, nil
	}
	// 3. 指定菜单
	var policies [][]string
	for _, menu := range menus {
		tenant := strconv.FormatUint(in.TenantId, 10)
		menuId := strconv.FormatUint(uint64(menu.ID), 10)
		policy := []string{roleEntity.Name, tenant, menuId, "read"}
		policies = append(policies, policy)
	}
	ok, err := l.svcCtx.Casbin.AddPolicies(policies)
	if err != nil || !ok {
		return nil, perr.Wrapf(err, "指定菜单失败: %v", err)
	}
	return &auth.AssignRoleMenuResp{}, nil
}
