package routes

import (
	"setup/adapter"
	"setup/helpers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var E = echo.New()

func Init() {

	Network := E.Group("", adapter.ConsoleAdapter)
	Network.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	Web(Network)
	port := helpers.GoDotEnvVariable("HOST") + ":" + helpers.GoDotEnvVariable("PORT")
	if port == ":" {
		port = "0.0.0.0:9000" // Default port if not specified
	}
	E.Logger.Fatal(E.Start(port))
}
