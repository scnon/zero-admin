package menulogic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrMenuNotFound = xerr.NewMsg("菜单不存在")
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *admin.UpdateMenuReq) (*admin.UpdateMenuResp, error) {
	entity := &model.SysMenu{}
	err := copier.Copy(entity, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
	}

	oldEntity, err := l.svcCtx.MenuModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.WithStack(ErrMenuNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find menu error: %v", err)
	}
	entity.Creator = oldEntity.Creator
	entity.TenantId = oldEntity.TenantId
	entity.CreateTime = oldEntity.CreateTime
	err = l.svcCtx.MenuModel.Update(l.ctx, entity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update menu error: %v", err)
	}
	return &admin.UpdateMenuResp{}, nil
}
