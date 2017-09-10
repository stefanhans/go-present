package main

import (
	"fmt"
	"time"
)

type Node struct {
	in    chan int
	cin   chan chan int
	out   chan int
	cout  chan chan int
	f     func(int) int
	cf    chan func(int) int
	open  bool
	copen chan bool
}

func NewNode() *Node {
	node := Node{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.cf = make(chan func(int) int)
	node.f = func(i int) int {
		return i
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
				node.in <- 0
			}
		}
	}()
	return node
}

func (node *Node) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Node) ConnectGroupBy(nextNode *GroupBy) *GroupBy {
	node.cout <- nextNode.in
	return nextNode
}

//####################################
type GroupBy struct {
	data    map[int]int
	insert  chan int
	request chan bool
	cdata   chan map[int]int
	in      chan int
	cin     chan chan int
	out     chan int
	cout    chan chan int
	f       func(int) int
	cf      chan func(int) int
	open    bool
	copen   chan bool
}

func NewGroupBy() *GroupBy {
	node := GroupBy{}
	node.data = make(map[int]int)
	node.insert = make(chan int)
	node.request = make(chan bool)
	node.cdata = make(chan map[int]int)
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.cf = make(chan func(int) int)
	node.f = func(i int) int {
		return i
	}
	node.copen = make(chan bool)
	node.Start()
	return &node
}

func (node *GroupBy) Start() {
	if node.open {
		return
	}
	go func() {
		fmt.Printf("GroupBy starting...\n")
		for {
			select {
			case flin := <-node.in:
				fmt.Printf("GroupBy<-node.in: %v\n", flin)
				node.out <- node.f(flin)
			case node.in = <-node.cin:
			case node.out = <-node.cout:
			case node.f = <-node.cf:
			case mod := <-node.insert:
				fmt.Printf("GroupBy<-node.insert: %v\n", mod)
				node.data[mod] = node.data[mod] + 1
			case <-node.request:
				node.cdata <- node.data
			case node.open = <-node.copen:
				if node.open {
					node.Start()
				} else {
					fmt.Printf("GroupBy out closing...\n")
					return
				}
			}
		}
	}()
	node.open = true
}

func (node *GroupBy) Stop() {
	if !node.open {
		return
	}
	node.copen <- false
	return
}

func (node *GroupBy) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node1 := NewNode()
	delay := NewNode()
	show := NewNode()

	n := 0
	node1.cf <- func(i int) int {
		n++
		return n
	}

	delay.cf <- func(i int) int {
		time.Sleep(time.Millisecond * 100)
		return i
	}

	groupby := NewGroupBy()
	groupby.f = func(i int) int {
		fmt.Printf("groupby.f...\n")
		//groupby.insert <- 1//i % 3    !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		return i
	}

	node1.Produce().Connect(delay).ConnectGroupBy(groupby).Connect(show)

	go func() {
		for {
			fmt.Printf("SHOW: %v\n", <-show.out)
		}
	}()

	go func() {
		fmt.Printf("groupby: %v\n", <-groupby.cdata)
	}()

	time.Sleep(time.Millisecond * 500)

	groupby.request <- true
}
