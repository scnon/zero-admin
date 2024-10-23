package models

import (
	"errors"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);uniqueIndex:idx_name_tenant;comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(255);comment:密码"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Nickname string `json:"nickname" gorm:"type:varchar(255);comment:昵称"`
	Status   int8   `json:"status" gorm:"default:1;comment:状态 0:禁用 1:启用"`
	ResModel
}

// 分表
//func (m *SysUser) TableName() string {
//	return fmt.Sprintf("sys_user_%d", m.TenantID)
//}

func (m *SysUser) BeforeUpdate(tx *gorm.DB) error {
	if m.ID == 1 {
		return errors.New("内置管理员不可修改")
	}
	return nil
}
func (m *SysUser) BeforeDelete(tx *gorm.DB) error {
	if m.ID == 1 {
		return errors.New("内置管理员不可删除")
	}
	return nil
}
