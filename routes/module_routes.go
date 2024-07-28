package routes

import (
	"github.com/labstack/echo/v4"
	moduleHandler "github.com/nawafilhusnul/NAWNAW-API/module/delivery/http/handler"
	moduleRepo "github.com/nawafilhusnul/NAWNAW-API/module/repository/mysql"
	moduleUsecase "github.com/nawafilhusnul/NAWNAW-API/module/usecase"
	"gorm.io/gorm"
)

func RegisterV1ModuleRoutes(v1 *echo.Group, db *gorm.DB) {
	moduleRepo := moduleRepo.NewModuleMySQLRepo(db)
	moduleUsecase := moduleUsecase.New(moduleRepo, db)
	moduleHandler := moduleHandler.New(moduleUsecase)

	g := v1.Group("/modules")
	g.GET("", moduleHandler.FindAll()).Name = "modules.FindAll"
}
