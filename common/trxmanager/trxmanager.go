package trxmanager

import (
	"errors"
	"fmt"

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
	tx := t.db.Begin()
	var err error
	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			err = errors.New("WithTrx Error: " + fmt.Sprintf("%v", p))
		} else if err != nil {
			// error occurred, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit().Error
		}
	}()

	err = fn(ctx)
	return err
}
