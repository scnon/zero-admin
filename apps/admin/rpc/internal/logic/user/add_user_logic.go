package userlogic

import (
	"context"
	"github.com/jinzhu/copier"
	perr "github.com/pkg/errors"
	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"
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
	entity := &model.SysUser{}
	err := copier.Copy(entity, in)
	if err != nil {
		return nil, perr.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
	}
	if _, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username); err == nil {
		return nil, perr.WithStack(ErrAlreadyExist)
	}

	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, perr.Wrapf(xerr.NewInternalErr(), "gen password err %v", err)
		}

		entity.Password = string(genPassword)
	}

	if _, err := l.svcCtx.UserModel.Insert(l.ctx, entity); err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "insert user err %v", err)
	}

	return &admin.AddUserResp{
		Id: entity.Id,
	}, nil
}
