package dept

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"

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
	res, err := l.svcCtx.DeleteDept(l.ctx, &auth.DeleteDeptReq{
		Ids:      req.Ids,
		TenantId: l.svcCtx.Config.Tenant,
	})

}
