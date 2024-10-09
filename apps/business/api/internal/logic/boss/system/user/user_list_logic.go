package user

import (
	"context"
	"time"
	"zero-admin/apps/admin/rpc/admin"

	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

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
	result, err := l.svcCtx.User.UserList(l.ctx, &admin.UserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   req.Status,
		Username: req.Username,
		Nickname: req.NickName,
		TenantId: l.svcCtx.Config.Tenant,
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
			Status:     item.Status,
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
