package userlogic

import (
	"context"
	"strconv"
	"xlife/models"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolesLogic {
	return &GetRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRolesLogic) GetRoles(in *auth.GetRolesReq) (*auth.GetRolesResp, error) {
	var user models.SysUser
	res := l.svcCtx.DB.Where("id = ?", in.UserId).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	roles := l.svcCtx.Casbin.GetRolesForUserInDomain(strconv.FormatUint(uint64(user.ID), 10),
		strconv.FormatUint(in.TenantId, 10))
	var roleIds []uint64
	for _, role := range roles {
		id, err := strconv.ParseUint(role, 10, 64)
		if err != nil {
			return nil, err
		}
		roleIds = append(roleIds, id)
	}
	var roleList []models.SysRole
	res = l.svcCtx.DB.Where("id IN (?)", roleIds).Find(&roleList)
	if res.Error != nil {
		return nil, res.Error
	}
	var list []*auth.RoleData
	for _, role := range roleList {
		data := &auth.RoleData{
			Id:     uint64(role.ID),
			Name:   role.Name,
			Sort:   int32(role.Sort),
			Remark: role.Remark,
		}
		if role.Creator != nil {
			data.Creator = role.Creator.Username
		}
		if role.Updater != nil {
			data.Updater = role.Updater.Username
		}
		list = append(list, data)
	}

	return &auth.GetRolesResp{
		Roles: list,
	}, nil
}
