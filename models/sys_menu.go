package models

import (
	"gorm.io/gorm"
)

type SysMenu struct {
	gorm.Model
	Name            string `json:"name" gorm:"type:varchar(255);uniqueIndex:idx_name_tenant;comment:菜单名称"`
	Title           string `json:"title" gorm:"type:varchar(255);comment:标题"`
	Path            string `json:"path" gorm:"type:varchar(255);comment:路径"`
	Component       string `json:"component" gorm:"type:varchar(255);comment:组件"`
	ParentID        uint   `json:"parent_id" gorm:"comment:父级菜单ID"`
	Icon            string `json:"icon" gorm:"type:varchar(255);comment:图标"`
	Type            int8   `json:"type" gorm:"comment:类型 0:目录 1:菜单 2:按钮"`
	ExtraIcon       string `json:"extra_icon" gorm:"type:varchar(255);comment:扩展图标"`
	FrameSrc        string `json:"frame_src" gorm:"type:varchar(255);comment:内嵌网页"`
	FrameLoading    bool   `json:"frame_loading" gorm:"comment:内嵌网页加载方式"`
	ShowLink        bool   `json:"show_link" gorm:"comment:是否显示链接"`
	ShowParent      bool   `json:"show_parent" gorm:"comment:是否显示父级"`
	EnterTransition string `json:"enter_transition" gorm:"type:varchar(255);comment:进入动画"`
	LeaveTransition string `json:"leave_transition" gorm:"type:varchar(255);comment:离开动画"`
	Redirection     string `json:"redirection" gorm:"type:varchar(255);comment:重定向地址"`
	Status          int8   `json:"status" gorm:"default:1;comment:状态 0:禁用 1:启用"`
	ResModel
}
