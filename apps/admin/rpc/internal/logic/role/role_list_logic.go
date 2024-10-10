package rolelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"zero-admin/pkg/xerr"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *admin.RoleListReq) (*admin.RoleListResp, error) {
	entityList, total, err := l.svcCtx.RoleModel.FindAll(l.ctx, in.TenantId, in.Page, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find role err %v", err)
	}

	roleList := make([]*admin.RoleData, 0)
	for _, entity := range entityList {
		var role admin.RoleData
		if err := copier.Copy(&role, &entity); err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		role.CreateTime = entity.CreateTime.Unix()
		role.UpdateTime = entity.UpdateTime.Unix()
		role.Creator = entity.CreatorName.String
		if entity.Creator == 0 {
			role.Creator = "系统"
		}
		role.Updater = entity.UpdaterName.String
		roleList = append(roleList, &role)
	}

	return &admin.RoleListResp{
		List:  roleList,
		Total: total,
	}, nil
}
