package rolelogic

import (
	"context"
	"strconv"

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
	// 1. 类型转换(casbin 中的用户ID 和 角色 ID 是字符串)
	userId := strconv.FormatUint(in.RoleId, 10)
	domainId := strconv.FormatUint(in.TenantId, 10)
	// 2. 查询角色的菜单权限
	menuIds := l.svcCtx.Casbin.GetRolesForUserInDomain(userId, domainId)
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
