package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

type IntMonad struct {
	NeutralElement int
	AssocFunc      func(int, int) int
}

func (list ListOfInt) Fold(monad IntMonad) int {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}

// END OMIT

func main() {
	var list = ListOfInt{-2, -1, 2, 2, 3}

	monad := IntMonad{0, func(x, y int) int { return x + y }}

	fmt.Printf("List %v: Fold(monad) yields %v\n", list, list.Fold(monad))
}

// END OMIT
