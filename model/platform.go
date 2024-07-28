package model

import datatypes "github.com/nawafilhusnul/NAWNAW-API/common/datatypes"

type Platform struct {
	ID   int                  `gorm:"column:id;primary_key" example:"YTFiMmMzZDRlNWY2ZzdoOGk5ajBrMWwybTNuNG81cDYyOQ=="`
	Name datatypes.NullString `gorm:"column:name" example:"Basic"`
	Slug datatypes.NullString `gorm:"column:slug" example:"basic"`
	DefaultModel
}

func (p Platform) TableName() string {
	return "platforms"
}
