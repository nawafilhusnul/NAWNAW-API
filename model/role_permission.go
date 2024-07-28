package model

type RolePermission struct {
	ID           int        `gorm:"column:id;primary_key"`
	RoleID       int        `gorm:"column:role_id;not null"`
	PermissionID int        `gorm:"column:permission_id;not null"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}

func (rp RolePermission) TableName() string {
	return "role_permissions"
}
