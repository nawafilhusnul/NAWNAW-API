package trxmanager

import (
	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"gorm.io/gorm"
)

type TrxManager struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TrxManager {
	return &TrxManager{db: db}
}

func (t *TrxManager) WithTrx(ctx *ctx.Ctx, fn func(ctx *ctx.Ctx) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		ctx.Tx = tx
		return fn(ctx)
	})
}
