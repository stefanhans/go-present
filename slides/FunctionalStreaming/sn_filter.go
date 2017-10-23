package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

// START_6 OMIT
func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	}

	node_2 := NewNodeOfInt()
	node_2.Cf <- func(in int) int { return in * 10 }

	node_1.Produce().Connect(node_2).Filter(func(in int) bool { return in >= 100 }).Consume()
	time.Sleep(time.Second)

}
// END_6 OMIT
