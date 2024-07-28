package model

type UserPlatform struct {
	ID         int `gorm:"column:id;primary_key" example:"YTFiMmMzZDRlNWY2ZzdoOGk5ajBrMWwybTNuNG81cDYyOQ=="`
	UserID     int `gorm:"column:user_id" example:"1"`
	PlatformID int `gorm:"column:platform_id" example:"1"`
	DefaultModel
	User     User     `gorm:"foreignKey:UserID"`
	Platform Platform `gorm:"foreignKey:PlatformID"`
}

func (u UserPlatform) TableName() string {
	return "user_platforms"
}
