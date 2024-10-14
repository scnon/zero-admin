package userlogic

import (
	"context"
	"zero-admin/ent/sysuser"
	"zero-admin/pkg/encrypt"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/xerr"

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
	oldEntity, err := l.svcCtx.Ent.SysUser.Query().Where(sysuser.IDEQ(in.Id)).Only(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}
	if oldEntity == nil {
		return nil, errors.WithStack(ErrUserNotFound)
	}

	var newPwd string
	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "gen password err %v", err)
		}

		newPwd = string(genPassword)
	}

	_, err = l.svcCtx.Ent.SysUser.UpdateOneID(in.Id).
		SetUsername(in.Username).
		SetNickname(in.Nickname).
		SetPassword(newPwd).
		SetAvatar(in.Avatar).
		SetSort(int8(in.Sort)).
		SetStatus(int8(in.Status)).
		SetCreator(oldEntity.Creator).
		SetUpdater(in.Op).
		Save(l.ctx)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update user err %v", err)
	}
	return &admin.UpdateUserResp{}, nil
}
