package model

import "github.com/nawafilhusnul/NAWNAW-API/common/datatypes"

type Permission struct {
	ID       datatypes.ID         `json:"id" gorm:"column:id;primary_key" example:"YTFiMmMzZDRlNWY2ZzdoOGk5ajBrMWwybTNuNG81cDYyOQ=="`
	Name     datatypes.NullString `json:"name" gorm:"column:name" example:"Register"`
	Slug     datatypes.NullString `json:"slug" gorm:"column:slug" example:"register"`
	ModuleID datatypes.ID         `json:"module_id" gorm:"column:module_id" example:"YTFiMmMzZDRlNWY2ZzdoOGk5ajBrMWwybTNuNG81cDYyOQ=="`
	Module   Module               `json:"module" gorm:"foreignKey:ModuleID"`
	DefaultModel
}

func (p Permission) TableName() string {
	return "permissions"
}

type CreatePermissionRequest struct {
	Name     string `json:"name" validate:"required"`
	ModuleID int    `json:"module_id" validate:"required"`
}
