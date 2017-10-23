package main

import (
	"fmt"
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	}

	// START_6 OMIT
	publisher := NewPublisherOfInt()
	subscriber_1 := NewNodeOfInt()
	subscriber_2 := NewNodeOfInt()
	subscriber_3 := NewNodeOfInt()
	node_1.Produce().ConnectPublisher(publisher) // HL
	publisher.SubscribePublisher("1st", subscriber_1)
	publisher.SubscribePublisher("2nd", subscriber_2)
	publisher.SubscribePublisher("3rd", subscriber_3)
	subscriber_1.Consume()
	subscriber_2.Calculate(func(i int) int { return i * 10 }).Consume()
	subscriber_3.Calculate(func(i int) int { return i * 100 }).Consume()
	time.Sleep(time.Second)
	fmt.Println()
	publisher.UnsubscribePublisher("2nd")
	time.Sleep(time.Second)
	// END_6 OMIT
}
