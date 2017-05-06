package main

import (
	"fmt"
	"time"
)

type ListOfInt []int

type listMapFunc func(int) int

func (list *ListOfInt) RefMap(f listMapFunc) {
	for i := 0; i < len(*list); i++ {
		(*list)[i] = f((*list)[i])
	}
}

func main() {
	tenTimes := func(x int) int {
		time.Sleep(time.Duration(1 * time.Millisecond))
		return x * 10
	}
	var list ListOfInt
	for i := 0; i < 10; i++ { list = append(list, i) }
	start := time.Now()
	fmt.Printf("%v.RefMap(tenTimes) ", list)
	list.RefMap(tenTimes)
	fmt.Printf("yields %v\n", list)
	fmt.Print(time.Since(start))
}

// END OMIT