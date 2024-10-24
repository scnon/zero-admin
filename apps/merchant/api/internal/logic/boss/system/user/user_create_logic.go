package user

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"
	"xlife/pkg/ctxdata"

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
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 创建用户
	if _, err = l.svcCtx.User.AddUser(l.ctx, &auth.AddUserReq{
		Username: req.Username,
		Nickname: req.NickName,
		Sort:     req.Sort,
		Remark:   req.Remark,
		Status:   req.Status,
		Avatar:   req.Avatar,
		TenantId: l.svcCtx.Config.Tenant,
		Op:       uid,
	}); err != nil {
		return nil, err
	}
	return &types.UserCreateResp{
		Base: l.svcCtx.Success(),
	}, nil
}
