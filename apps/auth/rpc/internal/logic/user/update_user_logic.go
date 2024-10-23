package userlogic

import (
	"context"
	"errors"
	perr "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *auth.UpdateUserReq) (*auth.UpdateUserResp, error) {
	res := l.svcCtx.DB.Where("id = ?", in.Id).Where("tenant_id = ?", in.TenantId).First(&models.SysUser{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrUserNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询用户失败: %v", res.Error)
	}

	updater := uint(in.Op)
	res = l.svcCtx.DB.Where("id = ?", in.Id).Where("tenant_id = ?", in.TenantId).Updates(&models.SysUser{
		Username: in.Username,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		ResModel: models.ResModel{
			Sort:      in.Sort,
			Remark:    in.Remark,
			UpdaterID: &updater,
		},
	})
	if res.Error != nil {
		return nil, perr.Wrapf(res.Error, "更新用户失败: %v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, perr.Wrapf(errors.New("没有更新数据"), "更新用户失败: %v", res.Error)
	}

	return &auth.UpdateUserResp{}, nil
}
