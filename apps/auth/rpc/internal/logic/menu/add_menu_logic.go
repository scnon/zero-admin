package menulogic

import (
	"context"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
	"xlife/pkg/xerr"

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

func (l *AddMenuLogic) AddMenu(in *auth.AddMenuReq) (*auth.AddMenuResp, error) {
	// 1. 查询父级菜单是否存在
	if in.ParentId != 0 {
		var existingMenu models.SysMenu
		res := l.svcCtx.DB.Where("id = ?", in.ParentId).Where("tenant_id = ?", in.TenantId).First(&existingMenu)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return nil, errors.WithStack(ErrParentNotFound)
			}
			return nil, errors.Wrapf(xerr.NewDBErr(), "查询父级菜单失败 %v", res.Error)
		}
	}
	// 2. 创建新菜单
	newMenu := models.SysMenu{
		Title:     in.Name,
		Path:      in.Path,
		Component: in.Component,
		ParentID:  uint(in.ParentId),
		ResModel: models.ResModel{
			Sort:      in.Sort,
			TenantID:  uint(in.TenantId),
			CreatorID: uint(in.Op),
		},
	}
	if err := l.svcCtx.DB.Create(&newMenu).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "创建菜单失败: %v", err)
	}
	return &auth.AddMenuResp{
		Id: uint64(newMenu.ID),
	}, nil
}
