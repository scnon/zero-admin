package rolelogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/models"
	"xlife/pkg/xerr"

	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *auth.RoleListReq) (*auth.RoleListResp, error) {
	// 1. 查询角色列表
	var roles []models.SysRole
	res := l.makeQuery(in).Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&roles)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户失败 %v", res.Error)
	}

	// 2. 构造返回数据
	var list []*auth.RoleData
	for _, role := range roles {
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

	var total int64
	res = l.makeQuery(in).Count(&total)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户总数失败 %v", res.Error)
	}

	return &auth.RoleListResp{
		List:  list,
		Total: uint64(total),
	}, nil
}

func (l *RoleListLogic) makeQuery(in *auth.RoleListReq) *gorm.DB {
	query := l.svcCtx.DB.Model(&models.SysUser{}).Preload("Creator").Preload("Updater")
	if in.Ids != nil {
		query = query.Where("id IN (?)", in.Ids)
	}
	if in.Status != nil {
		query = query.Where("status = ?", in.Status)
	}
	if in.TenantId != nil {
		query = query.Where("tenant_id = ?", in.TenantId)
	}
	return query
}
