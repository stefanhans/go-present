package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START_1 OMIT
type NodeOfInt struct {
	in  chan int      // Input channel
	cin chan chan int // can be exchanged.

	f  func(int) int      // Function
	cf chan func(int) int // can be exchanged.

	out   chan int      // Output channel
	cout  chan chan int // can be exchanged.
	close chan bool     // OMIT
}

// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for {
			select {

			case in := <-node.in:
				node.out <- node.f(in) // Handle data (DEADLOCKS!) // HL

			case node.in = <-node.cin: // Change input channel
			case node.f = <-node.cf: // Change function
			case node.out = <-node.cout: // Change output channel
			case <-node.close:
				return // OMIT
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in } // Default function returns input value
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
	go func() {
		for {
			select {
			case in := <-node.out:
				fmt.Printf(format, in) // HL
			}
		}
	}()
}

// END_PRINTF OMIT

// START_4 OMIT
func (node *NodeOfInt) ProduceRandPositivAtMs(max int, ms time.Duration) *NodeOfInt {
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			select {
			default:
				node.in <- rand.Intn(max) + 1
			} // HL
			time.Sleep(time.Millisecond * ms) // HL
		}
	}()
	return node
}

// END_4 OMIT

// START_AggregatorOfInt_1 OMIT
type AggregatorOfInt struct {
	in              chan int
	cin             chan chan int
	aggregator_map  *map[int]int            // HL
	caggregator_map chan *map[int]int       // HL
	aggregate       func(int, *map[int]int) // HL
	close           chan bool               // OMIT
}

// END_AggregatorOfInt_1 OMIT

// START_AggregatorOfInt_2 OMIT
func (aggregator *AggregatorOfInt) Start() {
	go func() {
		for {
			select {
			case in := <-aggregator.in: // HL
				aggregator.aggregate(in, aggregator.aggregator_map) // HL
			case aggregator.in = <-aggregator.cin: // HL
			case aggregator_map := <-aggregator.caggregator_map: // HL
				aggregator.caggregator_map <- aggregator.aggregator_map // HL
				aggregator.aggregator_map = aggregator_map              // HL
			case <-aggregator.close:
				return // OMIT
			}
		}
	}()
}

// END_AggregatorOfInt_2 OMIT

// START_AggregatorOfInt_3 OMIT
func NewAggregatorOfInt(aggr_map *map[int]int, f func(int, *map[int]int)) *AggregatorOfInt { // HL
	aggregator := AggregatorOfInt{}
	aggregator.in = make(chan int)
	aggregator.cin = make(chan chan int)
	aggregator.aggregator_map = aggr_map                 // HL
	aggregator.caggregator_map = make(chan *map[int]int) // HL
	aggregator.aggregate = f                             // HL
	aggregator.close = make(chan bool)                   // OMIT
	aggregator.Start()
	return &aggregator
}

// END_AggregatorOfInt_3 OMIT

// START_RESET OMIT
func (aggregator *AggregatorOfInt) Reset(aggr_map *map[int]int) *map[int]int {
	aggregator.caggregator_map <- aggr_map
	return <-aggregator.caggregator_map
}

// END_RESET OMIT

func (node *NodeOfInt) ConnectAggregator(nextNode *AggregatorOfInt) *AggregatorOfInt {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node_in := NewNodeOfInt()

	aggr_map_1 := make(map[int]int)
	aggr_func := func(i int, aggr_map *map[int]int) { (*aggr_map)[i] = (*aggr_map)[i] + 1 }
	aggregator := NewAggregatorOfInt(&aggr_map_1, aggr_func)

	node_in.ConnectAggregator(aggregator)
	node_in.ProduceRandPositivAtMs(4, 10)

	time.Sleep(time.Second)
	fmt.Printf("%v\n", aggr_map_1)
	time.Sleep(time.Second)

	aggr_map_2 := make(map[int]int)
	aggr_map_1 = (*aggregator.Reset(&aggr_map_2))
	time.Sleep(time.Second)

	fmt.Printf("%v\n", aggr_map_1)
	fmt.Printf("%v\n", aggr_map_2)
}
