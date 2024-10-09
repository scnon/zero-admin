package userlogic

import (
	"context"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *admin.GetUserReq) (*admin.GetUserResp, error) {
	entities, total, err := l.svcCtx.UserModel.FindAll(l.ctx, in.Ids, in.Nickname, in.Username, in.Status, in.TenantId,
		in.Page, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}

	userList := make([]*admin.UserData, 0)
	for _, entity := range entities {
		var user admin.UserData
		err := copier.Copy(&user, &entity)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}

		user.CreateTime = entity.CreateTime.Unix()
		user.UpdateTime = entity.UpdateTime.Unix()
		userList = append(userList, &user)
	}

	return &admin.GetUserResp{
		List:  userList,
		Total: total,
	}, nil
}
