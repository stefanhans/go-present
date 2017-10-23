package main

import (
	"fmt"
	"time"
)

// START_1 OMIT
type NodeOfFloat struct {
	in    chan float64
	out   chan float64
	f     func(float64) float64			// HL
	cf    chan func(float64) float64	// HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfFloat) Start() {
	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case fl := <-node.in: 		// HL
				node.out <- node.f(fl)	// HL
			case node.f = <-node.cf:	// HL
			case <-node.close: fmt.Printf("node closing...\n"); return
			}
		}
	}()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfFloat() *NodeOfFloat {
	node := NodeOfFloat{}
	node.in = make(chan float64)
	node.out = make(chan float64)
	node.f = func(fl float64) float64 { return fl }	// HL
	node.cf = make(chan func(float64) float64)		// HL
	node.close = make(chan bool)
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfFloat) Stop() {
	node.close <- true
}

// START_4 OMIT
func main() {
	node := NewNodeOfFloat()

	node.cf <- func(fl float64) float64 { return fl * 2.0 }	// HL

	go func() {
		var fl float64
		for {
			time.Sleep(time.Millisecond * 50)
			fl++; node.in <- fl
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
