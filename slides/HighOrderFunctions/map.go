package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

type listMapFunc func(int) int

func (list ListOfInt) Map(f listMapFunc) ListOfInt {
	var out ListOfInt
	for _, i := range list {
		out = append(out, f(i))
	}
	return out
}

// END OMIT

// MUTABLE MAP START OMIT
func (list ListOfInt) AnotherMap(f listMapFunc) {
	for n, i := range list {
		list[n] = f(i)
	}
}

func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}
	tenTimes := func(x int) int { return x * 10 }

	fmt.Printf("list%v: Map(tenTimes) returns list%v but still list%v\n",
		list, list.Map(tenTimes), list)

	fmt.Printf("list%v.AnotherMap(tenTimes) changed to ", list)
	list.AnotherMap(tenTimes)
	fmt.Printf("to list%v\n", list)
}
// MUTABLE MAP END OMIT

// END OMIT
