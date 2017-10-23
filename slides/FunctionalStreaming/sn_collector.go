package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	// START_1 OMIT
	node_1 := NewNodeOfInt()
	var i_1 int
	node_1.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 200)
		i_1++
		return in + i_1
	}

	node_2 := NewNodeOfInt()
	var i_2 int
	node_2.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 200)
		i_2++
		return in + i_2*10
	}

	node_3 := NewNodeOfInt()
	var i_3 int
	node_3.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 200)
		i_3++
		return in + i_3*100
	}
	// END_1 OMIT

	// START_2 OMIT
	node_out := NewNodeOfInt()

	node_1.Produce().Connect(node_out)
	node_2.Produce().Connect(node_out)
	node_3.Produce().Connect(node_out)

	node_out.Consume()
	time.Sleep(time.Second)
	// END_2 OMIT
}
