package main

import (
	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
	"fmt"
)

func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}

	tenTimes := func(x int) int { return x * 10 }

	fmt.Printf("List %v: Map(tenTimes) yields %v\n", list, list.Map(tenTimes))
}