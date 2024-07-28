package routes

import (
	"github.com/labstack/echo/v4"
	authHandler "github.com/nawafilhusnul/NAWNAW-API/auth/delivery/http/handler"
	authRepo "github.com/nawafilhusnul/NAWNAW-API/auth/repository/mysql"
	authUsecase "github.com/nawafilhusnul/NAWNAW-API/auth/usecase"
	"github.com/nawafilhusnul/NAWNAW-API/middleware"
	"gorm.io/gorm"
)

// RegisterV1AuthRoutes registers the authentication routes for version 1 of the API.
// It sets up the necessary repository, usecase, and handler for authentication.
//
// Example usage:
// e := echo.New()
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// routes.RegisterV1AuthRoutes(e, db)
func RegisterV1AuthRoutes(v1 *echo.Group, db *gorm.DB) {
	authRepo := authRepo.NewAuthMySQLRepo(db)
	authUsecase := authUsecase.New(authRepo, db)
	authHandler := authHandler.New(authUsecase)

	g := v1.Group("/auths")
	g.POST("/login", authHandler.Login()).Name = "auths.Login"
	g.POST("/register", authHandler.Register()).Name = "auths.Register"
	g.GET("", authHandler.GetOne(), middleware.Session(), middleware.Platform(), middleware.Permission()).Name = "auths.GetOne"
}
