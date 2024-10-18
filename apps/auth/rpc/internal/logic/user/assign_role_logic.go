package userlogic

import (
	"context"
	"strconv"
	"xlife/models"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleLogic {
	return &AssignRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssignRoleLogic) AssignRole(in *auth.AssignRoleReq) (*auth.AssignRoleResp, error) {
	var user models.SysUser
	res := l.svcCtx.DB.Where("id = ?", in.UserId).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	var roles []models.SysRole
	res = l.svcCtx.DB.Where("id IN (?)", in.RoleIds).Find(&roles)
	if res.Error != nil {
		return nil, res.Error
	}
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, strconv.FormatUint(uint64(role.ID), 10))
	}

	ok, err := l.svcCtx.Casbin.AddRolesForUser(strconv.FormatUint(uint64(user.ID), 10), roleNames,
		strconv.FormatUint(in.TenantId, 10))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, res.Error
	}
	return &auth.AssignRoleResp{}, nil
}
