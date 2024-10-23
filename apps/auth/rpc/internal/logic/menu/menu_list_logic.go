package menulogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"
	"xlife/pkg/xerr"

	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *auth.MenuListReq) (*auth.MenuListResp, error) {
	// 1. 根据条件查询菜单列表
	var menus []models.SysMenu
	res := l.makeQuery(in).Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&menus)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户失败 %v", res.Error)
	}

	// 2. 构造返回数据
	var list []*auth.MenuData
	for _, menu := range menus {
		data := &auth.MenuData{
			Id:       uint64(menu.ID),
			Title:    menu.Title,
			Path:     menu.Path,
			ParentId: uint64(menu.ParentID),
			Sort:     int32(menu.Sort),
		}
		if menu.Creator != nil {
			data.Creator = menu.Creator.Username
		}
		if menu.Updater != nil {
			data.Updater = menu.Updater.Username
		}
		list = append(list, data)
	}

	var total int64
	res = l.makeQuery(in).Count(&total)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户总数失败 %v", res.Error)
	}

	return &auth.MenuListResp{
		List:  list,
		Total: uint64(total),
	}, nil
}

func (l *MenuListLogic) makeQuery(in *auth.MenuListReq) *gorm.DB {
	query := l.svcCtx.DB.Model(&models.SysUser{}).Preload("Creator").Preload("Updater")
	if in.ParentId != nil {
		query = query.Where("id = ?", in.ParentId)
	}
	if in.Status != nil {
		query = query.Where("status = ?", in.Status)
	}
	if in.TenantId != nil {
		query = query.Where("tenant_id = ?", in.TenantId)
	}
	return query
}
