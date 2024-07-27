package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nawafilhusnul/big-app/common/ctx"
	"github.com/nawafilhusnul/big-app/config"
	"github.com/nawafilhusnul/big-app/routes"
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

	// Middleware
	e.Use(ctx.NewCtx)

	// Routes
	routes.RegisterV1AuthRoutes(e, db)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port)))
}
