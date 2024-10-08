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

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *admin.UpdateMenuReq) (*admin.UpdateMenuResp, error) {
	err := l.svcCtx.MenuModel.Update(l.ctx, &model.SysMenu{
		Id:       in.Id,
		ParentId: sql.NullInt64{Int64: in.ParentId, Valid: true},
		Sort:     int64(in.Sort),
		Path:     in.Path,
		Title:    in.Title,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "update menu error: %v", err)
	}
	return &admin.UpdateMenuResp{}, nil
}
