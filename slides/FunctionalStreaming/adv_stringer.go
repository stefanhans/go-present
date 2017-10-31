package main

import (
	"time"
	_ "fmt"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1, node_2 := NewNodeOfInt("node_1"), NewNodeOfInt("node_2") // node creation // HL
	var i int
	node_1.SetFunc(func(in int) int { i++; return in + i })
	node_2.SetFunc(func(in int) int { return in * 2 })

	node_1.Connect(node_2).Print() // stream configuration // HL
	node_1.ProduceAtMs(500) // sending data  // HL
	time.Sleep(time.Second)

	//fmt.Printf("node_1: %v\n", node_1)
	node_1.Report()
	time.Sleep(time.Second)
}