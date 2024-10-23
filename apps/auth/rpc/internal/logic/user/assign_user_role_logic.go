package userlogic

import (
	"context"
	"errors"
	perr "github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"xlife/models"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignUserRoleLogic {
	return &AssignUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssignUserRoleLogic) AssignUserRole(in *auth.AssignUserRoleReq) (*auth.AssignUserRoleResp, error) {
	// 1. 查询用户
	var user models.SysUser
	res := l.svcCtx.DB.Where("id = ?", in.UserId).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrUserNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询用户失败: %v", res.Error)
	}
	// 2. 查询角色
	var roles []models.SysRole
	res = l.svcCtx.DB.Where("id IN ?", in.RoleIds).Find(&roles)
	if res.Error != nil {
		return nil, perr.Wrapf(res.Error, "查询角色失败: %v", res.Error)
	}
	// 3. 指定角色
	tenantId := strconv.FormatUint(in.TenantId, 10)
	for _, role := range roles {
		ok, err := l.svcCtx.Casbin.AddRoleForUserInDomain(user.Username, role.Name, tenantId)
		if err != nil || !ok {
			return nil, perr.Wrapf(res.Error, "指定角色失败: %v", err)
		}
	}

	return &auth.AssignUserRoleResp{}, nil
}
