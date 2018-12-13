package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

func (list ListOfInt) Zip(otherList ListOfInt, f func(int, int) int) ListOfInt {
	var out ListOfInt

	for n, i := range list {
		out = append(out, f(i, otherList[n]))
	}
	return out
}

// END OMIT

func main() {
	list1 := ListOfInt{-20, -1, 0, 5, 2, 3}
	list2 := ListOfInt{-2, -10, 0, 2, 4, 3}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	fmt.Printf("list%v.\nZip(%v, max) yields\nlist%v\n",
		list1, list2, list1.Zip(list2, max))
}
