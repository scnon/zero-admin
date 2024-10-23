package menu

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

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
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 创建菜单
	_, err = l.svcCtx.Menu.AddMenu(l.ctx, &auth.AddMenuReq{
		ParentId: req.ParentID,
		Path:     req.Path,
		Title:    req.Title,
		Name:     req.Name,
		Sort:     req.Sort,
		Op:       uid,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}
	return &types.MenuCreateResp{
		Base: l.svcCtx.Success(),
	}, nil

}
