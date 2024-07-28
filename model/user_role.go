package model

type UserRole struct {
	ID     int  `gorm:"column:id;primary_key"`
	UserID int  `gorm:"column:user_id;not null"`
	RoleID int  `gorm:"column:role_id;not null"`
	User   User `gorm:"foreignKey:UserID"`
	Role   Role `gorm:"foreignKey:RoleID"`
}

func (ur UserRole) TableName() string {
	return "user_roles"
}
