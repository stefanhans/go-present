package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

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
// END OMIT

func main() {
	list := ListOfInt{-2, -1, 0, 2, 2, 3}
	isEven := func(x int) bool { return x%2 == 0 }

	fmt.Printf("List %v: Filter(isEven) yields %v\n",
		list, list.Filter(isEven))
}
