package authcontroller

import (
	"net/http"
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
)

type RegisterInput struct {
	Email                string `json:"email" validate:"required,email"`
	Name                 string `json:"name" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	Username             string `json:"username" validate:"required"`
}

func Register(c echo.Context) error {
	db := database.DBManager()
	var input RegisterInput
	var user models.User

	if err := c.Bind(&input); err != nil {
		return err
	}

	//validate inputs
	err := helpers.Validate(c, input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	//check email
	checkEmail := db.Where("email = '" + input.Email + "'").Find(&user)
	if checkEmail.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email Already Exist",
		})
	}

	if input.Password != input.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Passwords Do Not Match",
		})
	}

	//hashing
	user.Password = helpers.Hash(input.Password)
	user.Email = input.Email
	user.Name = input.Name
	user.Username = input.Username
	//creating
	db.Create(&user)

	return c.JSON(http.StatusOK, &user)
}
