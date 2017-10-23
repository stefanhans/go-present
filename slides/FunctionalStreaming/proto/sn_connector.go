package main

import (
	"fmt"
	"time"
)

// START_1 OMIT
type NodeOfInt struct {
	in    chan int
	out   chan int
	f     func(int) int
	cin   chan chan int		// HL
	cout  chan chan int		// HL
	cf    chan func(int) int
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		fmt.Printf("node starting...\n")
		for { select {
			case node.in = <-node.cin:		// HL
			case node.out = <-node.cout:	// HL
			case in := <-node.in: node.out <- node.f(in)
			case node.f = <-node.cf:
			case <-node.close: fmt.Printf("node closing...\n"); return
		}} }()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.out = make(chan int)
	node.f = func(in int) int { return in }
	node.cin = make(chan chan int)			// HL
	node.cout = make(chan chan int) 		// HL
	node.cf = make(chan func(int) int)
	node.close = make(chan bool)
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfInt) Stop() {
	node.close <- true
}

// START_4 OMIT
func (node *NodeOfInt) Produce() *NodeOfInt {
	go func() {
		for {
			select { default: node.in<- 0.0 }	// HL
		}}()
	return node
}
// END_4 OMIT

// START_5 OMIT
func (node *NodeOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_6 OMIT
func (node *NodeOfInt) Consume() {
	go func() {
		for {
			select {
			case in := <-node.out: 				// HL
				fmt.Printf("%v ", in)		// HL
			}
		}
	}()
}
// END_6 OMIT

// START_7 OMIT
func main() {
	node_1 := NewNodeOfInt()
	node_2 := NewNodeOfInt()

	var i int
	node_1.cf <- func(in int) int {	// HL
		time.Sleep(time.Millisecond * 50)	// HL
		i++								    // HL
		return in+i							// HL
	}
	node_2.cf <- func(in int) int { return in * 2 } // HL

	node_1.Produce().Connect(node_2).Consume() // HL
	time.Sleep(time.Second)
}
// END_7 OMIT
