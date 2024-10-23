package models

type ResModel struct {
	Sort     int32  `json:"sort" gorm:"comment:排序"`
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	TenantID uint   `json:"tenant_id" gorm:"uniqueIndex:idx_name_tenant;comment:租户 ID"`

	CreatorID uint     `json:"creator_id" gorm:"comment:创建者ID"` // 创建者 ID，不可变
	Creator   *SysUser `gorm:"foreignKey:CreatorID"`

	UpdaterID *uint    `json:"updater_id" gorm:"comment:更新者ID"` // 更新者 ID，可为空
	Updater   *SysUser `gorm:"foreignKey:UpdaterID"`

	DeleterID *uint    `json:"deleter_id" gorm:"comment:删除者ID"` // 删除者 ID，可为空
	Deleter   *SysUser `gorm:"foreignKey:DeleterID"`
}
