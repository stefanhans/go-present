package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()

	var i int
	node_1.SetFunc(func(in int) int {
		i++
		return in+i
	})

	node_1.ProduceAtMs(50).Calculate(func(in int) int { return in * 3 }).Print() // HL
	time.Sleep(time.Second)
}