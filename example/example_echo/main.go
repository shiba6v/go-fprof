package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shiba6v/fprof/v1"
)

func Hoge(c echo.Context) error {
	defer fprof.FProf()()
	t := 1000
	return c.JSON(http.StatusOK, map[string]int{"hoge": t})
}

func Fuga(c echo.Context) error {
	defer fprof.FProf()()
	t := 0
	for i := 0; i < 1000; i++ {
		t += 1
	}
	return c.JSON(http.StatusOK, map[string]int{"fuga": t})
}

func Piyo(c echo.Context) error {
	t := 0
	for i := 0; i < 1000; i++ {
		fpr := fprof.FProf()
		t += 1
		fpr()
	}
	return c.JSON(http.StatusOK, map[string]int{"piyo": t})
}

func GetAnalizeFProfResult(c echo.Context) error {
	result := fprof.AnalizeFProfResult()
	return c.String(http.StatusOK, result)
}

func main() {
	fprof.InitFProf()
	e := echo.New()
	e.GET("/hoge", Hoge)
	e.GET("/fuga", Fuga)
	e.GET("/piyo", Piyo)
	e.GET("/fprof_result", GetAnalizeFProfResult)
	e.Logger.Fatal(e.Start(":1323"))
}
