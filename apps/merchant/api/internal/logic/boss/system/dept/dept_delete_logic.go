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
	uid := ctxdata.GetUId(l.ctx)
	_, err = l.svcCtx.DeleteDept(l.ctx, &auth.DeleteDeptReq{
		Ids:      req.Ids,
		TenantId: l.svcCtx.Config.Tenant,
		Op:       uid,
	})

	if err != nil {
		return
	}

	return &types.DeptDeleteResp{
		Base: l.svcCtx.Success(),
	}, nil
}
