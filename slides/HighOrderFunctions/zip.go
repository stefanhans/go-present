package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

type listZipFunc func(x, y int) int

func (list ListOfInt) Zip(otherList ListOfInt, f func(int, int) int) ListOfInt {
	var out ListOfInt

	for n, i := range list {
		out = append(out, f(i, otherList[n]))
	}
	return out
}
// END OMIT

func main() {
	var list1 = ListOfInt{-20, -1, 0, 5, 2, 3}
	var list2 = ListOfInt{-2, -10, 0, 2, 4, 3}

	max := func(x, y int) int {
		if x > y { return x }
		return y
	}
	fmt.Printf("list%v: \nZip(%v, max) returns list%v\n",
		list1, list2, list1.Zip(list2, max))
}
