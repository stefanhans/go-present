package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1, node_2 := NewNodeOfInt("node_1"), NewNodeOfInt("node_2")

	var i int
	node_1.SetFunc(func(in int) int { i++; return in + i })
	node_1.SetDescription("var i int; func(in int) int { i++; return in + i }")

	node_2.SetFunc(func(in int) int { return in * 2 })
	node_2.SetDescription("func(in int) int { i++; return in * 2 }")

	node_1.Connect(node_2).Calculate(func(i int) int { return i }, "{ return i }").Print()

	node_1.Report()

	time.Sleep(time.Second)
}