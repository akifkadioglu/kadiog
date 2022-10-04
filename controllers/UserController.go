package controllers

import (
	"net/http"
	"setup/database"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := database.DBManager()
	db.Find("")
	return c.JSON(http.StatusOK, "")
}
