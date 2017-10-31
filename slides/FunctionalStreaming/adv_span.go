package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_counter := NewNodeOfInt("node_counter")
	var i int; node_counter.SetFunc(func(in int) int { i++; return in + i })
	node_TenTimes := NewNodeOfInt("node_TenTimes")
	node_TenTimes.SetFunc(func(in int) int { return in * 10 })
	node_HundredTimes := NewNodeOfInt("node_HundredTimes")
	node_HundredTimes.SetFunc(func(in int) int { return in * 100 })
	node_HundredTimes.Printf("%v ")

	node_TwoTimes := NewNodeOfInt("node_TwoTimes")
	node_TwoTimes.SetFunc(func(in int) int { return in * 2 })
	node_TwoTimes.Printf("%v ")

	node_counter.Connect(node_TenTimes).Connect(node_HundredTimes)
	node_counter.ProduceAtMs(200)					// START: counter * 10 * 100 // HL
	time.Sleep(time.Second)
	node_counter.Connect(node_TwoTimes)				// SPAN: counter [* 10 * 100] * 2 // HL
	time.Sleep(time.Second)
}