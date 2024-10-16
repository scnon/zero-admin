package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name string `gorm:"length:255" json:"name"`
	ResModel
}
