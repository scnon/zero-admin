package menulogic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-admin/apps/model"
	"zero-admin/pkg/utils"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotFound = xerr.NewMsg("菜单不存在")
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
	userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, in.AdminId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user error: %v", err)
	}

	roleIds := utils.GetInt64ArryFromStr(userEntity.Roles)
	menuList, err := l.svcCtx.MenuModel.FindAllByIds(l.ctx, roleIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find all menu error: %v", err)
	}

	var list []*admin.MenuData
	for _, menu := range menuList {
		data := &admin.MenuData{}
		err := copier.Copy(data, menu)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		data.Creator = menu.CreatorName.String
		if menu.Creator == 0 {
			data.Creator = "系统"
		}
		data.Updater = menu.UpdaterName.String
		data.CreateTime = menu.CreateTime.Unix()
		data.UpdateTime = menu.UpdateTime.Unix()
		list = append(list, data)
	}

	return &admin.GetMenuResp{
		Menu: list,
	}, nil
}
