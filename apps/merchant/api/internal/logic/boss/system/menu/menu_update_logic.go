package menu

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.MenuUpdateReq) (resp *types.MenuUpdateResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 更新菜单
	_, err = l.svcCtx.Menu.UpdateMenu(l.ctx, &auth.UpdateMenuReq{
		Id:       req.ID,
		ParentId: req.ParentID,
		Path:     req.Path,
		Title:    req.Title,
		Name:     req.Name,
		Op:       uid,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	return &types.MenuUpdateResp{
		Base: l.svcCtx.Success(),
	}, nil
}
