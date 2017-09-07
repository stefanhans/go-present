package main

import (
	"fmt"
	"time"
)

type Node struct {
	in    chan string
	cin   chan chan string
	out   chan string
	cout  chan chan string
	f     func(string) string
	cf    chan func(string) string
	open  bool
	copen chan bool
}

func NewNode() *Node {
	node := Node{}
	node.in = make(chan string)
	node.cin = make(chan chan string)
	node.out = make(chan string)
	node.cout = make(chan chan string)
	node.cf = make(chan func(string) string)
	node.f = func(str string) string {
		//fmt.Printf("node executing default f(%v)\n", str)
		return str
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
				node.out <- node.f(flin)
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
	if !node.open {
		return
	}
	node.copen <- false
	return
}

func (node *Node) Produce(str string) {
	go func() {
		for {
			select {
			default:
				node.in <- str
			}
		}
	}()
}

func (node *Node) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node1 := NewNode()
	node2 := NewNode()
	node3 := NewNode()

	n := "node1.f-> "
	node1.cf <- func(str string) string {
		time.Sleep(time.Millisecond * 100)
		//fmt.Printf("node executing node1.f(%v)\n", str)
		n += "."
		return n
	}

	node2.cf <- func(str string) string {
		time.Sleep(time.Millisecond * 100)
		return str + " 2"
	}

	node3.cf <- func(str string) string {
		time.Sleep(time.Millisecond * 100)
		return str + " 3"
	}

	node1.Produce("")
	node1.Connect(node2).Connect(node3)

	go func() {
		for {
			fmt.Printf("NODE 3: %v\n", <-node3.out)
		}
	}()
}
