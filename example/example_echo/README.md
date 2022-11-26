## Example Echo

```bash
go run example/example_echo/main.go
# In another tab,
$ curl localhost:1323/fuga
{"fuga":1000}
$ curl localhost:1323/hoge
{"hoge":1000}
$ curl localhost:1323/piyo
{"piyo":1000}
$ curl localhost:1323/fprof_result
FProf Result [us]
Sum           60, Max           60, Avg           60, Min           60, Count            1, L11 main.Hoge
Sum           63, Max           63, Avg           63, Min           63, Count            1, L17 main.Fuga
Sum          135, Max            6, Avg            0, Min            0, Count         1000, L28 main.Piyo
```

以下のような元のコードに対して、
開始時に`fprof.InitFProf()`、各関数の始めに`defer fprof.FProf()()`を追加し、結果を吐き出すエンドポイント`fprof.AnalizeFProfResult()`を作ってプロファイリングします。
Piyoでは、`fpr := fprof.FProf()`と`fpr()`で挟むことで、好きな区間を計測します。ただし、シンプルな作り故にオーバーヘッドがそれなりにあります。

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shiba6v/fprof/v1"
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

func Piyo(c echo.Context) error {
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
	e.GET("/piyo", Piyo)
	e.Logger.Fatal(e.Start(":1323"))
}
```