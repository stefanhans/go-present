package main

import (
	"fmt"
)

type ListOfInt []int

type listMapFunc func(int) int

func (list ListOfInt) Map(f listMapFunc) ListOfInt {
	var out ListOfInt
	for _, i := range list {
		out = append(out, f(i))
	}
	return out
}

type ListFilterFunc func(int) bool

func (list ListOfInt) Filter(f ListFilterFunc) ListOfInt {
	var out ListOfInt
	for _, i := range list {
		if f(i) {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}

	tenTimes := func(x int) int { return x * 10 }
	isEven := func(x int) bool { return x%2 == 0 }

	fmt.Printf("list%v.Filter(isEven).Map(tenTimes) yields %v\n", list,
		list.
			Filter(isEven).
			Map(tenTimes))
}

// END OMIT
