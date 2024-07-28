package auth

import (
	"errors"
	"net/http"

	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	"gorm.io/gorm"
)

type Repository interface {
	Login(ctx *ctx.Ctx, identifier, password string) (*model.Auth, error)
	Register(ctx *ctx.Ctx, user *model.Auth) error
	GetOne(ctx *ctx.Ctx, id int) (*model.User, error)
	FindUserRoles(ctx *ctx.Ctx, userID int) ([]model.Role, error)
	FindUserPlatforms(ctx *ctx.Ctx, userID int) ([]model.Platform, error)
}

type repository struct {
	db *gorm.DB
}

func NewAuthMySQLRepo(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Login(ctx *ctx.Ctx, identifier, password string) (*model.Auth, error) {
	user := &model.Auth{}
	err := r.db.WithContext(ctx.RequestContext()).Where("email = ? OR phone = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewError(http.StatusNotFound, constants.ErrorCodeUserNotFound, "User not found")
		}
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get user: "+err.Error())
	}
	return user, nil
}

func (r *repository) Register(ctx *ctx.Ctx, user *model.Auth) error {
	err := r.db.WithContext(ctx.RequestContext()).Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.NewError(http.StatusConflict, constants.ErrorCodeUserAlreadyExists, "User already exists")
		}
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to register user: "+err.Error())
	}

	return nil
}

func (r *repository) GetOne(ctx *ctx.Ctx, id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.WithContext(ctx.RequestContext()).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewError(http.StatusNotFound, constants.ErrorCodeUserNotFound, "User not found")
		}
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get user: "+err.Error())
	}
	return user, nil
}

func (r *repository) FindUserRoles(ctx *ctx.Ctx, userID int) ([]model.Role, error) {
	roles := []model.Role{}
	err := r.db.WithContext(ctx.RequestContext()).Table("roles").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *repository) FindUserPlatforms(ctx *ctx.Ctx, userID int) ([]model.Platform, error) {
	platforms := []model.Platform{}
	err := r.db.WithContext(ctx.RequestContext()).Table("platforms").
		Joins("JOIN user_platforms ON user_platforms.platform_id = platforms.id").
		Where("user_platforms.user_id = ?", userID).
		Find(&platforms).Error
	return platforms, err
}
