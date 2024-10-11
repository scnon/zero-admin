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
	ErrParentNotFound = errors.New("选择的父级菜单不存在")
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMenuLogic) AddMenu(in *admin.AddMenuReq) (*admin.AddMenuResp, error) {
	if in.ParentId != 0 {
		if _, err := l.svcCtx.MenuModel.FindOne(l.ctx, in.ParentId); err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, errors.WithStack(ErrParentNotFound)
			}
			return nil, errors.Wrapf(xerr.NewDBErr(), "find one menu error: %v", err)
		}
	}

	entity := &model.SysMenu{}
	err := copier.Copy(entity, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
	}

	if _, err := l.svcCtx.MenuModel.Insert(l.ctx, entity); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert menu error: %v", err)
	}

	return &admin.AddMenuResp{
		Id: entity.Id,
	}, nil
}
