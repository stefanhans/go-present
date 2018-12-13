package main

import "fmt"

// START OMIT
type List []int

type ListFoldFunc func(int, int) int

func (list List) Fold(f ListFoldFunc) int {
	var out int
	for _, i := range list {
		out = f(out, i)
	}
	return out
}

func main() {
	var list = List{-2, -1, 2, 2, 3}

	sum := func(x, y int) int { return x + y }

	fmt.Printf("List %v: Fold(sum) yields %v\n", list, list.Fold(sum))
}

// END OMIT
