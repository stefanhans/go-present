package functionalstreams

import (
	"fmt"
)

// START_1 OMIT
type NodeOfInt struct {
	in    chan int
	out   chan int
	f     func(int) int
	cin   chan chan int		// HL
	cout  chan chan int		// HL
	cf    chan func(int) int // HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for { select {
		case in := <-node.in: node.out <- node.f(in) // HL
		case node.in = <-node.cin:		// HL
		case node.out = <-node.cout:	// HL
		case node.f = <-node.cf: // HL
		case <-node.close: return
		}} }()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.out = make(chan int)
	node.f = func(in int) int { return in }
	node.cin = make(chan chan int)
	node.cout = make(chan chan int)
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
			select { default: node.in <- 0 }	// HL
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
		for { select {
			case in := <-node.out: 				// HL
				fmt.Printf("%v ", in)		// HL
	}}}()
}
// END_6 OMIT

// START_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) {
	node.cf <- f
}
// END_SETFUNC OMIT

// START_CALC OMIT
// node(f) -> node
func (node *NodeOfInt) Calculate(calc func(int) int) *NodeOfInt {
	nextNode := NewNodeOfInt()
	nextNode.cf <- calc

	node.Connect(nextNode)
	return nextNode
}
// END_CALC OMIT

