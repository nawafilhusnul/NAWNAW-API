package module

import (
	"net/http"

	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx *cc.Ctx, module *model.Module) error
	FindAll(ctx *cc.Ctx) ([]model.Module, error)
	FindByID(ctx *cc.Ctx, id int) (*model.Module, error)
	Update(ctx *cc.Ctx, module *model.Module) error
	Delete(ctx *cc.Ctx, id int) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) checkTrx(ctx *cc.Ctx) {
	if ctx.Tx != nil {
		r.db = ctx.Tx
	}
}

func NewModuleMySQLRepo(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx *cc.Ctx, module *model.Module) error {
	r.checkTrx(ctx)

	err := r.db.WithContext(ctx.RequestContext()).Create(module).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return nil
}

func (r *repository) FindAll(ctx *cc.Ctx) ([]model.Module, error) {
	var modules []model.Module
	err := r.db.WithContext(ctx.RequestContext()).Find(&modules).Error
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return modules, nil
}

func (r *repository) FindByID(ctx *cc.Ctx, id int) (*model.Module, error) {
	var module model.Module
	err := r.db.WithContext(ctx.RequestContext()).First(&module, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, response.NewError(http.StatusNotFound, constants.ErrorCodeModuleNotFound, err.Error())
		}
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return &module, nil
}

func (r *repository) Update(ctx *cc.Ctx, module *model.Module) error {
	err := r.db.WithContext(ctx.RequestContext()).Model(&model.Module{}).Where("id = ?", module.ID).Updates(module).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return nil
}

func (r *repository) Delete(ctx *cc.Ctx, id int) error {
	err := r.db.WithContext(ctx.RequestContext()).Model(&model.Module{}).Where("id = ?", id).Delete(&model.Module{}).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, err.Error())
	}
	return nil
}
