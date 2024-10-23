package casbinlogic

import (
	"context"
	"xlife/apps/auth/rpc/auth"

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

func (l *CheckLogic) Check(in *auth.CasbinCheckReq) (*auth.CasbinCheckResp, error) {
	// todo: add your logic here and delete this line

	return &auth.CasbinCheckResp{}, nil
}
