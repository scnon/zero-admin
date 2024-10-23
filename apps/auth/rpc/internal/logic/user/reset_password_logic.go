package userlogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/models"
	"xlife/pkg/encrypt"
	"xlife/pkg/xerr"

	perr "github.com/pkg/errors"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPasswordLen = xerr.NewMsg("密码长度不能小于8")
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *auth.ResetPasswordReq) (*auth.ResetPasswordResp, error) {
	// 1. 查询用户是否存在
	var entity models.SysUser
	res := l.svcCtx.DB.Where("id = ?", in.UserId).First(&entity)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, perr.WithStack(ErrUserNotFound)
		}
		return nil, perr.Wrapf(xerr.NewDBErr(), "查询用户失败: %v", res.Error)
	}
	// 2. 验证密码规则
	if len(in.Password) < 8 {
		return nil, perr.WithStack(ErrPasswordLen)
	}
	// 3. 密码加密
	newPwd, err := encrypt.GenPasswordHash([]byte(in.Password))
	if err != nil {
		return nil, perr.Wrapf(xerr.NewInternalErr(), "生成密码失败: %v", err)
	}
	// 4. 更新密码
	res = l.svcCtx.DB.Model(&entity).Update("password", newPwd)
	if res.Error != nil {
		return nil, perr.Wrapf(xerr.NewDBErr(), "更新密码失败: %v", res.Error)
	}
	return &auth.ResetPasswordResp{}, nil
}
