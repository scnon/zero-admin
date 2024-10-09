package rolelogic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	perr "github.com/pkg/errors"
	"zero-admin/apps/model"
	"zero-admin/pkg/xerr"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrRoleNotFound = perr.New("角色不存在")
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *admin.UpdateRoleReq) (*admin.UpdateRoleResp, error) {
	entity := model.SysRole{}
	if err := copier.Copy(&entity, &in); err != nil {
		return nil, perr.Wrapf(xerr.NewInternalErr(), "copy entity failed: %v", err)
	}
	if _, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, perr.WithStack(ErrRoleNotFound)
		}
		return nil, perr.Wrapf(xerr.NewDBErr(), "find role error: %v", err)
	}
	if err := l.svcCtx.RoleModel.Update(l.ctx, &entity); err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "update role error: %v", err)
	}
	return &admin.UpdateRoleResp{}, nil
}
