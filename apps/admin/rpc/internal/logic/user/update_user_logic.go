package userlogic

import (
	"context"
	"zero-admin/pkg/encrypt"

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
	err := copier.Copy(&entity, &in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
	}

	oldEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}

	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "gen password err %v", err)
		}

		entity.Password = string(genPassword)
	} else {
		entity.Password = oldEntity.Password
	}

	entity.Creator = oldEntity.Creator
	entity.TenantId = oldEntity.TenantId
	entity.CreateTime = oldEntity.CreateTime
	if err := l.svcCtx.UserModel.Update(l.ctx, &entity); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update user err %v", err)
	}
	return &admin.UpdateUserResp{}, nil
}
