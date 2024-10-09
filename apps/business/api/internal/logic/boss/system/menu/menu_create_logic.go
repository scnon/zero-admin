package menu

import (
	"context"
	"zero-admin/apps/admin/rpc/admin"

	"zero-admin/apps/business/api/internal/svc"
	"zero-admin/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuCreateLogic {
	return &MenuCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuCreateLogic) MenuCreate(req *types.MenuCreateReq) (resp *types.MenuCreateResp, err error) {
	_, err = l.svcCtx.Menu.AddMenu(l.ctx, &admin.AddMenuReq{
		ParentId: req.ParentID,
		Path:     req.Path,
		Title:    req.Title,
		Name:     req.Name,
		Sort:     req.Sort,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	return &types.MenuCreateResp{
		Base: l.svcCtx.Success(),
	}, nil

}
