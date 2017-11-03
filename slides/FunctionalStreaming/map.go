package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()                                     // node creation // HL
	var i int                                                    //
	node_1.SetFunc(func(in int) int { i++; return in+i })        //

	node_1.Map(func(in int) int { return in * 3 }).Printf("%v ") // stream configuration // HL

	node_1.ProduceAtMs(50)                                       // sending data // HL
	time.Sleep(time.Second)
}