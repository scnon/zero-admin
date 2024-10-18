package role

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	uid := ctxdata.GetUId(l.ctx)
	roles, err := l.svcCtx.User.GetRoles(l.ctx, &auth.GetRolesReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	roleList := make([]types.RoleData, 0)
	for _, role := range roles.Roles {
		roleList = append(roleList, types.RoleData{
			Id:     int64(role.Id),
			Name:   role.Name,
			Remark: role.Remark,
			Sort:   int64(role.Sort),
		})
	}
	return &types.RoleListResp{
		Base: l.svcCtx.Success(),
		Data: types.RoleList{
			List: roleList,
		},
	}, nil
}
