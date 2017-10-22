package main

import (
	"fmt"
	"time"
)

type NodeOfFloat struct {
	in    chan float64
	cin   chan chan float64
	out   chan float64
	cout  chan chan float64
	f     func(float64) float64
	cf    chan func(float64) float64
	running  bool
	close chan bool
}

func (node *NodeOfFloat) String() string {
	return fmt.Sprintf("NodeOfFloat: %v\n", node.running)
}

func NewNodeOfFloat() *NodeOfFloat {
	node := NodeOfFloat{}
	node.in = make(chan float64)
	node.cin = make(chan chan float64)
	node.out = make(chan float64)
	node.cout = make(chan chan float64)
	node.cf = make(chan func(float64) float64)
	node.f = func(fl float64) float64 {
		return fl
	}
	node.running = false
	node.close = make(chan bool)
	node.Start()
	return &node
}

func (node *NodeOfFloat) Start() {
	if node.running {
		return
	}
	node.running = true
	go func(node *NodeOfFloat) {
		fmt.Printf("node starting...\n")
		for {
			select {
			case fl := <-node.in:
				node.out <- node.f(fl)
			case node.in = <-node.cin:
			case node.out = <-node.cout:
			case node.f = <-node.cf:

			case stop := <-node.close:
				if node.running && stop {
					fmt.Printf("node out closing...\n")
					node.running = false
					return
				}
			}
		}
	}(node)
}

func main() {
	node1 := NewNodeOfFloat()

	var i int
	go func() {
		for {
			time.Sleep(time.Millisecond * 50)
			i++
			node1.in <- float64(i)
		}
	}()

	go func() {
		for {
			fmt.Printf("%v ", <-node1.out)
		}
	}()

	time.Sleep(time.Second)
	node1.close <- true
	fmt.Println()

	go func() {
		for {
			fmt.Printf("%v ", <-node1.out)
		}
	}()
	node1.Start()
	time.Sleep(time.Second)
	fmt.Println()
}
