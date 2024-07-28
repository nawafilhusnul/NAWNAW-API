package auth

import (
	"errors"
	"net/http"

	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	"gorm.io/gorm"
)

type Repository interface {
	Login(ctx *cc.Ctx, identifier, password string) (*model.Auth, error)
	Register(ctx *cc.Ctx, user *model.Auth) error
	GetOne(ctx *cc.Ctx, id int) (*model.User, error)
	FindUserRoles(ctx *cc.Ctx, userID int) ([]model.Role, error)
	FindUserPlatforms(ctx *cc.Ctx, userID int) ([]model.Platform, error)
	AssignDefaultPlatform(ctx *cc.Ctx, userID int, platformSlugs ...string) error
	AssignDefaultRole(ctx *cc.Ctx, userID int, roleSlugs ...string) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) checkTrx(ctx *cc.Ctx) {
	if ctx.Tx != nil {
		r.db = ctx.Tx
	}
}

func NewAuthMySQLRepo(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Login(ctx *cc.Ctx, identifier, password string) (*model.Auth, error) {
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

func (r *repository) Register(ctx *cc.Ctx, user *model.Auth) error {
	r.checkTrx(ctx)

	err := r.db.WithContext(ctx.RequestContext()).Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.NewError(http.StatusConflict, constants.ErrorCodeUserAlreadyExists, "User already exists")
		}
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to register user: "+err.Error())
	}

	return nil
}

func (r *repository) GetOne(ctx *cc.Ctx, id int) (*model.User, error) {
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

func (r *repository) FindUserRoles(ctx *cc.Ctx, userID int) ([]model.Role, error) {
	roles := []model.Role{}
	err := r.db.WithContext(ctx.RequestContext()).Table("roles").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get user roles: "+err.Error())
	}
	return roles, nil
}

func (r *repository) FindUserPlatforms(ctx *cc.Ctx, userID int) ([]model.Platform, error) {
	platforms := []model.Platform{}
	err := r.db.WithContext(ctx.RequestContext()).Table("platforms").
		Joins("JOIN user_platforms ON user_platforms.platform_id = platforms.id").
		Where("user_platforms.user_id = ?", userID).
		Find(&platforms).Error
	if err != nil {
		return nil, response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get user platforms: "+err.Error())
	}
	return platforms, nil
}

func (r *repository) AssignDefaultPlatform(ctx *cc.Ctx, userID int, platformSlugs ...string) error {
	r.checkTrx(ctx)

	userPlatforms := []model.UserPlatform{}
	for _, platformSlug := range platformSlugs {
		platform := model.Platform{}
		err := r.db.WithContext(ctx.RequestContext()).Where("slug = ?", platformSlug).First(&platform).Error
		if err != nil {
			return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get platform: "+err.Error())
		}
		userPlatforms = append(userPlatforms, model.UserPlatform{
			UserID:     userID,
			PlatformID: platform.ID,
		})
	}
	err := r.db.WithContext(ctx.RequestContext()).Create(&userPlatforms).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to assign default platform: "+err.Error())
	}
	return nil
}

func (r *repository) AssignDefaultRole(ctx *cc.Ctx, userID int, roleSlugs ...string) error {
	r.checkTrx(ctx)

	userRoles := []model.UserRole{}
	for _, roleSlug := range roleSlugs {
		role := model.Role{}
		err := r.db.WithContext(ctx.RequestContext()).Where("slug = ?", roleSlug).First(&role).Error
		if err != nil {
			return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to get role: "+err.Error())
		}
		userRoles = append(userRoles, model.UserRole{
			UserID: userID,
			RoleID: int(role.ID),
		})
	}
	err := r.db.WithContext(ctx.RequestContext()).Create(&userRoles).Error
	if err != nil {
		return response.NewError(http.StatusInternalServerError, constants.ErrorCodeInternalServerError, "Failed to assign default role: "+err.Error())
	}
	return nil
}
