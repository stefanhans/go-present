package main

import "fmt"

// START OMIT
type ListOfInt []int

func (list ListOfInt) Fold(f func(int, int) int) int {
	var out int
	for _, i := range list {
		out = f(out, i)
	}
	return out
}
// END OMIT

func main() {
	list := ListOfInt{-2, -1, 2, 2, 3}
	sum := func(x, y int) int { return x + y }

	fmt.Printf("list%v.Fold(sum) yields %v\n",
		list, list.Fold(sum))
}
// END OMIT
