package main

import (
	"fmt"
	"time"
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

// START_1 OMIT
type ConverterIntToFloat struct {
	in      chan int
	cin     chan chan int
	convert func(int) float64 // HL
	out     chan float64
	cout    chan chan float64
	close   chan bool // OMIT
}
// END_1 OMIT

// START_2 OMIT
func (converter *ConverterIntToFloat) Start() {
	go func() {
		for {
			select {
			case converter.in = <-converter.cin:
			case in := <-converter.in: converter.out <- converter.convert(in) // HL
			case converter.out = <-converter.cout:
			case <-converter.close: return // OMIT
			}
		}
	}()
}
// END_2 OMIT

// START_3 OMIT
func NewConverterIntToFloat() *ConverterIntToFloat {
	converter := ConverterIntToFloat{}
	converter.in = make(chan int)
	converter.cin = make(chan chan int)
	converter.convert = func(in int) float64 { return float64(in) } // HL
	converter.out = make(chan float64)
	converter.cout = make(chan chan float64)
	converter.close = make(chan bool) // OMIT
	converter.Start()
	return &converter
}
// END_3 OMIT

func (converter *ConverterIntToFloat) Stop() {
	converter.close <- true
}

// START_CONNECT OMIT
func (converter *ConverterIntToFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	converter.cout <- nextNode.in
	return nextNode
}
// END_CONNECT OMIT

// START_1 OMIT
type NodeOfFloat struct {
	in    chan float64
	out   chan float64
	f     func(float64) float64
	cin   chan chan float64		// HL
	cout  chan chan float64		// HL
	cf    chan func(float64) float64 // HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfFloat) Start() {
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
func NewNodeOfFloat() *NodeOfFloat {
	node := NodeOfFloat{}
	node.in = make(chan float64)
	node.out = make(chan float64)
	node.f = func(in float64) float64 { return in }
	node.cin = make(chan chan float64)
	node.cout = make(chan chan float64)
	node.cf = make(chan func(float64) float64)
	node.close = make(chan bool)
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfFloat) Stop() {
	node.close <- true
}

// START_5 OMIT
func (node *NodeOfFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_CONNECTCONVERTER OMIT
func (node *NodeOfInt) ConnectConverterIntToFloat(nextNode *ConverterIntToFloat) *ConverterIntToFloat { // HL
	node.cout <- nextNode.in
	return nextNode
}
// END_CONNECTCONVERTER OMIT

func (node *NodeOfFloat) Printf(format string) {
	go func() {
		for {
			select {
			case in := <-node.out: // HL
				fmt.Printf(format, in) // HL
			}
		}
	}()
}


func main() {
	node_in := NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	converter := NewConverterIntToFloat()
	node_out := NewNodeOfFloat()

	node_in.ConnectConverterIntToFloat(converter).Connect(node_out).Printf("%f ")

	node_in.ProduceAtMs(200)

	time.Sleep(time.Second)
}
