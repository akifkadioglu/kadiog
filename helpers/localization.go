package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

const (
	TURKISH string = "tr"
	ENGLISH string = "en"
)

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/Language/en.json")
	bundle.LoadMessageFile("resources/Language/tr.json")

}

func TR(key string, c echo.Context) string {

	language := c.Request().Header.Get("Accept-Language")
	if len(language) < 2 {
		language = ENGLISH
	}
	switch language[0:2] {
	case TURKISH:
		localizer = i18n.NewLocalizer(bundle, TURKISH)
	case ENGLISH:
		localizer = i18n.NewLocalizer(bundle, ENGLISH)
	default:
		localizer = i18n.NewLocalizer(bundle, ENGLISH)
	}
	localizedValue := i18n.LocalizeConfig{
		MessageID: key,
	}
	value, err := localizer.Localize(&localizedValue)
	if err != nil {
		fmt.Println(err.Error())
	}
	return value
}
