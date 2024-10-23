package userlogic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xlife/apps/auth/rpc/auth"
	"xlife/apps/auth/rpc/internal/svc"
	"xlife/models"
	"xlife/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrAlreadyExist = xerr.NewMsg("用户名已存在")
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *auth.AddUserReq) (*auth.AddUserResp, error) {
	// 1. 检查用户是否已存在
	var existingUser models.SysUser
	res := l.svcCtx.DB.Where("username = ?", in.Username).First(&existingUser)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "查询用户失败: %v", res.Error)
	}
	if existingUser.ID != 0 {
		return nil, errors.WithStack(ErrAlreadyExist)
	}

	// 2. 创建新用户
	newUser := models.SysUser{
		Username: in.Username,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		ResModel: models.ResModel{
			Sort:      int(in.Sort),
			Remark:    in.Remark,
			TenantID:  uint(in.TenantId),
			CreatorID: uint(in.Op),
		},
	}
	if err := l.svcCtx.DB.Create(&newUser).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "创建用户失败: %v", err)
	}

	return &auth.AddUserResp{
		Id: uint64(newUser.ID),
	}, nil
}
