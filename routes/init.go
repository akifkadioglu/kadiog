package routes

import (
	"chatApp/env"

	"github.com/labstack/echo/v4"
)

var E = echo.New()

func Init() {
	Api()
	Web()
	port := ":" + env.GoDotEnvVariable("APP_PORT")
	E.Logger.Fatal(E.Start(port))

}
