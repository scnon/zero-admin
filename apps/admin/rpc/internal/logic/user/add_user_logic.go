package userlogic

import (
	"context"
	perr "github.com/pkg/errors"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/ent/sysuser"
	"zero-admin/pkg/encrypt"
	"zero-admin/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrAlreadyExist = xerr.NewMsg("用户名已存在")
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *admin.AddUserReq) (*admin.AddUserResp, error) {
	exist, err := l.svcCtx.Ent.SysUser.Query().Where(sysuser.UsernameEQ(in.Username)).Exist(l.ctx)
	if err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}
	if exist {
		return nil, perr.WithStack(ErrAlreadyExist)
	}

	var secretPwd string
	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, perr.Wrapf(xerr.NewInternalErr(), "gen password err %v", err)
		}

		secretPwd = string(genPassword)
	}

	entity, err := l.svcCtx.Ent.SysUser.Create().
		SetUsername(in.Username).
		SetNickname(in.Nickname).
		SetPassword(secretPwd).
		SetAvatar(in.Avatar).
		SetSort(int8(in.Sort)).
		SetStatus(int8(in.Sort)).
		SetCreator(in.Op).
		Save(l.ctx)

	if err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "create user err %v", err)
	}

	return &admin.AddUserResp{
		Id: entity.ID,
	}, nil
}
