package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_in, node_out := NewNodeOfInt(), NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	buffer := NewBufferOfInt()

	node_in.ConnectBuffer(buffer).Connect(node_out).Print()
	node_in.ProduceAtMs(200)

	time.Sleep(time.Second) // Not buffering by default
	buffer.Buffer()
	time.Sleep(time.Second)
	buffer.Flush()
	time.Sleep(time.Second)
}
