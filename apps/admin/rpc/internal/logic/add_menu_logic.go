package logic

import (
	"context"
	"database/sql"

	"zero-admin/apps/admin/rpc/admin"
	"zero-admin/apps/admin/rpc/internal/svc"
	"zero-admin/apps/model"
	"zero-admin/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMenuLogic) AddMenu(in *admin.AddMenuReq) (*admin.AddMenuResp, error) {
	menuEntity := model.SysMenu{
		ParentId: sql.NullInt64{Int64: in.ParentId, Valid: true},
		Title:    in.Title,
		Sort:     int64(in.Sort),
		Path:     in.Path,
		Name:     in.Name,
		TenantId: in.TenantId,
	}

	if _, err := l.svcCtx.MenuModel.Insert(l.ctx, &menuEntity); err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert menu error: %v", err)
	}

	return &admin.AddMenuResp{
		Id: menuEntity.Id,
	}, nil
}
