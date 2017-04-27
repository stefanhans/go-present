package main

import (
	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
	"fmt"
)

func main() {
	var list = ListOfInt{-2, -1, 0, 2, 2, 3}

	var isEven ListFilterFunc = func(x int) bool { return x%2 == 0 }

	fmt.Printf("List %v: Filter(isEven) yields %v\n", list, list.Filter(isEven))
}