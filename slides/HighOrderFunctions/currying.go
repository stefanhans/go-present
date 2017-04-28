package main

import (
	"fmt"

	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
)

func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}

	tenTimes := func(x int) int { return x * 10 }
	isEven := func(x int) bool { return x%2 == 0 }

	fmt.Printf("list%v.Map(tenTimes).Filter(isEven) yields %v\n", list,
		list.
			Filter(isEven).
			Map(tenTimes))
}

// END OMIT
