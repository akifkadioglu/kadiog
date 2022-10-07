package adapter

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func ConsoleAdapter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Request().URL
		fmt.Print(strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute()) + ":" + strconv.Itoa(time.Now().Second()) + " - \t")
		fmt.Println(url)
		return next(c)
	}
}
