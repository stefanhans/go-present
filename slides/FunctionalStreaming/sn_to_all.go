package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node := NewNodeOfInt()
	var i int
	node.SetFunc(func(in int) int { i++; return in + i })

	publisher := NewPublisherOfInt()	// HL
	publisher.DistributeRoundRobin()	// HL
	node.ProduceAtMs(50).ConnectPublisher(publisher) // HL

	subscriber_1, subscriber_2, subscriber_3 := NewNodeOfInt(), NewNodeOfInt(), NewNodeOfInt()	// HL
	publisher.SubscribePublisher("1st", subscriber_1)	// HL
	publisher.SubscribePublisher("2nd", subscriber_2)	// HL
	publisher.SubscribePublisher("3rd", subscriber_3)	// HL

	subscriber_1.Print()	// HL
	subscriber_2.Calculate(func(i int) int { return i * 10 }).Print()	// HL
	subscriber_3.Calculate(func(i int) int { return i * 100 }).Print()	// HL

	time.Sleep(time.Second)
}
