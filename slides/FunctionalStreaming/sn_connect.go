package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()
	node_2 := NewNodeOfInt()

	var i int // HL
	node_1.SetFunc(func(in int) int {	// HL
		i++; return in+i				// HL
	}) // HL

	node_2.SetFunc(func(in int) int { return in * 2 }) // HL

	node_1.ProduceAtMs(50).Connect(node_2).Print() // HL

	time.Sleep(time.Second)
}