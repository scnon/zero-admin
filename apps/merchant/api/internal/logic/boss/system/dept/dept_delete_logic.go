package dept

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptDeleteLogic) DeptDelete(req *types.DeptDeleteReq) (resp *types.DeptDeleteResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 删除部门
	if _, err = l.svcCtx.DeleteDept(l.ctx, &auth.DeleteDeptReq{
		Ids:      req.Ids,
		Op:       uid,
		TenantId: l.svcCtx.Config.Tenant,
	}); err != nil {
		return nil, err
	}

	return &types.DeptDeleteResp{
		Base: l.svcCtx.Success(),
	}, nil
}
