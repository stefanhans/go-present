package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

type MapIntAggregates map[int]int

type IntMonad struct {
	NeutralElement int
	AssocFunc      func(int, int) int
}

func (list ListOfInt) Aggregate(monad IntMonad) MapIntAggregates {
	var out MapIntAggregates = make(map[int]int)

	for _, inK := range list {
		v, ok := out[inK]
		if !ok {
			out[inK] = monad.AssocFunc(monad.NeutralElement, inK)
		} else {
			out[inK] = monad.AssocFunc(v, inK)
		}
	}
	return out
}

// END OMIT

func main() {
	var list = ListOfInt{-2, -2, -1, -1, 0, 2, 3, 3, 10}

	// count
	monad := IntMonad{0, func(aggr, val int) int { return aggr + 1 }}
	// sum
	//monad := IntMonad{0, func(aggr, val int) int { return aggr + val }}

	fmt.Printf("List %v: Aggregate(monad) yields %v\n", list, list.Aggregate(monad))
}
