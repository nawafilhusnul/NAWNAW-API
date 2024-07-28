package module

import (
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	module "github.com/nawafilhusnul/NAWNAW-API/module/repository/mysql"
	"gorm.io/gorm"
)

type Usecase interface {
	Create(ctx *cc.Ctx, module *model.Module) error
	FindAll(ctx *cc.Ctx) ([]model.Module, error)
	FindByID(ctx *cc.Ctx, id int) (*model.Module, error)
	Update(ctx *cc.Ctx, module *model.Module) error
	Delete(ctx *cc.Ctx, id int) error
}

type usecase struct {
	repo module.Repository
	db   *gorm.DB
}

func New(repo module.Repository, db *gorm.DB) Usecase {
	return &usecase{repo: repo, db: db}
}

func (u *usecase) Create(ctx *cc.Ctx, module *model.Module) error {
	return u.repo.Create(ctx, module)
}

func (u *usecase) FindAll(ctx *cc.Ctx) ([]model.Module, error) {
	return u.repo.FindAll(ctx)
}

func (u *usecase) FindByID(ctx *cc.Ctx, id int) (*model.Module, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *usecase) Update(ctx *cc.Ctx, module *model.Module) error {
	return u.repo.Update(ctx, module)
}

func (u *usecase) Delete(ctx *cc.Ctx, id int) error {
	return u.repo.Delete(ctx, id)
}
