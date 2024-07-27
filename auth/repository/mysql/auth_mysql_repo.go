package auth

import (
	"net/http"

	"github.com/nawafilhusnul/big-app/common/constants"
	"github.com/nawafilhusnul/big-app/common/ctx"
	"github.com/nawafilhusnul/big-app/common/response"
	"github.com/nawafilhusnul/big-app/model"
	"gorm.io/gorm"
)

type Repository interface {
	Login(ctx *ctx.Ctx, identifier, password string) (*model.Auth, error)
	Register(ctx *ctx.Ctx, user *model.Auth) error
}

type repository struct {
	db *gorm.DB
}

func NewAuthMySQLRepo(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Login(ctx *ctx.Ctx, identifier, password string) (*model.Auth, error) {
	panic("not implemented")
}

func (r *repository) Register(ctx *ctx.Ctx, user *model.Auth) error {
	err := r.db.WithContext(ctx.RequestContext()).Create(user).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to register user")
	}

	return nil
}
