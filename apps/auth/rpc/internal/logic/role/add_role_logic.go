package rolelogic

import (
	"context"
	"github.com/pkg/errors"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
	"xlife/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *auth.AddRoleReq) (*auth.AddRoleResp, error) {
	// 1. 创建角色
	newRole := models.SysRole{
		Name: in.Name,
		ResModel: models.ResModel{
			Sort:      int(in.Sort),
			Remark:    in.Remark,
			TenantID:  uint(in.TenantId),
			CreatorID: uint(in.Op),
		},
	}
	if err := l.svcCtx.DB.Create(&newRole).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "创建用户失败: %v", err)
	}

	return &auth.AddRoleResp{
		Id: uint64(newRole.ID),
	}, nil
}
