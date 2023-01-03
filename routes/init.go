package routes

import (
	"net/http"
	_ "net/http/pprof"
	"setup/adapter"
	"setup/environment"
	"setup/helpers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	var E = echo.New()
	Network := E.Group("", adapter.ConsoleAdapter)
	Network.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	Web(Network)
	E.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
	port := helpers.GoDotEnvVariable(environment.HOST) + ":" + helpers.GoDotEnvVariable(environment.PORT)
	if port == ":" {
		port = "0.0.0.0:9000" // Default port if not specified
	}
	E.Logger.Fatal(E.Start(port))
}
