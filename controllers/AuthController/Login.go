package authcontroller

import (
	"net/http"
	"time"

	variables "setup/controllers/Variables"
	"setup/database"
	"setup/env"
	"setup/helpers"
	models "setup/models"
	language "setup/resources/Language"
	"setup/localization"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {

	var input LoginInput
	var user models.User
	db := database.DBManager()
	c.Bind(&input)

	if err := helpers.Validate(c, input); err != nil {
		return err
	}

	if err := db.Where("email = ?", input.Email).First(&user); err.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.PASSWORD_AND_EMAIL_DONT_MATCH, c),
		})
	}

	if err := helpers.CompareHash(user.Password, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.PASSWORD_AND_EMAIL_DONT_MATCH, c),
		})
	}

	var claims = &models.JwtCustomClaims{
		UserId: int(user.ID),
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(), // For 1 year permission
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(env.Getenv(env.APP_KEY)))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		variables.TOKEN: t,
	})
}
