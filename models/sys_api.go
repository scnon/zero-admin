package models

import "gorm.io/gorm"

type SysAPI struct {
	gorm.Model
	Path   string `gorm:"type:varchar(255)" json:"path"`
	Method string `gorm:"type:varchar(10)" json:"method"`
	ResModel
}
