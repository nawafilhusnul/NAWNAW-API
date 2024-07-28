package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	customValidator "github.com/nawafilhusnul/NAWNAW-API/common/validator"
	"github.com/nawafilhusnul/NAWNAW-API/config"
	"github.com/nawafilhusnul/NAWNAW-API/routes"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	// Initialize database connection (replace with actual connection details)
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	appConfig := config.LoadAppConfig()

	dbConfig := config.NewDatabase()
	db := dbConfig.GetDB()

	// error handler
	e.HTTPErrorHandler = response.CustomHTTPErrorHandler

	// Initialize validator
	e.Validator = customValidator.NewCustomValidator()

	// Middleware
	e.Use(
		middleware.Recover(),
		middleware.RemoveTrailingSlashWithConfig(
			middleware.TrailingSlashConfig{
				RedirectCode: http.StatusMovedPermanently,
			},
		),
		ctx.NewCtx,
	)

	// V1 Routes
	v1 := e.Group("/api/v1")
	routes.RegisterV1AuthRoutes(v1, db)
	routes.RegisterV1ModuleRoutes(v1, db)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port)))
}
