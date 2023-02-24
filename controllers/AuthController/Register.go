package authcontroller

import (
	"net/http"
	variables "setup/controllers/Variables"
	"setup/database"
	"setup/helpers"
	"setup/localization"
	"setup/models"
	language "setup/resources/Language"
	emails "setup/resources/Views/Emails"

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

	err := helpers.Validate(c, input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: err.Error(),
		})
	}

	checkEmail := db.Where("email = ?", input.Email).Find(&user)
	if checkEmail.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.EMAIL_ALREADY_EXIST, c),
		})
	}

	if input.Password != input.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.PASSWORDS_DONT_MATCH, c),
		})
	}

	user.Password = helpers.Hash(input.Password)
	user.Email = input.Email
	user.Name = input.Name
	user.Username = input.Username

	if err := db.Create(&user); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			variables.MESSAGE: localization.TR(language.SOMETHING_WENT_WRONG, c),
		})
	}

	if err := helpers.SendEmail(user.Email, localization.TR(language.NEW_ACCOUNT, c), emails.Register(user.Name, c)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			variables.MESSAGE: localization.TR(language.SOMETHING_WENT_WRONG, c),
		})
	}

	return c.JSON(http.StatusOK, &user)
}
