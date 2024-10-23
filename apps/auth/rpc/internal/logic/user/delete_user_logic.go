package userlogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"
	"xlife/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"xlife/apps/auth/rpc/internal/svc"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *auth.DeleteUserReq) (*auth.DeleteUserResp, error) {
	// 1. 使用事务逻辑删除用户
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// Step 1: 更新 DeleterID
		if err := tx.Model(&models.SysUser{}).
			Where("id IN (?)", in.Ids).
			Where("tenant_id = ?", in.TenantId).
			Update("deleter_id", in.Op).Error; err != nil {
			return err
		}

		// Step 2: 逻辑删除
		return tx.Where("id IN (?)", in.Ids).Where("tenant_id = ?", in.TenantId).Delete(&models.SysUser{}).Error
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除用户失败: %v", err)
	}
	return &auth.DeleteUserResp{}, nil
}
