package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"setup/adapter"
	"setup/env"
)

var E = echo.New()
var Network = E.Group("", adapter.ConsoleAdapter)

func Init() {
	E.Use(middleware.Logger())
	E.Use(middleware.Recover())
	Api()
	Web()
	port := ":" + env.GoDotEnvVariable("APP_PORT")
	E.Logger.Fatal(E.Start(port))
}
