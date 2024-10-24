package user

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// 1. 获取用户列表
	res, err := l.svcCtx.User.UserList(l.ctx, &auth.UserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   req.Status,
		Username: req.Username,
		Nickname: req.NickName,
		TenantId: &l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	// 2. 构造返回数据
	list := make([]types.UserData, 0)
	for _, item := range res.List {
		list = append(list, types.UserData{
			Id:         item.Id,
			NickName:   item.Nickname,
			Username:   item.Username,
			Status:     item.Status,
			Sort:       item.Sort,
			Remark:     item.Remark,
			Avatar:     item.Avatar,
			Creator:    item.Creator,
			Updator:    item.Updater,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		})
	}
	return &types.UserListResp{
		Base: l.svcCtx.Success(),
		Data: types.UserListData{
			Total: res.Total,
			List:  list,
		},
	}, nil
}
