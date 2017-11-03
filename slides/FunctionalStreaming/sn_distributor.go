package main

import (
	"fmt"
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.SetFunc(func(in int) int {
		i++
		return in + i
	})

	// START_6 OMIT
	distributor := NewDistributorOfInt()                                // nodes' creation
	subscriber_1 := NewNodeOfInt()                                      //
	subscriber_2 := NewNodeOfInt()                                      //
	subscriber_3 := NewNodeOfInt()                                      //
	                                                                    //
	subscriber_1.Printf("%v ")                                          //
	subscriber_2.Map(func(i int) int { return i * 10 }).Printf("%v ")   //
	subscriber_3.Map(func(i int) int { return i * 100 }).Printf("%v ")  //


	node_1.ConnectDistributor(distributor)                          // stream configuration
	distributor.SubscribeDistributor("1st", subscriber_1)           //
	distributor.SubscribeDistributor("2nd", subscriber_2)           //
	distributor.SubscribeDistributor("3rd", subscriber_3)           //

	node_1.ProduceAtMs(200)                                         // sending data

	time.Sleep(time.Second)
	fmt.Println()

	distributor.UnsubscribeDistributor("2nd")                       // unsubscribe

	time.Sleep(time.Second)
	// END_6 OMIT
}
