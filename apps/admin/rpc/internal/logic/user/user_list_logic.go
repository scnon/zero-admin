package userlogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"zero-admin/ent/predicate"
	"zero-admin/ent/sysuser"
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
	var conditions []predicate.SysUser
	if in.Ids != nil {
		conditions = append(conditions, sysuser.IDIn(in.Ids...))
	}
	if in.Nickname != nil {
		conditions = append(conditions, sysuser.NicknameEQ(*in.Nickname))
	}
	if in.Username != nil {
		conditions = append(conditions, sysuser.UsernameEQ(*in.Username))
	}
	if in.Status != nil {
		conditions = append(conditions, sysuser.StatusEQ(int8(*in.Status)))
	}
	entityList, err := l.svcCtx.Ent.SysUser.Query().Where(conditions...).
		WithCreateBy().WithUpdateBy().All(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user err %v", err)
	}
	total, err := l.svcCtx.Ent.SysUser.Query().Where(conditions...).Count(l.ctx)

	userList := make([]*admin.UserData, 0)
	for _, entity := range entityList {
		var user admin.UserData
		if err := copier.Copy(&user, &entity); err != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "copy entity err %v", err)
		}
		user.CreateTime = entity.CreatedAt
		user.UpdateTime = entity.UpdatedAt
		if entity.Edges.CreateBy == nil {
			user.Creator = entity.Edges.CreateBy.Username
		}
		if entity.Creator == 0 {
			user.Creator = "系统"
		}
		if entity.Edges.UpdateBy == nil {
			user.Updater = entity.Edges.UpdateBy.Username
		}
		userList = append(userList, &user)
	}

	return &admin.UserListResp{
		List:  userList,
		Total: int64(total),
	}, nil
}
