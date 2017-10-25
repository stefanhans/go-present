package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()

	var i int
	node_1.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in+i
	})

	node_1.Produce().Calculate(func(in int) int { return in * 3 }).Print() // HL
	time.Sleep(time.Second)
}