package rolelogic

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

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *auth.DeleteRoleReq) (*auth.DeleteRoleResp, error) {
	// 1. 使用事务逻辑删除角色
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// Step 1: 更新 DeleterID
		if err := tx.Model(&models.SysRole{}).
			Where("id IN (?)", in.Ids).
			Where("tenant_id = ?", in.TenantId).
			Update("deleter_id", in.Op).Error; err != nil {
			return err
		}

		// Step 2: 逻辑删除
		return tx.Where("id IN (?)", in.Ids).Where("tenant_id = ?", in.TenantId).
			Delete(&models.SysRole{}).Error
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除角色失败: %v", err)
	}
	return &auth.DeleteRoleResp{}, nil
}
