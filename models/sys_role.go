package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name   string `gorm:"length:255" json:"name"`
	Status int8   `json:"status"`
	ResModel
}
