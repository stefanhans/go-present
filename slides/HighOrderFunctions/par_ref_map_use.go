package main

import (
	"fmt"
	"runtime"
	"time"

	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
)

func main() {
	tenTimes := func(x int) int {
		time.Sleep(time.Duration(1 * time.Millisecond))
		return x * 10
	}
	var list ListOfInt
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
	start := time.Now()
	fmt.Printf("%v.ParRefMap(tenTimes, %v) ", list, runtime.NumCPU())
	list.ParRefMap(tenTimes, runtime.NumCPU())
	fmt.Printf("yields %v\n", list)
	fmt.Print(time.Since(start))
}

// END OMIT