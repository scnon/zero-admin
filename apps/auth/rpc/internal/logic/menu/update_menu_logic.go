package menulogic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
	"xlife/pkg/xerr"

	perr "github.com/pkg/errors"
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

func (l *UpdateMenuLogic) UpdateMenu(in *auth.UpdateMenuReq) (*auth.UpdateMenuResp, error) {
	res := l.svcCtx.DB.Where("id = ?", in.Id).First(&models.SysMenu{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrMenuNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询菜单失败: %v", res.Error)
	}

	updater := uint(in.Op)
	res = l.svcCtx.DB.Where("id = ?", in.Id).Updates(&models.SysMenu{
		Title:    in.Title,
		ParentID: uint(in.ParentId),
		Path:     in.Path,
		ResModel: models.ResModel{
			Sort:      int(in.Sort),
			TenantID:  uint(in.TenantId),
			UpdaterID: &updater,
		},
	})
	if res.Error != nil {
		return nil, perr.Wrapf(res.Error, "更新菜单失败: %v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, perr.Wrapf(errors.New("没有更新数据"), "更新菜单失败: %v", res.Error)
	}
	return &auth.UpdateMenuResp{}, nil
}
