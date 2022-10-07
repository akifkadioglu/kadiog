package routes

import (
	jwt "github.com/golang-jwt/jwt"
	"net/http"
	authcontroller "setup/controllers/AuthController"
	"setup/helpers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Api() {
	api := Network.Group("/api")

	//Authenticate
	api.POST("/register", authcontroller.Register)
	api.POST("/login", authcontroller.Login)

	auth := api.Group("")
	config := middleware.JWTConfig{
		KeyFunc: helpers.GetKey,
	}
	auth.Use(middleware.JWTWithConfig(config))
	auth.GET("/naber", restricted)
	//api.POST("/login", authcontroller.Login)
}
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
