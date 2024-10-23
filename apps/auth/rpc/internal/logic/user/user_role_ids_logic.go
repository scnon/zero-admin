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

type UserRoleIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleIdsLogic {
	return &UserRoleIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleIdsLogic) UserRoleIds(in *auth.UserRoleIdsReq) (*auth.UserRoleIdsResp, error) {
	// 1. 查询用户
	var entity models.SysUser
	res := l.svcCtx.DB.Where("id = ?", in.UserId).Where("tenant_id = ?", in.TenantId).First(&entity)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrUserNotFound)
		}
		return nil, perr.Wrapf(res.Error, "查询用户失败: %v", res.Error)
	}
	domainId := strconv.FormatUint(in.TenantId, 10)
	// 2. 查询用户的角色
	roleNames := l.svcCtx.Casbin.GetUsersForRoleInDomain(entity.Username, domainId)
	var roles []models.SysRole
	res = l.svcCtx.DB.Where("name IN ?", roleNames).Find(&roles)
	if res.Error != nil {
		return nil, perr.Wrapf(res.Error, "查询用户角色失败: %v", res.Error)
	}
	// 3. 返回角色id
	ids := make([]uint64, 0)
	for _, role := range roles {
		ids = append(ids, uint64(role.ID))
	}
	return &auth.UserRoleIdsResp{
		RoleIds: ids,
	}, nil
}
