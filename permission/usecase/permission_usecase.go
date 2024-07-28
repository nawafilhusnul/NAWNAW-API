package permission

import (
	"strings"

	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	module "github.com/nawafilhusnul/NAWNAW-API/module/repository/mysql"
	permission "github.com/nawafilhusnul/NAWNAW-API/permission/repository/mysql"
	"gorm.io/gorm"
)

type Usecase interface {
	Create(c *cc.Ctx, permission *model.Permission) error
}

type usecase struct {
	repo       permission.Repository
	moduleRepo module.Repository
	db         *gorm.DB
}

func New(repo permission.Repository, moduleRepo module.Repository, db *gorm.DB) Usecase {
	return &usecase{repo: repo, moduleRepo: moduleRepo, db: db}
}

func (uc *usecase) Create(c *cc.Ctx, permission *model.Permission) error {
	module, err := uc.moduleRepo.FindByID(c, int(permission.ModuleID))
	if err != nil {
		return err
	}

	permission.ModuleID = datatypes.ID(module.ID)

	lowerModuleName := strings.ToLower(module.Name)
	sanitizedPermissionName := strings.ReplaceAll(permission.Name.String, " ", "")
	slug := lowerModuleName + "." + sanitizedPermissionName
	permission.Slug = datatypes.SetNullString(slug)
	permission.Name = datatypes.SetNullString(module.Name + ": " + sanitizedPermissionName)

	return uc.repo.Create(c, permission)
}
