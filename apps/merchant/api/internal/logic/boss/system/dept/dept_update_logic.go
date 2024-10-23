package dept

import (
	"context"
	"xlife/apps/auth/rpc/auth"
	"xlife/pkg/ctxdata"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptUpdateLogic) DeptUpdate(req *types.DeptUpdateReq) (resp *types.DeptUpdateResp, err error) {
	// 1. 获取当前用户
	uid := ctxdata.GetUId(l.ctx)
	// 2. 更新部门
	if _, err := l.svcCtx.Dept.UpdateDept(l.ctx, &auth.UpdateDeptReq{
		Id:       req.ID,
		ParentId: req.ParentID,
		Name:     req.Name,
		Sort:     req.Sort,
		Status:   req.Status,
		Op:       uid,
	}); err != nil {
		return nil, err
	}
	return &types.DeptUpdateResp{
		Base: l.svcCtx.Success(),
	}, nil
}
