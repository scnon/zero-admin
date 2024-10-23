package models

import "gorm.io/gorm"

type SysDept struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(128);comment:部门名称"`
	ParentID uint   `json:"parent_id" gorm:"comment:上级部门ID"`
	Status   int8   `json:"status" gorm:"default:1;comment:部门状态"`
	ResModel
}
