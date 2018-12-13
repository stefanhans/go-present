package main

import "fmt"

// START OMIT
type ListOfInt []int

func (list ListOfInt) Fold(f func(int, int) int) int {
	var out int // HL
	for _, i := range list {
		out = f(out, i)
	}
	return out
}

// HALF OMIT

func main() {
	list := ListOfInt{-2, -1, 2, 2, 3}

	fmt.Printf("list%v.Fold(func(x, y int) int { return x + y }) yields %v\n", // HL
		list, list.Fold(func(x, y int) int { return x + y })) // HL
}

// END OMIT
