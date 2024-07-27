package model

import (
	"time"

	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
)

type DefaultModel struct {
	CreatedAt time.Time          `json:"-" gorm:"column:created_at" example:"2021-01-01 00:00:00"`
	CreatedBy int                `json:"-" gorm:"column:created_by" example:"0"`
	UpdatedAt time.Time          `json:"-" gorm:"column:updated_at" example:"2021-01-01 00:00:00"`
	UpdatedBy int                `json:"-" gorm:"column:updated_by" example:"0"`
	DeletedAt datatypes.NullTime `json:"-" gorm:"column:deleted_at" example:"2021-01-01 00:00:00"`
	DeletedBy datatypes.NullInt  `json:"-" gorm:"column:deleted_by" example:"0"`
	Ctx       *ctx.Ctx           `json:"-" gorm:"-"`
}
