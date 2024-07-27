package auth

import (
	"net/http"

	auth "github.com/nawafilhusnul/big-app/auth/repository/mysql"
	"github.com/nawafilhusnul/big-app/common/constants"
	"github.com/nawafilhusnul/big-app/common/ctx"
	"github.com/nawafilhusnul/big-app/common/helper"
	"github.com/nawafilhusnul/big-app/common/response"
	"github.com/nawafilhusnul/big-app/model"
)

type Usecase interface {
	Login(ctx *ctx.Ctx, identifier, password string) (*model.Auth, error)
	Register(ctx *ctx.Ctx, user *model.Auth) error
}

type usecase struct {
	repo auth.Repository
}

func NewAuthUsecase(repo auth.Repository) Usecase {
	return &usecase{repo: repo}
}

func (uc *usecase) Login(ctx *ctx.Ctx, identifier, password string) (*model.Auth, error) {
	return uc.repo.Login(ctx, identifier, password)
}

func (uc *usecase) Register(ctx *ctx.Ctx, user *model.Auth) error {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return response.NewError(http.StatusBadRequest, constants.ErrorCodeInvalidPassword, "Failed to hash password")
	}
	user.Password = hashedPassword
	return uc.repo.Register(ctx, user)
}
