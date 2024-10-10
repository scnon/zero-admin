package user

import (
	"context"
	"zero-admin/apps/admin/rpc/admin"

	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	_, err = l.svcCtx.User.AddUser(l.ctx, &admin.AddUserReq{
		Username: req.Username,
		Nickname: req.NickName,
		Password: req.Password,
		Sort:     req.Sort,
		Remark:   req.Remark,
		Status:   req.Status,
		Roles:    req.Roles,
		Avatar:   req.Avatar,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserCreateResp{
		Base: l.svcCtx.Success(),
	}, nil

}
