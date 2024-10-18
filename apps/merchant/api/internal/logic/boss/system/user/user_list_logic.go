package user

import (
	"context"
	"time"
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
	var username *string
	var nickname *string
	if req.Username != "" {
		username = new(string)
		*username = req.Username
	}
	if req.NickName != "" {
		nickname = new(string)
		*nickname = req.NickName
	}
	result, err := l.svcCtx.User.UserList(l.ctx, &auth.UserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Username: username,
		Nickname: nickname,
		TenantId: &l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.UserData, 0)
	for _, item := range result.List {
		list = append(list, types.UserData{
			Id:         item.Id,
			NickName:   item.Nickname,
			Username:   item.Username,
			Status:     int64(item.Status),
			Avatar:     item.Avatar,
			Creator:    item.Creator,
			Updator:    item.Updater,
			CreateTime: time.Unix(item.CreateTime, 0).Format("2006-01-02 15:04:05"),
			UpdateTime: time.Unix(item.UpdateTime, 0).Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.UserListResp{
		Base: l.svcCtx.Success(),
		Data: types.UserListData{
			Total: result.Total,
			List:  list,
		},
	}
	return

}
