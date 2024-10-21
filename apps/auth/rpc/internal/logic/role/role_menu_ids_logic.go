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
	userId := strconv.FormatUint(in.RoleId, 10)
	domainId := strconv.FormatUint(in.TenantId, 10)
	menuIds := l.svcCtx.Casbin.GetRolesForUserInDomain(userId, domainId)
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
