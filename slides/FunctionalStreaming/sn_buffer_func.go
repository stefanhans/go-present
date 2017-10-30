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

	// START_BUFFUNC OMIT
	node_in.ConnectBuffer(buffer).Connect(node_out).Print()
	node_in.ProduceAtMs(200)

	buffer.SetFunc(func(i int) bool {
		return buffer.Len()%5 == 0
	})
	time.Sleep(time.Second * 5)
	// END_BUFFUNC OMIT


}