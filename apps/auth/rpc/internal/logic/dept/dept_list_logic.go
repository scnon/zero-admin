package deptlogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/models"
	"xlife/pkg/xerr"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *auth.DeptListReq) (*auth.DeptListResp, error) {
	// 1. 根据条件查询部门列表
	var depts []models.SysDept
	res := l.makeQuery(in).Find(&depts)
	if res.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询部门列表失败 %v", res.Error)
	}

	// 2. 构造返回数据
	var list []*auth.DeptData
	for _, dept := range depts {
		data := &auth.DeptData{
			Id:       uint64(dept.ID),
			Name:     dept.Name,
			ParentId: uint64(dept.ParentId),
			Sort:     int32(dept.Sort),
			Status:   int32(dept.Status),
		}
		if dept.Creator != nil {
			data.Creator = dept.Creator.Username
		}
		if dept.Updater != nil {
			data.Updater = dept.Updater.Username
		}
		list = append(list, data)
	}

	return &auth.DeptListResp{
		List: list,
	}, nil
}

func (l *DeptListLogic) makeQuery(in *auth.DeptListReq) *gorm.DB {
	query := l.svcCtx.DB.Model(&models.SysDept{}).Preload("Creator").Preload("Updater")
	if in.ParentId != nil {
		query = query.Where("id = ?", in.ParentId)
	}
	if in.Status != nil {
		query = query.Where("status = ?", in.Status)
	}
	if in.TenantId != nil {
		query = query.Where("tenant_id = ?", in.TenantId)
	}
	return query
}
