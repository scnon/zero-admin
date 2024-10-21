package models

import (
	"gorm.io/gorm"
	"time"
)

type SysLoginLog struct {
	gorm.Model
	UserId    uint64    `json:"user_id"`
	User      *SysUser  `gorm:"foreignKey:UserId"`
	IP        string    `json:"ip" gorm:"type:varchar(64)"`
	Address   string    `json:"address" gorm:"type:varchar(255)"`
	System    string    `json:"system" gorm:"type:varchar(255)"`
	Browser   string    `json:"browser" gorm:"type:varchar(255)"`
	Status    int8      `json:"status"`
	Behavior  string    `json:"behavior" gorm:"type:varchar(255)"`
	LoginTime time.Time `json:"login_time"`
	TenantID  uint64    `json:"tenant_id"`
}
