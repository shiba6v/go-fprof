## Example Goroutine
以下のような元のコードに対して、
開始時に`fprof.InitFProf()`、終了時に`r := fprof.AnalizeFProfResult(); fmt.Println(r)`、各関数の始めに`defer fprof.FProf()()`を付けてプロファイリングします。

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