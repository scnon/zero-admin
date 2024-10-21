package models

import "gorm.io/gorm"

type SysAPI struct {
	gorm.Model
	Path   string `gorm:"length:255" json:"path"`
	Method string `gorm:"length:10" json:"method"`
	ResModel
}
