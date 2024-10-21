package deptlogic

import (
	"context"
	"errors"
	perr "github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/models"
	"xlife/pkg/xerr"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrDeptNotFound = xerr.NewMsg("部门不存在")
)

type UpdateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeptLogic) UpdateDept(in *auth.UpdateDeptReq) (*auth.UpdateDeptResp, error) {
	res := l.svcCtx.DB.Where("id = ?", in.Id).First(&models.SysDept{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrDeptNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询菜单失败: %v", res.Error)
	}

	updater := uint(in.Op)
	res = l.svcCtx.DB.Where("id = ?", in.Id).Updates(&models.SysDept{
		Name:     in.Name,
		ParentId: uint(in.ParentId),
		ResModel: models.ResModel{
			Sort:      int(in.Sort),
			TenantID:  uint(in.TenantId),
			UpdaterID: &updater,
		},
	})
	if res.Error != nil {
		return nil, perr.Wrapf(res.Error, "更新部门信息失败: %v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, perr.Wrapf(errors.New("没有更新数据"), "更新部门信息失败: %v", res.Error)
	}

	return &auth.UpdateDeptResp{}, nil
}
