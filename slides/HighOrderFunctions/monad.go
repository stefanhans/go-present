package main

import (
	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
	"fmt"
)

func main() {
	var list = ListOfInt{-2, -1, 2, 2, 3}

	monad := ListFoldMonad{0, func(x, y int) int { return x + y }}

	fmt.Printf("List %v: Fold(monad) yields %v\n", list, list.Fold(monad))
}
