package model

import "github.com/nawafilhusnul/NAWNAW-API/common/datatypes"

type Permission struct {
	ID     datatypes.ID         `json:"id" gorm:"column:id;primary_key" example:"YTFiMmMzZDRlNWY2ZzdoOGk5ajBrMWwybTNuNG81cDYyOQ=="`
	Name   datatypes.NullString `json:"name" gorm:"column:name" example:"Register"`
	Slug   datatypes.NullString `json:"slug" gorm:"column:slug" example:"register"`
	Module datatypes.NullString `json:"module" gorm:"column:module" example:"auth"`
	DefaultModel
}

func (p Permission) TableName() string {
	return "permissions"
}
