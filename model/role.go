package model

import datatypes "github.com/nawafilhusnul/NAWNAW-API/common/datatypes"

type Role struct {
	ID     datatypes.ID         `json:"id" gorm:"column:id;primary_key" example:"YTFiMmMzZDRlNWY2ZzdoOGk5ajBrMWwybTNuNG81cDYyOQ=="`
	Name   datatypes.NullString `json:"name" gorm:"column:name" example:"Admin"`
	Slug   datatypes.NullString `json:"slug" gorm:"column:slug" example:"admin"`
	Module datatypes.NullString `json:"module" gorm:"column:module" example:"auth"`
}
