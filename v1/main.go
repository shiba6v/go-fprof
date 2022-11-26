package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

var FProfStorage []*FProfElement
var FProfStorageMutex sync.Mutex

// key: lineNumber, value: functionName
// 1行に2つ以上の関数がないことを前提としている。
var FProfFuncNameMap map[uint16]string
var FProfFuncNameMapMutex sync.RWMutex

type FProfElement struct {
	// Elementあたり10 Byteで、100万回呼ばれても10MBなのでメモリに乗せても大丈夫。
	LineNumber  uint16
	ElapsedTime uint32
}

type FProfAnalyzeElement struct {
	Sum   uint64
	Count uint64
	Max   uint32
	Min   uint32
}

func InitFProf() { // FPROF_IGNORE
	FProfStorageMutex = sync.Mutex{}
	FProfFuncNameMapMutex = sync.RWMutex{}
	FProfStorage = make([]*FProfElement, 0, 1024)
	// FProfFuncNameMap = map[uint16]string{}
}

func FProf() func() { // FPROF_IGNORE
	pt, _, l, ok := runtime.Caller(1)
	lineNumber := uint16(l)
	if !ok {
		fmt.Println("Warning: runtime.Caller(1) failed in FProf()")
		return func() {} // FPROF_IGNORE
	}
	// FProfFuncNameMapMutex.RLock()
	// _, ok = FProfFuncNameMap[lineNumber]
	// // TODO: ここの実装は怪しい。
	// // 書き込みロックをかけるには、読み込みロックを解除しなければならないが、読み込みロックを外した後にほかから読み込まれるかもしれない。
	// // しかし、全部を書き込みロックで取るのはパフォーマンスを低下させる可能性がある。(ほとんどの場合が読み込みロックで十分なため。)
	// // 今回はmapに値を入れる冪等な操作なので、panicにさえならなければ良い。
	// FProfFuncNameMapMutex.RUnlock()
	// if !ok {
	// 	FProfFuncNameMapMutex.Lock()
	// 	funcName := runtime.FuncForPC(pt).Name()
	// 	FProfFuncNameMap[lineNumber] = funcName
	// 	FProfFuncNameMapMutex.Unlock()
	// }

	FProfFuncNameMapMutex.Lock()
	if FProfFuncNameMap == nil {
		log.Print("init FProfFuncNameMap")
		FProfFuncNameMap = map[uint16]string{}
	}
	_, ok = FProfFuncNameMap[lineNumber]
	if !ok {
		funcName := runtime.FuncForPC(pt).Name()
		FProfFuncNameMap[lineNumber] = funcName
	}
	FProfFuncNameMapMutex.Unlock()

	t1 := time.Now().UnixNano()
	return func() { // FPROF_IGNORE
		t2 := time.Now().UnixNano()
		if t2-t1 < 0 {
			fmt.Println("Warning: t2 < t1 in FProf()")
		}
		s := FProfElement{
			LineNumber:  lineNumber,
			ElapsedTime: uint32(t2 - t1),
		}
		FProfStorageMutex.Lock()
		FProfStorage = append(FProfStorage, &s)
		FProfStorageMutex.Unlock()
	}
}

func FProfMaxUint32(x uint32, y uint32) uint32 { // FPROF_IGNORE
	if x > y {
		return x
	} else {
		return y
	}
}

func FProfMinUint32(x uint32, y uint32) uint32 { // FPROF_IGNORE
	if x < y {
		return x
	} else {
		return y
	}
}

func analyzeFProfResultAggregate() map[uint16]*FProfAnalyzeElement { // FPROF_IGNORE
	FProfStorageMutex.Lock()
	aMap := make(map[uint16]*FProfAnalyzeElement)
	for _, p := range FProfStorage {
		a, ok := aMap[p.LineNumber]
		if !ok {
			a = &FProfAnalyzeElement{
				Sum:   0,
				Count: 0,
				Max:   0,
				Min:   math.MaxUint32,
			}
		}
		a.Sum += uint64(p.ElapsedTime)
		a.Count += 1
		a.Max = FProfMaxUint32(a.Max, p.ElapsedTime)
		a.Min = FProfMinUint32(a.Min, p.ElapsedTime)
		aMap[p.LineNumber] = a
	}
	FProfStorageMutex.Unlock()
	return aMap
}

func analyzeFProfResultGetLineNumbers() []uint16 { // FPROF_IGNORE
	lineNumbers := []uint16{}
	FProfFuncNameMapMutex.Lock()
	for k := range FProfFuncNameMap {
		lineNumbers = append(lineNumbers, k)
	}
	FProfFuncNameMapMutex.Unlock()
	sort.Slice(lineNumbers, func(i int, j int) bool { // FPROF_IGNORE
		return lineNumbers[i] < lineNumbers[j]
	})
	return lineNumbers
}

func analyzeFProfResultBuildResult(lineNumbers []uint16, aMap map[uint16]*FProfAnalyzeElement) string { // FPROF_IGNORE
	result := strings.Builder{}
	result.WriteString("FProf Result [us]\n")
	FProfFuncNameMapMutex.RLock()
	for _, line := range lineNumbers {
		a, ok := aMap[line]
		if !ok {
			fmt.Printf("Warning: key %d does not exists in aMap %d", line, len(aMap))
		} else {
			r := fmt.Sprintf("Sum %12d, Max %12d, Avg %12d, Min %12d, Count %12d, L%d %s\n",
				a.Sum/1000,
				a.Max/1000,
				a.Sum/(1000*a.Count),
				a.Min/1000,
				a.Count,
				line,
				FProfFuncNameMap[line],
			)
			result.WriteString(r)
		}
	}
	FProfFuncNameMapMutex.RUnlock()
	return result.String()
}

func AnalizeFProfResult() string { // FPROF_IGNORE
	// line number to aggregated time information.
	aMap := analyzeFProfResultAggregate()
	// line numbers of target functions
	lineNumbers := analyzeFProfResultGetLineNumbers()
	// create a profiling report
	result := analyzeFProfResultBuildResult(lineNumbers, aMap)
	return result
}

// func GetAnalizeFProfResult(c echo.Context) error { // FPROF_IGNORE
// 	result := AnalizeFProfResult()
// 	return c.String(http.StatusOK, result)
// }
