package logic

import (
	"context"
	"database/sql"
	"errors"
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
	userEntity := model.SysUser{
		Username: in.Username,
		Nickname: in.Nickname,
		Status:   int64(in.Status),
		Sort:     int64(in.Sort),
		TenantId: sql.NullInt64{Int64: in.TenantId, Valid: true},
	}
	if _, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username); err == nil {
		if !errors.Is(err, model.ErrNotFound) {
			return nil, perr.WithStack(ErrAlreadyExist)
		}
	}

	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, perr.Wrapf(xerr.NewInternalErr(), "gen password err %v", err)
		}

		userEntity.Password = string(genPassword)
	}

	if _, err := l.svcCtx.UserModel.Insert(l.ctx, &userEntity); err != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "insert user err %v", err)
	}

	return &admin.AddUserResp{}, nil
}
