package userlogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"zero-admin/pkg/xerr"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *admin.UserListReq) (*admin.UserListResp, error) {
	entityList, total, err := l.svcCtx.UserModel.FindAll(l.ctx, in.Ids, in.Nickname, in.Username, in.Status, in.TenantId, in.Page, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}
	userList := make([]*admin.UserData, 0)
	for _, entity := range entityList {
		var user admin.UserData
		if err := copier.Copy(&user, &entity); err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		user.CreateTime = entity.CreateTime.Unix()
		user.UpdateTime = entity.UpdateTime.Unix()
		user.Creator = entity.CreatorName.String
		if entity.Creator == 0 {
			user.Creator = "系统"
		}
		user.Updater = entity.UpdaterName.String
		userList = append(userList, &user)
	}

	return &admin.UserListResp{
		List:  userList,
		Total: total,
	}, nil
}
