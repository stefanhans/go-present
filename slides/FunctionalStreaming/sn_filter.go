package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

// START_6 OMIT
func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.SetFunc(func(in int) int { i++; return in + i })

	node_1.Filter(func(in int) bool { return in%2 == 1 }).Printf("%v ") // HL

	node_1.ProduceAtMs(50)

	time.Sleep(time.Second)

}
// END_6 OMIT
