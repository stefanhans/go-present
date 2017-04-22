package main

import (
	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
	"fmt"
)

func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}

	tenTimes := func(x int) int { return x * 10 }
	isEven := func(x int) bool { return x%2 == 0 }

	fmt.Printf("List %v: Map(tenTimes).Filter(isEven) yields %v\n", list,
		list.
			Map(tenTimes).
			Filter(isEven))
}
