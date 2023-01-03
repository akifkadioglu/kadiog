package testinitializer

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

var Network = echo.New()

func Method(method,json string) (echo.Context, *httptest.ResponseRecorder) {
	
	req := httptest.NewRequest(method, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := Network.NewContext(req, rec)
	return c, rec
}
