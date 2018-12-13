package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

func (list ListOfInt) GroupBy(f func(int, int) int) map[int]int {
	out := make(map[int]int)
	for _, i := range list {
		out[i] = f(i, out[i])
	}
	return out
}

// END OMIT

func main() {
	list := ListOfInt{1, 1, 2, 3, 3, 4, 4}
	count := func(i int, old int) int {
		return old + 1
	}
	fmt.Printf("list%v.GroupBy(count) yields %v\n", list, list.GroupBy(count))
}
