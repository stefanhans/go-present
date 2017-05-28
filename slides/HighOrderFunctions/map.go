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
func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}

	tenTimes := func(x int) int { return x * 10 }

	fmt.Printf("List %v: Map(tenTimes) yields %v\n", list, list.Map(tenTimes))
	fmt.Printf("and list%v is immutable\n", list)
}

func (list ListOfInt) MutableMap(f listMapFunc) ListOfInt {
	for n, i := range list {
		list[n] = f(i)
	}
	return list
}

// MUTABLE MAP END OMIT

// END OMIT