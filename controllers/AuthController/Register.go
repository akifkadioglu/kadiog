package authcontroller

import (
	"errors"
	"net/http"
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

func Register(c echo.Context) error {
	db := database.DBManager()
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	//validate inputs
	err := registerValidate(c, user)
	if err != nil {
		return err
	}

	//check username is exist
	err = checkDBFor(db, c, "username", user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "this username is already exist")
	}

	//check email is exist
	err = checkDBFor(db, c, "email", user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "this email is already exist")
	}

	//hashing
	user.Password = helpers.Hash(user.Password)

	//creating
	db.Create(&user)

	return c.JSON(http.StatusOK, &user)
}

func registerValidate(c echo.Context, user models.User) error {

	v := validator.New()

	if err := v.Struct(user); err != nil {
		return err
	}
	return nil
}

func checkDBFor(db gorm.DB, c echo.Context, column string, value string) error {
	var user models.User

	query := "`" + column + "` = '" + value + `'`
	result := db.Where(query).Find(&user)
	if result.RowsAffected > 0 {
		return errors.New("an error")
	}
	return nil
}
