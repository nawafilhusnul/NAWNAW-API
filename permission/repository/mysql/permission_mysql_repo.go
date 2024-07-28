package permission

import (
	"net/http"

	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx *cc.Ctx, permission *model.Permission) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) checkTrx(ctx *cc.Ctx) {
	if ctx.Tx != nil {
		r.db = ctx.Tx
	}
}

func NewPermissionMySQLRepo(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx *cc.Ctx, permission *model.Permission) error {
	r.checkTrx(ctx)

	err := r.db.WithContext(ctx.RequestContext()).Create(permission).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return nil
}

func (r *repository) Update(ctx *cc.Ctx, permission *model.Permission) error {
	err := r.db.WithContext(ctx.RequestContext()).Updates(permission).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return nil
}
