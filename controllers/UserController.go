package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"setup/database"
	models "setup/models"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	db := database.DBManager()
	db.Find(&users)
	return c.JSON(http.StatusOK, map[string][]models.User{
		"users": users,
	})
}
