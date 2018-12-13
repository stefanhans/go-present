package main

import (
	"fmt"
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

// START_NODE_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) { node.cf <- f }

// END_NODE_SETFUNC OMIT

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

// START_ProduceAtMs OMIT
func (node *NodeOfInt) ProduceAtMs(n time.Duration) *NodeOfInt {
	go func() {
		for {
			select {
			default:
				node.in <- 0
			} // Trigger permanently // HL
			time.Sleep(time.Millisecond * n) // with delay in ms // HL
		}
	}()
	return node
}

// END_ProduceAtMs OMIT

func main() {
	node_1, node_2 := NewNodeOfInt(), NewNodeOfInt()        // nodes' creation // HL
	var i int                                               //
	node_1.SetFunc(func(in int) int { i++; return in + i }) //
	node_2.SetFunc(func(in int) int { return in * 2 })      //

	node_1.Connect(node_2).Printf("%v ") // stream configuration // HL
	node_1.ProduceAtMs(50)               // sending data  // HL
	time.Sleep(time.Second)
	fmt.Println()

	node_2.SetFunc(func(in int) int { return in * 10 }) // change function // HL
	time.Sleep(time.Second)
}
