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
		fmt.Printf("node executing default f(%v)\n", fl)
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

			case node.open = <-node.copen:
				if node.open {
					node.Start()
				} else {
					fmt.Printf("node out closing...\n")
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

func (node *Node) Produce(fl float64){
	go func() {
		for {
			select {
			default:
				node.in<- fl
			}
		}
	}()
}

func (node *Node) Connect(nextNode *Node){
	node.cout<- nextNode.in
}

func main() {
	node1 := NewNode()
	node2 := NewNode()
	node3 := NewNode()

	n := 0.0
	node1.cf<- func(fl float64) float64 {
		time.Sleep(time.Millisecond*100)
		fmt.Printf("node executing node1.f(%v)\n", fl)
		n++
		return n
	}

	// node2.f is default, i.e. identity

	node3.cf<- func(fl float64) float64 {
		time.Sleep(time.Millisecond*100)
		fmt.Printf("node executing node3.f(%v)\n", fl)
		return fl*10
	}

	node1.Produce(0.0)
	node1.Connect(node2)

	go func() {
		for {
			fmt.Printf("from node2.out: %v\n", <-node2.out)
		}
	}()

	time.Sleep(time.Millisecond*300)

	node1.Connect(node3)

	go func() {
		for {
			fmt.Printf("from node3.out: %v\n", <-node3.out)
		}
	}()

	//time.Sleep(time.Millisecond*50)
}