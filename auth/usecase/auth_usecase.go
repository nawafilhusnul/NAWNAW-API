package auth

import (
	"net/http"
	"time"

	auth "github.com/nawafilhusnul/NAWNAW-API/auth/repository/mysql"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
	"github.com/nawafilhusnul/NAWNAW-API/common/helper"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/common/token"
	"github.com/nawafilhusnul/NAWNAW-API/common/trxmanager"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	"gorm.io/gorm"
)

type Usecase interface {
	Login(c *cc.Ctx, identifier, password, tz string) (*model.Auth, error)
	Register(c *cc.Ctx, user *model.Auth) error
	GetOne(c *cc.Ctx, id int) (*model.User, error)
}

type usecase struct {
	repo auth.Repository
	db   *gorm.DB
}

func NewAuthUsecase(repo auth.Repository, db *gorm.DB) Usecase {
	return &usecase{repo: repo, db: db}
}

func (uc *usecase) Login(ctx *cc.Ctx, identifier, password, tz string) (*model.Auth, error) {
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

	roles, err := uc.repo.FindUserRoles(ctx, int(user.ID))
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get user roles")
	}

	userRoles := make(map[string]bool)
	for _, role := range roles {
		userRoles[role.Slug.String] = true
	}
	user.Roles = userRoles

	platforms, err := uc.repo.FindUserPlatforms(ctx, int(user.ID))
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get user platforms")
	}
	userPlatforms := make(map[string]bool)
	for _, platform := range platforms {
		userPlatforms[platform.Slug.String] = true
	}
	user.Platforms = userPlatforms

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

func (uc *usecase) Register(ctx *cc.Ctx, user *model.Auth) error {
	hashedPassword, err := helper.HashPassword(string(user.Password))
	if err != nil {
		return response.NewError(http.StatusBadRequest, constants.ErrorCodeInvalidPassword, "Failed to hash password")
	}
	user.Password = datatypes.HashString(hashedPassword)

	err = trxmanager.New(uc.db).WithTrx(ctx, func(ctx *cc.Ctx) error {
		err = uc.repo.Register(ctx, user)
		if err != nil {
			return err
		}

		err = uc.repo.AssignDefaultPlatform(ctx, int(user.ID), constants.DefaultPlatformSlugs...)
		if err != nil {
			return err
		}

		err = uc.repo.AssignDefaultRole(ctx, int(user.ID), constants.DefaultRoleSlugs...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) GetOne(ctx *cc.Ctx, id int) (*model.User, error) {
	return uc.repo.GetOne(ctx, id)
}
