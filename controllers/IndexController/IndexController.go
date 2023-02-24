package indexcontroller

import (
	"net/http"
	"setup/localization"
	language "setup/resources/Language"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, localization.TR(language.WELCOME_TO_INDEX, c))
}
