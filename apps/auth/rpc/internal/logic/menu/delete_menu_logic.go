package menulogic

import (
	"context"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"

	"xlife/apps/auth/rpc/internal/svc"
	"xlife/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *auth.DeleteMenuReq) (*auth.DeleteMenuResp, error) {
	// 1. 使用事务逻辑删除菜单
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// Step 1: 更新 DeleterID
		if err := tx.Model(&models.SysMenu{}).
			Where("id IN (?)", in.Ids).
			Where("tenant_id = ?", in.TenantId).
			Update("deleter_id", in.Op).Error; err != nil {
			return err
		}
		// Step 2: 逻辑删除
		return tx.Where("id IN (?)", in.Ids).Where("tenant_id = ?", in.TenantId).Delete(&models.SysMenu{}).Error
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除菜单失败: %v", err)
	}
	return &auth.DeleteMenuResp{}, nil
}
