package main

import (
	"fmt"
	"time"
)

type Node struct {
	in   chan float64
	cin chan chan float64
	out  chan float64
	cout chan chan float64
	f     func(float64) float64
	cf    chan func(float64) float64
	open bool
	copen chan bool
}



func NewNode() *Node {
	node := Node{}
	node.in = make(chan float64)
	node.cin = make(chan chan float64)
	node.out = make(chan float64)
	node.cout = make(chan chan float64)
	node.cf = make(chan func(float64) float64)
	node.f = func(fl float64) float64 {
		fmt.Printf("node executing f(%v)\n", fl)
		return fl
	}
	node.copen = make(chan bool)
	node.Start()
	return &node
}

func (node *Node) Start() {
	if node.open {
		return
	}
	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case flin := <-node.in:
				node.out<- node.f(flin)
			case node.in = <-node.cin:
			case node.out = <-node.cout:
			case node.f = <-node.cf:

			case opin := <-node.copen:
				if opin {
					node.Start()
				} else {
					close(node.out)
					return
				}
			}
		}
	}()
	node.open = true
}

func (node *Node) Stop() {
	if ! node.open {
		return
	}
	node.copen<- false
	return
}

func main() {
	node1 := NewNode()

	n := 0.0
	node1.cf<- func(fl float64) float64 {
		n++
		return n
	}

	go func() {
		for {
			fmt.Printf("received from node1: %v\n", <-node1.out)
		}
	}()
	node1.in<- 0.0
	node1.in<- 0.0

	time.Sleep(time.Second)
}