package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
	"fmt"
)

func main() {
	node := NewNodeOfInt()
	var i int; node.SetFunc(func(in int) int { i++; return in + i })

	distributor := NewDistributorOfInt()
	distributor.DistributeRoundRobin()	                 // Distribute round-robin // HL

	subscriber_1, subscriber_2, subscriber_3 := NewNodeOfInt(), NewNodeOfInt(), NewNodeOfInt()
	subscriber_1.Printf("%v ")
	subscriber_2.Map(func(i int) int { return i * 10 }).Printf("%v ")
	subscriber_3.Map(func(i int) int { return i * 100 }).Printf("%v ")

	node.ConnectDistributor(distributor)
	distributor.SubscribeDistributor("1st", subscriber_1)
	distributor.SubscribeDistributor("2nd", subscriber_2)
	distributor.SubscribeDistributor("3rd", subscriber_3)

	node.ProduceAtMs(100)
	time.Sleep(time.Second)
	fmt.Println()
	distributor.DistributeToAll()                          // Distribute to all	// HL
	time.Sleep(time.Second)
}
