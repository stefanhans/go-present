package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_in := NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	aggregator := NewAggregatorOfInt()

	node_in.ConnectAggregator(aggregator).Print()
	node_in.ProduceAtMs(200)

	time.Sleep(time.Second)
}
