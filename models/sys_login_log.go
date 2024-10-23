package models

import (
	"gorm.io/gorm"
	"time"
)

type SysLoginLog struct {
	gorm.Model
	UserID    uint64    `json:"user_id" gorm:"comment: 操作用户ID"`
	User      *SysUser  `gorm:"foreignKey:UserID;comment:操作用户"`
	IP        string    `json:"ip" gorm:"type:varchar(64);comment:IP地址"`
	Address   string    `json:"address" gorm:"type:varchar(255);comment:地址"`
	System    string    `json:"system" gorm:"type:varchar(255);comment:操作系统"`
	Browser   string    `json:"browser" gorm:"type:varchar(255);comment:浏览器"`
	Status    int8      `json:"status" gorm:"comment:状态 0:失败 1:成功"`
	Behavior  string    `json:"behavior" gorm:"type:varchar(255);comment:行为"`
	LoginTime time.Time `json:"login_time" gorm:"comment:登录时间"`
	TenantID  uint64    `json:"tenant_id" gorm:"comment:租户 ID"`
}
