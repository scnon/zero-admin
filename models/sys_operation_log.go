package models

import (
	"gorm.io/gorm"
	"time"
)

type SysOperationLog struct {
	gorm.Model
	UserId        uint64    `json:"user_id"`
	User          *SysUser  `gorm:"foreignKey:UserId"`
	IP            string    `json:"ip" gorm:"type:varchar(64)"`
	Address       string    `json:"address" gorm:"type:varchar(255)"`
	System        string    `json:"system" gorm:"type:varchar(255)"`
	Browser       string    `json:"browser" gorm:"type:varchar(255)"`
	Status        int8      `json:"status"`
	Summary       string    `json:"summary" gorm:"type:varchar(255)"`
	Module        string    `json:"module" gorm:"type:varchar(255)"`
	OperationTime time.Time `json:"operation_time"`
	TenantID      uint64    `json:"tenant_id"`
}
