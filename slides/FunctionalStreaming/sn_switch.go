package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node := NewNodeOfInt()
	var i int
	node.SetFunc(func(in int) int { i++; return in + i })

	node_true, node_false := node.Switch(func(i int) bool { return i%2 == 0 }) // HL

	node_true.Printf("%v ")                                                                          // HL
	node_false.Map(func(i int) int { return i * 10 }).Printf("%v ")                             // HL

	node.ProduceAtMs(50)

	time.Sleep(time.Second)
}
