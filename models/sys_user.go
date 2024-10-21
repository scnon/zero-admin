package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Username string `gorm:"uniqueIndex length:255" json:"username"`
	Password string `gorm:"length:255" json:"password"`
	Avatar   string `gorm:"length:255" json:"avatar"`
	Nickname string `gorm:"length:255" json:"nickname"`
	Status   int8   `json:"status"`
	ResModel
}
