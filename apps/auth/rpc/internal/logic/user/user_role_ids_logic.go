package userlogic

import (
	"context"
	"strconv"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleIdsLogic {
	return &UserRoleIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleIdsLogic) UserRoleIds(in *auth.UserRoleIdsReq) (*auth.UserRoleIdsResp, error) {
	userId := strconv.FormatUint(in.UserId, 10)
	domainId := strconv.FormatUint(in.TenantId, 10)
	roleIds := l.svcCtx.Casbin.GetUsersForRoleInDomain(userId, domainId)
	ids := make([]uint64, 0)
	for _, roleId := range roleIds {
		id, err := strconv.ParseUint(roleId, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return &auth.UserRoleIdsResp{
		RoleIds: ids,
	}, nil
}
