package logic

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"
	"zero-admin/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UpdateUserLogic) UpdateUser(in *admin.UpdateUserReq) (*admin.UpdateUserResp, error) {
	var entity model.SysUser
	copier.Copy(&entity, &in)

	if err := l.svcCtx.UserModel.Update(l.ctx, &entity); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update user err %v", err)
	}
	return &admin.UpdateUserResp{}, nil
}
