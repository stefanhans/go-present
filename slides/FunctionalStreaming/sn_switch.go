package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	}

	node_2, node_3 := node_1.Produce().Switch(func(in int) bool { return (in%2 == 0) }) // HL
	node_2.Cf <- func(in int) int { return in * 10 }									// HL
	node_3.Cf <- func(in int) int { return in * 2 }										// HL

	node_2.Consume() 	// HL
	node_3.Consume() 	// HL
	time.Sleep(time.Second)
}
