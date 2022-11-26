## Example Goroutine
```
$ go run example/example_goroutine/main.go 
100
FProf Result [us]
Sum          150, Max          150, Avg          150, Min          150, Count            1, L13 main.A
Sum           11, Max            1, Avg            0, Min            0, Count          100, L23 main.B
```

以下のような元のコードに対して、
開始時に`fprof.InitFProf()`、各関数の始めに`defer fprof.FProf()()`を付け、終了時に`r := fprof.AnalizeFProfResult(); fmt.Println(r)`でプロファイリング結果を出力します。

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func A() {
	defer wg.Done()
	sum := 0
	for i := 0; i < 100; i++ {
		sum += B()
	}
	fmt.Println(sum)
}

func B() int {
	t := func() int {
		return 1
	}()
	return t
}

func main() {
	wg.Add(1)
	go A()
	wg.Wait()
}

```