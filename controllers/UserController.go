package controllers

import (
	"chatApp/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := database.DBManager()
	db.Find("")
	return c.JSON(http.StatusOK, "")
}
