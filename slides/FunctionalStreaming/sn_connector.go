package main

import (
"fmt"
"time"
)

// START_1 OMIT
type NodeOfFloat struct {
	in    chan float64
	out   chan float64
	f     func(float64) float64
	cin   chan chan float64		// HL
	cout  chan chan float64		// HL
	cf    chan func(float64) float64
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfFloat) Start() {
	go func() {
		fmt.Printf("node starting...\n")
		for { select {
			case node.in = <-node.cin:		// HL
			case node.out = <-node.cout:	// HL
			case fl := <-node.in: node.out <- node.f(fl)
			case node.f = <-node.cf:
			case <-node.close: fmt.Printf("node closing...\n"); return
		}} }()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfFloat() *NodeOfFloat {
	node := NodeOfFloat{}
	node.in = make(chan float64)
	node.out = make(chan float64)
	node.f = func(fl float64) float64 { return fl }
	node.cin = make(chan chan float64)			// HL
	node.cout = make(chan chan float64) 		// HL
	node.cf = make(chan func(float64) float64)
	node.close = make(chan bool)
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfFloat) Stop() {
	node.close <- true
}

// START_4 OMIT
func (node *NodeOfFloat) Produce() *NodeOfFloat {
	go func() {
		for {
			select { default: node.in<- 0.0 }	// HL
		}}()
	return node
}
// END_4 OMIT

// START_5 OMIT
func (node *NodeOfFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_6 OMIT
func (node *NodeOfFloat) Consume() {
	go func() {
		for {
			select {
			case fl := <-node.out: 				// HL
				fmt.Printf("%v ", fl)		// HL
			}
		}
	}()
}
// END_6 OMIT

// START_7 OMIT
func main() {
	node_1 := NewNodeOfFloat()
	node_2 := NewNodeOfFloat()

	var i float64
	node_1.cf <- func(fl float64) float64 {	// HL
		time.Sleep(time.Millisecond * 50)	// HL
		i++								    // HL
		return fl+i							// HL
	}
	node_2.cf <- func(fl float64) float64 { return fl * 2.0 } // HL

	node_1.Produce().Connect(node_2).Consume() // HL
	time.Sleep(time.Second)
}
// END_7 OMIT
