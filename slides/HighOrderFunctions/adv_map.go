package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

func (list ListOfInt) Map(f func(int) int) ListOfInt {
	var out ListOfInt
	for _, i := range list {
		out = append(out, f(i))
	}
	return out
}

// END OMIT

func main() {
	list := ListOfInt{-2, -1, 0, 2, 2, 3}
	tenTimes := func(x int) int { return x * 10 }

	fmt.Printf("list%v.Map(tenTimes) yields %v\n",
		list, list.Map(tenTimes))
}
