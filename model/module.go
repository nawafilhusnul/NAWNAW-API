package model

import "github.com/nawafilhusnul/NAWNAW-API/common/datatypes"

type Module struct {
	ID          int                  `json:"id" gorm:"column:id;primaryKey"`
	Name        string               `json:"name" gorm:"column:name;unique"`
	Description datatypes.NullString `json:"description" gorm:"column:description"`
	DefaultModel
}
