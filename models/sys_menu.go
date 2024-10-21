package models

import "gorm.io/gorm"

type SysMenu struct {
	gorm.Model
	Name            string `gorm:"length:255" json:"name"`
	Title           string `gorm:"length:255" json:"title"`
	Path            string `gorm:"length:255" json:"path"`
	Component       string `gorm:"length:255" json:"component"`
	ParentID        uint   `json:"parent_id"`
	Icon            string `gorm:"length:255" json:"icon"`
	Type            int8   `json:"type"`
	ExtraIcon       string `gorm:"length:255" json:"extra_icon"`
	FrameSrc        string `gorm:"length:255" json:"frame_src"`
	FrameLoading    bool   `json:"frame_loading"`
	ShowLink        bool   `json:"show_link"`
	ShowParent      bool   `json:"show_parent"`
	EnterTransition string `gorm:"length:255" json:"enter_transition"`
	LeaveTransition string `gorm:"length:255" json:"leave_transition"`
	Redirection     string `gorm:"length:255" json:"redirection"`
	Status          int8   `json:"status"`
	ResModel
}
