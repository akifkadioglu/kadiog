package routes

import (
	"setup/adapter"
	"setup/helpers"
	_ "net/http/pprof"
	"net/http"
	
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
	port := helpers.GoDotEnvVariable("HOST") + ":" + helpers.GoDotEnvVariable("PORT")
	if port == ":" {
		port = "0.0.0.0:9000" // Default port if not specified
	}
	E.Logger.Fatal(E.Start(port))
}
