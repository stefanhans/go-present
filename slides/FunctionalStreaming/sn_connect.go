package main

import (
	"time"
	"fmt"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1, node_2 := NewNodeOfInt(), NewNodeOfInt() // node creation // HL
	var i int
	node_1.SetFunc(func(in int) int { i++; return in + i })
	node_2.SetFunc(func(in int) int { return in * 2 })

	node_1.Connect(node_2).Print() // stream configuration // HL
	node_1.ProduceAtMs(50) // sending data  // HL
	time.Sleep(time.Second)

	fmt.Println()
	node_2.SetFunc(func(in int) int { return in * 10 }) // change function // HL
	time.Sleep(time.Second)
}