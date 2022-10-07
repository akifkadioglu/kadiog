package routes

import (
	"setup/adapter"
	"setup/env"
	"github.com/labstack/echo/v4"
)

var E = echo.New()
var Network = E.Group("", adapter.ConsoleAdapter)

func Init() {
	Api()
	Web()
	port := ":" + env.GoDotEnvVariable("APP_PORT")
	E.Logger.Fatal(E.Start(port))
}
