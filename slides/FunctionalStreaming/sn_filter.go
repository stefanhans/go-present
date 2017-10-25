package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

// START_6 OMIT
func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	})

	node_1.Produce().Filter(func(in int) bool { return in%2 == 1 }).Consume() // HL
	time.Sleep(time.Second)

}
// END_6 OMIT
