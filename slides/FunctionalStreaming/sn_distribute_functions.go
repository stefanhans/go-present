package main

import (
	"time"
	"fmt"
)
// START_1 OMIT
type NodeOfInt struct {
	in    chan int                                 // Input channel
	cin   chan chan int                            // can be exchanged.

	f     func(int) int                            // Function
	cf    chan func(int) int                       // can be exchanged.

	out   chan int                                 // Output channel
	cout  chan chan int                            // can be exchanged.
	close chan bool // OMIT
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for { select {

		case in := <-node.in: node.out <- node.f(in) // Handle data (DEADLOCKS!) // HL

		case node.in = <-node.cin:   	            // Change input channel
		case node.f = <-node.cf: 		            // Change function
		case node.out = <-node.cout: 	            // Change output channel
		case <-node.close: return // OMIT
		}
		}}()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in }        // Default function returns input value
	node.cf = make(chan func(int) int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.close = make(chan bool) // OMIT
	node.Start()
	return &node
}
// END_3 OMIT

// START_5 OMIT
func (node *NodeOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) { node.cf <- f }
// END_SETFUNC OMIT

// START_PRINTF OMIT
func (node *NodeOfInt) Printf(format string) {
	go func() { for { select {
	case in := <-node.out: fmt.Printf(format, in)		// HL
	}}}()
}
// END_PRINTF OMIT


// START_3 OMIT
func (node *NodeOfInt) ProduceAtMs(n time.Duration) *NodeOfInt {
	go func() { for { select {
	default: node.in <- 0 }	               // Trigger permanently // HL
		time.Sleep(time.Millisecond * n)	      // with delay in ms // HL
	}}()
	return node
}
// END_3 OMIT

// START_TEE OMIT
func (node *NodeOfInt) Tee() (*NodeOfInt, *NodeOfInt) {
	distributor := NewDistributorOfInt()
	node.ConnectDistributor(distributor)

	nodeA := NewNodeOfInt()
	nodeB := NewNodeOfInt()

	distributor.SubscribeDistributor("A", nodeA)
	distributor.SubscribeDistributor("B", nodeB)

	return nodeA, nodeB
}
// END_TEE OMIT

// START_SUBSCRIPTION OMIT
type SubscriptionToInt struct {
	name string
	cint chan int
}

// END_SUBSCRIPTION OMIT

// START_1 OMIT
type DistributorOfInt struct {
	in                   chan int
	cin                  chan chan int
	f                    func(int)				 // Distribute over subscriptions // HL
	cf                   chan func(int)			// // HL
	out_map              map[string]chan int       // Handle subscriptions // HL
	cout_map_subscribe   chan SubscriptionToInt    // // HL
	cout_map_unsubscribe chan string               // // HL
	out_index            []SubscriptionToInt	   // Subscriptions ordered by number // HL
	last_index           int                       // // HL
	clast_index          chan int 	// OMIT
	close                chan bool 	// OMIT
}

// END_1 OMIT

// START_2 OMIT
func (distributor *DistributorOfInt) Start() {
	go func() {
		for { select {
		case distributor.in = <-distributor.cin:

			// Function distributes the input value // HL
		case in := <-distributor.in: distributor.f(in)

		case distributor.f = <-distributor.cf:

			// Subscribe to the distributor // HL
		case subscription := <-distributor.cout_map_subscribe:
			distributor.out_map[subscription.name] = subscription.cint
			distributor.out_index = append(distributor.out_index, subscription)

			// Unsubscribe from the distributor // HL
		case name := <-distributor.cout_map_unsubscribe:
			delete(distributor.out_map, name)

			// delete from distributor.out_index accordingly
			// (not shown for brevity) ...
			i := -1; _ = i 	// OMIT
			for n, subscription := range distributor.out_index { 	// OMIT
				if subscription.name == name { i = n }} 	// OMIT
			if i != -1 { 	// OMIT
				distributor.out_index = append(distributor.out_index[:i], 	// OMIT
					distributor.out_index[i+1:]...)} 	// OMIT
		case <-distributor.close: return 	// OMIT
		}}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewDistributorOfInt() *DistributorOfInt {
	distributor := DistributorOfInt{}
	distributor.in = make(chan int)
	distributor.cin = make(chan chan int)

	distributor.out_map = make(map[string]chan int)
	distributor.cout_map_subscribe = make(chan SubscriptionToInt)
	distributor.cout_map_unsubscribe = make(chan string)

	distributor.out_index = make([]SubscriptionToInt, 0)

	// Each subsription gets a goroutine for sending to its channel  // HL
	distributor.f = func(in int) {
		for _, cout := range distributor.out_map {
			go func(cout chan int, in int) { cout <- in }(cout, in)
		}
	}
	distributor.cf = make(chan func(int))
	distributor.close = make(chan bool) // OMIT
	distributor.Start()
	return &distributor
}

// END_3 OMIT

// START_CONNECTD OMIT
func (node *NodeOfInt) ConnectDistributor(distributor *DistributorOfInt) *DistributorOfInt {
	node.cout <- distributor.in
	return distributor
}
// END_CONNECTD OMIT

// START_SUBSCRIBE OMIT
func (distributor *DistributorOfInt) SubscribeDistributor(name string, nextNode *NodeOfInt) {
	distributor.cout_map_subscribe <- SubscriptionToInt{name, nextNode.in}
}

func (distributor *DistributorOfInt) UnsubscribeDistributor(name string) {
	distributor.cout_map_unsubscribe <- name
}
// END_SUBSCRIBE OMIT


// START_MAP OMIT
func (node *NodeOfInt) Map(f func(int) int) *NodeOfInt {
	nextNode := NewNodeOfInt()
	nextNode.cf <- f

	node.Connect(nextNode)
	return nextNode
}
// END_MAP OMIT


// START_RR OMIT
func (distributor *DistributorOfInt) DistributeRoundRobin() {
	distributor.cf <- func(in int) {
		if distributor.last_index == len(distributor.out_map) {        // Reset index
			distributor.last_index = 0
		}
		for i, subscription := range distributor.out_index {           // Loop until last index
			if i == distributor.last_index {
				go func(cout chan int, in int) { cout <- in }(subscription.cint, in)
				distributor.last_index++
				return
			}}}
}
// END_RR OMIT

// START_TOALL OMIT
func (distributor *DistributorOfInt) DistributeToAll() {
	distributor.cf <- func(in int) {
		for _, cout := range distributor.out_map {
			go func(cout chan int, in int) { cout <- in }(cout, in)
		}}
}
// END_TOALL OMIT

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
