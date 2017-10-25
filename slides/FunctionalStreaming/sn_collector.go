package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	// START_1 OMIT
	node_1 := NewNodeOfInt()
	var i_1 int
	node_1.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 200)
		i_1++
		return in + i_1
	})

	node_2 := NewNodeOfInt()
	var i_2 int
	node_2.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 200)
		i_2++
		return in + i_2*10
	})

	node_3 := NewNodeOfInt()
	var i_3 int
	node_3.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 200)
		i_3++
		return in + i_3*100
	})
	// END_1 OMIT

	// START_2 OMIT
	node_out := NewNodeOfInt()

	node_1.Produce().Connect(node_out)	// 1, 2, 3, 4, ...
	node_2.Produce().Connect(node_out)	// 10, 20, 30, 40, ...
	node_3.Produce().Connect(node_out)	// 100, 200, 300, 400, ...

	node_out.Print()
	time.Sleep(time.Second)
	// END_2 OMIT
}
