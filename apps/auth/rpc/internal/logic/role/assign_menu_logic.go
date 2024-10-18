package rolelogic

import (
	"context"
	"strconv"
	"xlife/models"

	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignMenuLogic {
	return &AssignMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssignMenuLogic) AssignMenu(in *auth.AssignMenuReq) (*auth.AssignMenuResp, error) {
	var role models.SysRole
	res := l.svcCtx.DB.Where("id = ?", in.RoleId).First(&role)
	if res.Error != nil {
		return nil, res.Error
	}
	var menus []models.SysMenu
	res = l.svcCtx.DB.Where("id IN (?)", in.MenuIds).Find(&menus)

	for _, menu := range menus {
		_, err := l.svcCtx.Casbin.AddPolicy(strconv.FormatUint(uint64(role.ID), 10),
			strconv.FormatUint(in.TenantId, 10),
			strconv.FormatUint(uint64(menu.ID), 10), "read")
		if err != nil {
			return nil, err
		}
	}
	err := l.svcCtx.Casbin.SavePolicy()
	if err != nil {
		return nil, err
	}

	return &auth.AssignMenuResp{}, nil
}
