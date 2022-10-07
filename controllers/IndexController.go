package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to index")
}
