package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_counter, node_TenTimes := NewNodeOfInt("node_counter"), NewNodeOfInt("node_TenTimes")
	var i int; node_counter.SetFunc(func(in int) int { i++; return in + i })

	node_TenTimes.SetFunc(func(in int) int { return in * 10 })
	node_TenTimes.Printf("%v ")

	node_HundredTimes := NewNodeOfInt("node_HundredTimes")
	node_HundredTimes.SetFunc(func(in int) int { return in * 11 })
	node_HundredTimes.Printf("%v ")

	node_counter.Connect(node_TenTimes)
	node_counter.ProduceAtMs(200)							// START: counter * 10 // HL
	time.Sleep(time.Second)
	node_HundredTimes.AddTee(node_counter, node_TenTimes)	// TEE: counter * 11 // HL
	time.Sleep(time.Second)
}