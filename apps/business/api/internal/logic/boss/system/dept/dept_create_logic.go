package dept

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"

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
	uid := ctxdata.GetUId(l.ctx)
	res, err := l.svcCtx.Dept.AddDept(l.ctx, &auth.AddDeptReq{
		Name:     req.Name,
		ParentId: req.ParentID,
		Sort:     req.Sort,
		TenantId: l.svcCtx.Config.Tenant,
		Op:       uid,
	})
	if err != nil {
		return
	}
	return &types.DeptCreateResp{
		Base: l.svcCtx.Success(),
		Data: res.Id,
	}, nil
}
