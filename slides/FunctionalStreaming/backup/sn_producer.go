package main

import (
	"fmt"
	"time"
)

// START_1 OMIT
type NodeOfFloat struct {
	in    chan float64
	out   chan float64
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfFloat) Start() {
	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case fl := <-node.in:
				node.out <- fl
			case <-node.close:
				fmt.Printf("\nnode closing...\n")
				return
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
	node.close = make(chan bool)
	node.Start()
	return &node
}

func (node *NodeOfFloat) Stop() {
	node.close <- true
}
// END_3 OMIT

// START_4 OMIT
func main() {
	node := NewNodeOfFloat()

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

// node.Start();time.Sleep(time.Second);node.Stop();time.Sleep(time.Second)
