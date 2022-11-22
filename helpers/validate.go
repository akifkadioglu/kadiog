package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Validate(c echo.Context, input interface{}) error {

	v := validator.New()

	if err := v.Struct(input); err != nil {
		return err
	}

	return nil
}

