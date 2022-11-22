package authcontroller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"setup/database"
	"setup/helpers"
	models "setup/models"
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

	//validate inputs
	err := helpers.Validate(c, input)
	if err != nil {
		return err
	}

	//check email
	checkEmail := db.Where("email = '" + input.Email + "'").First(&user)
	if checkEmail.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Password And Email Do Not Match",
		})
	}

	err = helpers.CompareHash(user.Password, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Password And Email Do Not Match",
		})	}

	// Set custom claims
	var claims = &models.JwtCustomClaims{
		UserId: int(user.ID),
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(), // For 1 year permission
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(helpers.GoDotEnvVariable("APP_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
