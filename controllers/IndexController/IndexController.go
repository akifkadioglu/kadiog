package indexcontroller

import (
	"net/http"
	"setup/helpers"
	language "setup/resources/Language"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, helpers.TR(language.WELCOME_TO_INDEX, c))
}
