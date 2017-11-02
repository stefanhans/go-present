package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1, node_2, node_3 := NewNodeOfInt(), NewNodeOfInt(), NewNodeOfInt()

	var i, j, k int
	node_1.SetFunc(func(in int) int { i++; return in + i })
	node_2.SetFunc(func(in int) int { j++; return in + j * 10 })
	node_3.SetFunc(func(in int) int { k++; return in + k * 100})

	node_out := NewNodeOfInt()
	node_out.Print()

	node_1.Connect(node_out)
	node_2.Connect(node_out)
	node_3.Connect(node_out)

	node_1.ProduceAtMs(200)
	node_2.ProduceAtMs(200)
	node_3.ProduceAtMs(200)

	time.Sleep(time.Second)
}
