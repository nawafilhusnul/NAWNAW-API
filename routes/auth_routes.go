package routes

import (
	"github.com/labstack/echo/v4"
	authHandler "github.com/nawafilhusnul/big-app/auth/delivery/http/handler"
	authRepo "github.com/nawafilhusnul/big-app/auth/repository/mysql"
	authUsecase "github.com/nawafilhusnul/big-app/auth/usecase"
	"gorm.io/gorm"
)

func RegisterV1AuthRoutes(e *echo.Echo, db *gorm.DB) {
	authRepo := authRepo.NewAuthMySQLRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	v1 := e.Group("/api/v1")
	g := v1.Group("/auth")
	g.POST("/login", authHandler.Login())
	g.POST("/register", authHandler.Register())
}
