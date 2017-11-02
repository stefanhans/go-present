package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node := NewNodeOfInt()
	var i int
	node.SetFunc(func(in int) int { i++; return in + i })

	node_left, node_right := node.Tee()         	// HL

	node_left.Print()                                               // HL
	node_right.Map(func(i int) int { return i * 10 }).Printf("%v ") // HL

	node.ProduceAtMs(100)

	time.Sleep(time.Second)
}
