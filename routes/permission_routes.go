package routes

import (
	"github.com/labstack/echo/v4"
	moduleRepo "github.com/nawafilhusnul/NAWNAW-API/module/repository/mysql"
	permissionHandler "github.com/nawafilhusnul/NAWNAW-API/permission/delivery/http/handler"
	permissionRepo "github.com/nawafilhusnul/NAWNAW-API/permission/repository/mysql"
	permissionUsecase "github.com/nawafilhusnul/NAWNAW-API/permission/usecase"
	"gorm.io/gorm"
)

func RegisterV1PermissionRoutes(v1 *echo.Group, db *gorm.DB) {
	permissionRepo := permissionRepo.NewPermissionMySQLRepo(db)
	moduleRepo := moduleRepo.NewModuleMySQLRepo(db)
	permissionUsecase := permissionUsecase.New(permissionRepo, moduleRepo, db)
	permissionHandler := permissionHandler.New(permissionUsecase)

	g := v1.Group("/permissions")
	g.POST("", permissionHandler.Create()).Name = "permissions.Create"
}
