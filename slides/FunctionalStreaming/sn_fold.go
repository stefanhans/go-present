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

func (node *Node) Produce() *Node {
	go func() {
		for {
			select {
			default:
				node.in<- ""
			}
		}
	}()
	return node
}

func (node *Node) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Node) ConnectFold(nextNode *Fold) *Fold {
	node.cout <- nextNode.in
	return nextNode
}

//####################################
type Fold struct {
	data interface{}
	request chan bool
	cdata chan interface{}
	in    chan string
	cin   chan chan string
	out   chan string
	cout  chan chan string
	f     func(string) string
	cf    chan func(string) string
	open  bool
	copen chan bool
}

func NewFold() *Fold {
	node := Fold{}
	node.request = make(chan bool)
	node.cdata = make(chan interface{})
	node.in = make(chan string)
	node.cin = make(chan chan string)
	node.out = make(chan string)
	node.cout = make(chan chan string)
	node.cf = make(chan func(string) string)
	node.f = func(str string) string {
		return str
	}
	node.copen = make(chan bool)
	node.Start()
	return &node
}

func (node *Fold) Start() {
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
			case <-node.request:
				node.cdata<- node.data
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

func (node *Fold) Stop() {
	if !node.open {
		return
	}
	node.copen <- false
	return
}

func (node *Fold) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node1 := NewNode()
	delay := NewNode()
	node2 := NewNode()
	node3 := NewNode()

	node1.cf <- func(str string) string {
		return "1"
	}

	delay.cf<- func(str string) string {
		time.Sleep(time.Millisecond * 100)
		return str
	}

	node2.cf <- func(str string) string {
		return str + " 2"
	}

	node3.cf <- func(str string) string {
		return str + " 3"
	}

	counter := NewFold()
	counter.data = 0
	counter.f = func(str string) string {
		if v, ok := counter.data.(int); ok {
			counter.data = v + 1
		}
		return str
	}

	node1.Produce().Connect(delay).Connect(node2).ConnectFold(counter).Connect(node3)

	go func() {
		for {
			fmt.Printf("NODE 3: %v\n", <-node3.out)
		}
	}()

	go func() {
		fmt.Printf("COUNTER: %v passed\n", <-counter.cdata)
	}()

	time.Sleep(time.Millisecond * 500)

	counter.request<- true
}
