package main

import (
	"fmt"
	"strconv"
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
				node.in <- ""
			}
		}
	}()
	return node
}

func (node *Node) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Node) ConnectFilter(nextNode *Filter) *Filter {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Node) ConnectDistributor(nextNode *Distributor) *Distributor {
	node.cout <- nextNode.in
	return nextNode
}


type Filter struct {
	in    chan string
	cin   chan chan string
	out   chan string
	cout  chan chan string
	f     func(string) bool
	cf    chan func(string) bool
	open  bool
	copen chan bool
}

func NewFilter() *Filter {
	node := Filter{}
	node.in = make(chan string)
	node.cin = make(chan chan string)
	node.out = make(chan string)
	node.cout = make(chan chan string)
	node.cf = make(chan func(string) bool)
	node.f = func(str string) bool {
		return true
	}
	node.copen = make(chan bool)
	node.Start()
	return &node
}

func (node *Filter) Start() {
	if node.open {
		return
	}
	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case flin := <-node.in:
				if node.f(flin) {
					node.out <- flin
				}
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

func (node *Filter) Stop() {
	if !node.open {
		return
	}
	node.copen <- false
	return
}

func (node *Filter) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Filter) ConnectFilter(nextNode *Filter) *Filter {
	node.cout <- nextNode.in
	return nextNode
}

type Distributor struct {
	in    chan string
	cin   chan chan string
	outs  []chan string
	cout  chan chan string
	index int
	f     func(string) int
	cf    chan func(string) int
	open  bool
	copen chan bool
}

func NewDistributor() *Distributor {
	node := Distributor{}
	node.in = make(chan string)
	node.cin = make(chan chan string)
	node.cout = make(chan chan string)
	node.cf = make(chan func(string) int)
	node.f = func(str string) int {
		if node.index >= len(node.outs) {
			node.index = 0
		}
		node.index++
		return node.index - 1
	}
	node.copen = make(chan bool)
	node.Start()
	return &node
}

func (node *Distributor) Start() {
	if node.open {
		return
	}
	go func() {
		fmt.Printf("Distributor starting...\n")
		for {
			select {
			case str := <-node.in:
				//node.outs[0] <- str
				node.outs[node.f(str)] <- str
				//fmt.Printf("node.outs[%v]\n", node.f(str))
				//fmt.Printf("node.index: %v\n", node.index)
			case node.in = <-node.cin:
			case cout := <-node.cout:
				node.outs = append(node.outs, cout)
			case node.f = <-node.cf:
			case node.open = <-node.copen:
				if node.open {
					node.Start()
				} else {
					fmt.Printf("Distributor out closing...\n")
					return
				}
			}
		}
	}()
	node.open = true
}

func (node *Distributor) Stop() {
	if !node.open {
		return
	}
	node.copen <- false
	return
}

func (node *Distributor) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node1 := NewNode()
	delay := NewNode()
	node2 := NewNode()
	node3 := NewNode()
	node4 := NewNode()
	node5 := NewNode()

	var strings []string = make([]string, 0)
	for i := 0; i < 10; i++ {
		strings = append(strings, strconv.Itoa(i))
	}
	n := 0
	node1.cf <- func(str string) string {
		n++
		return strconv.Itoa(n)
	}

	delay.cf <- func(str string) string {
		time.Sleep(time.Millisecond * 100)
		return str
	}

	node2.cf <- func(str string) string {
		return str + " 2"
	}

	node3.cf <- func(str string) string {
		return str + " 3"
	}

	node4.cf <- func(str string) string {
		return str + " 4"
	}

	node5.cf <- func(str string) string {
		return str + " 5"
	}

	distributor1 := NewDistributor()
	distributor1.cout<- node2.in
	distributor1.cout<- node3.in
	distributor1.cout<- node4.in


	node1.Produce().Connect(delay).ConnectDistributor(distributor1)

	go func() {
		for {
			fmt.Printf("NODE 2: %v\n", <-node2.out)
		}
	}()

	go func() {
		for {
			fmt.Printf("NODE 3: %v\n", <-node3.out)
		}
	}()

	go func() {
		for {
			fmt.Printf("NODE 4: %v\n", <-node4.out)
		}
	}()
}
