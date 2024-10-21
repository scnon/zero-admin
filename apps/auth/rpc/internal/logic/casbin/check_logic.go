package casbinlogic

import (
	"context"

	"xlife/apps/auth/rpc/admin"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *admin.CasbinCheckReq) (*admin.CasbinCheckResp, error) {
	// todo: add your logic here and delete this line

	return &admin.CasbinCheckResp{}, nil
}
