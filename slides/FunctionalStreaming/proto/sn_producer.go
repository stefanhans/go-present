package main

import (
	"fmt"
	"time"
)

// START_1 OMIT
type NodeOfInt struct {
	in    chan int
	out   chan int
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case in := <-node.in:
				node.out <- in
			case <-node.close:
				fmt.Printf("\nnode closing...\n")
				return
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
	node.close = make(chan bool)
	node.Start()
	return &node
}

func (node *NodeOfInt) Stop() {
	node.close <- true
}
// END_3 OMIT

// START_4 OMIT
func main() {
	node := NewNodeOfInt()

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

// node.Start();time.Sleep(time.Second);node.Stop();time.Sleep(time.Second)
