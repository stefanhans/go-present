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

	tee := NewTeeOfInt()
	node_a := NewNodeOfInt()
	node_b := NewNodeOfInt()
	node_b.Cf <- func(in int) int { return in * -2 }

	node_1.Produce().ConnectTee(tee).ConnectNodes(node_a, node_b)
	node_a.Consume()
	node_b.Consume()
	time.Sleep(time.Second)
}

