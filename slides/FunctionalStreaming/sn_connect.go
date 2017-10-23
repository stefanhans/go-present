package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()
	node_2 := NewNodeOfInt()

	var i int
	node_1.Cf <- func(in int) int {	// HL
		time.Sleep(time.Millisecond * 50)	// HL
		i++								    // HL
		return in+i							// HL
	}
	node_2.Cf <- func(in int) int { return in * 2 } // HL

	node_1.Produce().Connect(node_2).Consume() // HL
	time.Sleep(time.Second)
}