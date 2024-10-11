package rolelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *admin.AddRoleReq) (*admin.AddRoleResp, error) {
	entity := model.SysRole{}
	err := copier.Copy(&entity, &in)
	if err != nil {
		return nil, err
	}
	id, err := l.svcCtx.RoleModel.InsertWithMenus(l.ctx, &entity, in.Menus)
	if err != nil {
		return nil, err
	}

	return &admin.AddRoleResp{
		Id: id,
	}, nil
}
