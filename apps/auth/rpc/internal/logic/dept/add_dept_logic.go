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

var (
	ErrParentNotFound = errors.New("选择的父级部门不存在")
)

type AddDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDeptLogic) AddDept(in *auth.AddDeptReq) (*auth.AddDeptResp, error) {
	if in.ParentId != 0 {
		var existingDept models.SysDept
		res := l.svcCtx.DB.Where("id = ?", in.ParentId).First(&existingDept)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return nil, errors.WithStack(ErrParentNotFound)
			}
			return nil, errors.Wrapf(xerr.NewDBErr(), "查询父级部门失败 %v", res.Error)
		}
	}

	newDept := models.SysDept{
		Name:     in.Name,
		ParentId: uint(in.ParentId),
		ResModel: models.ResModel{
			Sort:      int(in.Sort),
			TenantID:  uint(in.TenantId),
			CreatorID: uint(in.Op),
		},
	}
	if err := l.svcCtx.DB.Create(&newDept).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "创建部门失败: %v", err)
	}

	return &auth.AddDeptResp{
		Id: uint64(newDept.ID),
	}, nil
}