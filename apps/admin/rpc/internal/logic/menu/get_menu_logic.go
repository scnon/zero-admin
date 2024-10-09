package menulogic

import (
	"context"
	"github.com/jinzhu/copier"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuLogic) GetMenu(in *admin.GetMenuReq) (*admin.GetMenuResp, error) {
	menuList, total, err := l.svcCtx.MenuModel.FindAll(l.ctx, in.TenantId, in.Page, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find all menu error: %v", err)
	}

	var list []*admin.MenuData
	for _, menu := range *menuList {
		data := &admin.MenuData{}
		err := copier.Copy(data, menu)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		list = append(list, data)
	}

	return &admin.GetMenuResp{
		Menu:  list,
		Total: total,
	}, nil
}
