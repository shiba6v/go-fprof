package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hoge(c echo.Context) error {
	t := 1000
	return c.JSON(http.StatusOK, map[string]int{"hoge": t})
}

func Fuga(c echo.Context) error {
	t := 0
	for i := 0; i < 1000; i++ {
		t += 1
	}
	return c.JSON(http.StatusOK, map[string]int{"fuga": t})
}

func main() {
	e := echo.New()
	e.GET("/hoge", Hoge)
	e.GET("/fuga", Fuga)
	e.Logger.Fatal(e.Start(":1323"))
}
