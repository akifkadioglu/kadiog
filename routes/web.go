package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Web() *echo.Echo {

	E.GET("/index", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "index")
	})
	return E
}
