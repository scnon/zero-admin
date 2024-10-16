package models

import "gorm.io/gorm"

type SysMenu struct {
	gorm.Model
	Name     string `gorm:"length:255" json:"name"`
	Title    string `gorm:"length:255" json:"title"`
	Path     string `gorm:"length:255" json:"path"`
	ParentID uint   `json:"parent_id"`
	ResModel
}
