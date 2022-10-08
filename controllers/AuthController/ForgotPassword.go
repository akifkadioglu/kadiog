package authcontroller

import (
	"net/http"
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
)

type ForgotPasswordInputs struct {
	Email string `json:"email"`
}

func ForgotPassword(c echo.Context) error {
	db := database.DBManager()
	var input ForgotPasswordInputs
	var user models.User
	c.Bind(&input)
	db.Where("`email` ='" + input.Email + "'").Find(&user).First(&user)
	//fmt.Println(&c.Request().Host)
	newPasswordUri := string(rune(user.ID))
	text := "Hello " + user.Name + " <a href='facebook.com'>click me</a> "
	helpers.SendEmail(user.Email, "nassın", text+newPasswordUri)
	return c.JSON(http.StatusOK, "sended")
}
