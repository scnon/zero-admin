package userlogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
	"xlife/pkg/xerr"

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

func (l *UserListLogic) UserList(in *auth.UserListReq) (*auth.UserListResp, error) {
	// 1. 查询用户列表
	var users []models.SysUser
	res := l.makeQuery(in).Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&users)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户失败 %v", res.Error)
	}
	// 2. 构造返回数据
	var list []*auth.UserData
	for _, user := range users {
		data := &auth.UserData{
			Id:         uint64(user.ID),
			Username:   user.Username,
			Nickname:   user.Nickname,
			Avatar:     user.Avatar,
			Sort:       user.Sort,
			Remark:     user.Remark,
			TenantId:   uint64(user.TenantID),
			CreateTime: uint64(user.CreatedAt.Unix()),
			UpdateTime: uint64(user.UpdatedAt.Unix()),
		}
		if user.Creator != nil {
			data.Creator = user.Creator.Username
		}
		if user.Updater != nil {
			data.Updater = user.Updater.Username
		}
		list = append(list, data)
	}

	// 3. 查询用户总数
	var total int64
	res = l.makeQuery(in).Count(&total)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户总数失败 %v", res.Error)
	}

	return &auth.UserListResp{
		List:  list,
		Total: uint64(total),
	}, nil
}

func (l *UserListLogic) makeQuery(in *auth.UserListReq) *gorm.DB {
	query := l.svcCtx.DB.Model(&models.SysUser{}).Preload("Creator").Preload("Updater")
	if in.Ids != nil {
		query = query.Where("id IN (?)", in.Ids)
	}
	if in.Username != nil {
		query = query.Where("username = ?", in.Username)
	}
	if in.Status != nil {
		query = query.Where("status = ?", in.Status)
	}
	if in.TenantId != nil {
		query = query.Where("tenant_id = ?", in.TenantId)
	}
	return query
}
