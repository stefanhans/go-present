package main

import (
	"fmt"
)

// START OMIT
type ListOfInt []int

type MapIntAggregates map[int]int

type intAggregateFunc func(aggrValue, listValue int) (aggregateValue int)

func (list ListOfInt) Aggregate(f intAggregateFunc) MapIntAggregates {
	var out MapIntAggregates = make(map[int]int)

	for _, inK := range list {
		out[inK] = f(out[inK], inK)
	}
	return out
}
// END OMIT

func main() {
	var list = ListOfInt{-2, -2, -1, -1, 0, 2, 3, 3, 10}

	var sum intAggregateFunc = func(inAggr, inV int) int {
		return inAggr + inV
	}
	fmt.Printf("list%v: \nGroup(count) returns list%v\n",
		list, list.Aggregate(sum))
}
