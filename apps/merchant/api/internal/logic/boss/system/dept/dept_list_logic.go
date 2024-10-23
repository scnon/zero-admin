package dept

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/merchant/api/internal/svc"
	"xlife/apps/merchant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(req *types.DeptListReq) (resp *types.DeptListResp, err error) {
	res, err := l.svcCtx.Dept.DeptList(l.ctx, &auth.DeptListReq{})
	if err != nil {
		return nil, err
	}
	var list []*types.DeptData
	for _, item := range res.List {
		list = append(list, &types.DeptData{
			ID:       item.Id,
			Name:     item.Name,
			ParentID: item.ParentId,
			Sort:     item.Sort,
			Status:   item.Status,
		})
	}
	return &types.DeptListResp{
		Base: l.svcCtx.Success(),
		Data: list,
	}, nil
}
