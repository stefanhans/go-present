package main

import (
	"fmt"
)

type ListOfInt []int

// START OMIT
type ListFoldMonad struct {
	NeutralElement int
	AssocFunc      func(int, int) int
}

func (list ListOfInt) Fold(monad ListFoldMonad) int {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}
// END OMIT

func main() {
	var list = ListOfInt{-2, -1, 2, 2, 3}

	monad := ListFoldMonad{0, func(x, y int) int { return x + y }}

	fmt.Printf("List %v: Fold(monad) yields %v\n", list, list.Fold(monad))
}

// END OMIT