package models

import (
	"gorm.io/gorm"
)

type SysRole struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(255);uniqueIndex:idx_name_tenant;comment:角色名称"`
	Status int8   `json:"status" gorm:"default:1;comment:状态 0:禁用 1:启用"`
	ResModel
}
