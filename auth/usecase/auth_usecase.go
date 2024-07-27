package auth

import (
	"net/http"
	"time"

	auth "github.com/nawafilhusnul/NAWNAW-API/auth/repository/mysql"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
	"github.com/nawafilhusnul/NAWNAW-API/common/helper"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/common/token"
	"github.com/nawafilhusnul/NAWNAW-API/model"
)

type Usecase interface {
	Login(ctx *ctx.Ctx, identifier, password, tz string) (*model.Auth, error)
	Register(ctx *ctx.Ctx, user *model.Auth) error
	GetOne(ctx *ctx.Ctx, id int) (*model.User, error)
}

type usecase struct {
	repo auth.Repository
}

func NewAuthUsecase(repo auth.Repository) Usecase {
	return &usecase{repo: repo}
}

func (uc *usecase) Login(ctx *ctx.Ctx, identifier, password, tz string) (*model.Auth, error) {
	user, err := uc.repo.Login(ctx, identifier, password)
	if err != nil {
		return nil, err
	}

	_, err = time.LoadLocation(tz)
	if err != nil {
		return nil, response.NewError(http.StatusBadRequest, constants.ErrorCodeInvalidTimezone, "Invalid timezone")
	}

	if err := helper.ComparePassword(string(user.Password), password); err != nil {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidPassword, "Invalid password")
	}

	user.Timezone = tz

	accessToken, err := token.GenerateAccessToken(user)
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to generate access token")
	}

	refreshToken, err := token.GenerateRefreshToken(user)
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to generate refresh token")
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	return user, nil
}

func (uc *usecase) Register(ctx *ctx.Ctx, user *model.Auth) error {
	hashedPassword, err := helper.HashPassword(string(user.Password))
	if err != nil {
		return response.NewError(http.StatusBadRequest, constants.ErrorCodeInvalidPassword, "Failed to hash password")
	}
	user.Password = datatypes.HashString(hashedPassword)
	return uc.repo.Register(ctx, user)
}

func (uc *usecase) GetOne(ctx *ctx.Ctx, id int) (*model.User, error) {
	return uc.repo.GetOne(ctx, id)
}
