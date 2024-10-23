package rolelogic

import (
	"context"
	"errors"
	perr "github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"
	"xlife/pkg/xerr"

	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrRoleNotFound = xerr.NewMsg("角色不存在")
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

func (l *UpdateRoleLogic) UpdateRole(in *auth.UpdateRoleReq) (*auth.UpdateRoleResp, error) {
	// 1. 查询要更新的角色是否存在
	res := l.svcCtx.DB.Where("id = ?", in.Id).First(&models.SysRole{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrRoleNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询菜单失败: %v", res.Error)
	}
	// 2. 更新角色信息
	updater := uint(in.Op)
	res = l.svcCtx.DB.Where("id = ?", in.Id).Updates(&models.SysRole{
		Name: in.Name,
		ResModel: models.ResModel{
			Sort:      int(in.Sort),
			Remark:    in.Remark,
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
	return &auth.UpdateRoleResp{}, nil
}
