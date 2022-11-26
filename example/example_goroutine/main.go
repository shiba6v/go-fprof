package main

import (
	"fmt"
	"sync"

	"github.com/shiba6v/fprof"
)

var wg sync.WaitGroup

func A() {
	defer fprof.FProf()()
	defer wg.Done()
	sum := 0
	for i := 0; i < 100; i++ {
		sum += B()
	}
	fmt.Println(sum)
}

func B() int {
	defer fprof.FProf()()
	t := func() int {
		return 1
	}()
	return t
}

func main() {
	fprof.InitFProf()

	wg.Add(1)
	go A()

	wg.Wait()
	r := fprof.AnalizeFProfResult()
	fmt.Println(r)
}
