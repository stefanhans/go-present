package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_in := NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	converter := NewConverterIntToFloat()
	node_out := NewNodeOfFloat()
	node_in.ConnectConverterIntToFloat(converter).Connect(node_out).Print()


	node_in.ProduceAtMs(200)
	time.Sleep(time.Second)



}
