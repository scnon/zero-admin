package models

type ResModel struct {
	Sort     int    `json:"sort"`
	Remark   string `gorm:"size:255" json:"remark"`
	TenantID uint   `json:"tenant_id"`

	CreatorID uint     `json:"creator_id"`           // 创建者 ID，不可变
	Creator   *SysUser `gorm:"foreignKey:CreatorID"` // 关联创建者的详细信息

	UpdaterID *uint    `json:"updater_id"`           // 更新者 ID，可为空
	Updater   *SysUser `gorm:"foreignKey:UpdaterID"` // 关联更新者的详细信息

	DeleterID *uint    `json:"deleter_id"`           // 删除者 ID，可为空
	Deleter   *SysUser `gorm:"foreignKey:DeleterID"` // 关联删除者的详细信息
}
