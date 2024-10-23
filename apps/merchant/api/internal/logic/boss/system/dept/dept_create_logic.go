package dept

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptCreateLogic {
	return &DeptCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptCreateLogic) DeptCreate(req *types.DeptCreateReq) (resp *types.DeptCreateResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 创建部门
	res, err := l.svcCtx.Dept.AddDept(l.ctx, &auth.AddDeptReq{
		Name:     req.Name,
		ParentId: req.ParentID,
		Sort:     req.Sort,
		Op:       uid,
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return
	}
	return &types.DeptCreateResp{
		Base: l.svcCtx.Success(),
		Data: res.Id,
	}, nil
}
