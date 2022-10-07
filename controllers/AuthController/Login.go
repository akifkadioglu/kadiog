package authcontroller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"

	"setup/database"
	"setup/env"
	models "setup/models"
)

type LoginInput struct {
	Username string `json:"username" gorm:"not null; size:60" validate:"required,min=3"`
	Password string `json:"password" gorm:"not null" validate:"required,min=5"`
}

func Login(c echo.Context) error {

	var input LoginInput
	var user models.User
	db := database.DBManager()
	c.Bind(&input)

	//validate inputs
	err := loginValidate(c, input)
	if err != nil {
		return err
	}

	//get user
	db.Where("`username` = '" + input.Username + "'").Find(&user).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, echo.ErrNotFound)
	}
	userPassword := []byte(user.Password)
	password := []byte(input.Password)

	err = bcrypt.CompareHashAndPassword(userPassword, password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	var claims = &models.JwtCustomClaims{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(env.GoDotEnvVariable("APP_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func loginValidate(c echo.Context, input LoginInput) error {

	v := validator.New()

	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
