package main

import (
	"fmt"
	"time"
)

// START_1 OMIT
type NodeOfInt struct {
	in    chan int
	out   chan int
	f     func(int) int			// HL
	cf    chan func(int) int	// HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case in := <-node.in: 		// HL
				node.out <- node.f(in)	// HL
			case node.f = <-node.cf:	// HL
			case <-node.close: fmt.Printf("node closing...\n"); return
			}
		}
	}()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.out = make(chan int)
	node.f = func(in int) int { return in }	// HL
	node.cf = make(chan func(int) int)		// HL
	node.close = make(chan bool)
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfInt) Stop() {
	node.close <- true
}

// START_4 OMIT
func main() {
	node := NewNodeOfInt()

	node.cf <- func(in int) int { return in * 2 }	// HL

	go func() {
		var in int
		for {
			time.Sleep(time.Millisecond * 50)
			in++; node.in <- in
		}
	}()
	go func() {
		for {
			fmt.Printf("%v ", <-node.out)
		}
	}()

	time.Sleep(time.Second)
	node.Stop()
}
// END_4 OMIT
