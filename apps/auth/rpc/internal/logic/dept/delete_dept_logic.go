package deptlogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/models"
	"xlife/pkg/xerr"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeptLogic {
	return &DeleteDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeptLogic) DeleteDept(in *auth.DeleteDeptReq) (*auth.DeleteDeptResp, error) {
	// 1. 使用事务逻辑删除部门
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// Step 1: 更新 DeleterID
		if err := tx.Model(&models.SysDept{}).
			Where("id IN (?)", in.Ids).
			Update("deleter_id", in.Op).Error; err != nil {
			return err
		}

		// Step 2: 逻辑删除
		if err := tx.Where("id IN (?)", in.Ids).Delete(&models.SysDept{}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "删除部门失败: %v", err)
	}

	return &auth.DeleteDeptResp{}, nil
}
