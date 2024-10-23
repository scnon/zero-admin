package rolelogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"xlife/models"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuIdsLogic {
	return &RoleMenuIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleMenuIdsLogic) RoleMenuIds(in *auth.RoleMenuIdsReq) (*auth.RoleMenuIdsResp, error) {
	// 1. 查询角色
	var role models.SysRole
	if res := l.svcCtx.DB.Where("id = ?", in.RoleId).First(&role); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(ErrRoleNotFound)
		}
		return nil, errors.Wrapf(res.Error, "查询角色失败: %v", res.Error)
	}
	// 2. 查询角色的菜单权限
	domainId := strconv.FormatUint(in.TenantId, 10)
	menuIds := l.svcCtx.Casbin.GetRolesForUserInDomain(role.Name, domainId)
	// 3. 类型转换
	ids := make([]uint64, 0)
	for _, menuId := range menuIds {
		id, err := strconv.ParseUint(menuId, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return &auth.RoleMenuIdsResp{
		MenuIds: ids,
	}, nil
}
