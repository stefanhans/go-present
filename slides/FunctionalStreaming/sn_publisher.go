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
	publisher := NewPublisherOfInt()
	subscriber_1 := NewNodeOfInt()
	subscriber_2 := NewNodeOfInt()
	subscriber_3 := NewNodeOfInt()

	node_1.ProduceAtMs(200).ConnectPublisher(publisher)
	publisher.SubscribePublisher("1st", subscriber_1)
	publisher.SubscribePublisher("2nd", subscriber_2)
	publisher.SubscribePublisher("3rd", subscriber_3)

	subscriber_1.Print()
	subscriber_2.Calculate(func(i int) int { return i * 10 }).Print()
	subscriber_3.Calculate(func(i int) int { return i * 100 }).Print()
	time.Sleep(time.Second)

	fmt.Println()

	publisher.UnsubscribePublisher("2nd")
	time.Sleep(time.Second)
	// END_6 OMIT
}
