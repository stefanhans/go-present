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

func (node *Node) ConnectDistributor(nextNode *Distributor) *Distributor {
	node.cout <- nextNode.in
	return nextNode
}

type Distributor struct {
	in    chan string
	cin   chan chan string
	outs  []chan string
	cout  chan chan string
	open  bool
	copen chan bool
}

func NewDistributor() *Distributor {
	node := Distributor{}
	node.in = make(chan string)
	node.cin = make(chan chan string)
	node.cout = make(chan chan string)
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
				for i:=0; i<len(node.outs); i++ {
					node.outs[i]<- str
				}
			case node.in = <-node.cin:
			case cout := <-node.cout:
				node.outs = append(node.outs, cout)
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

func (node *Distributor) AddNode(nextNode *Node) *Distributor {
	node.cout<- nextNode.in
	return node
}

type Collector struct {
	in    chan string
	cin   chan chan string
	out   chan string
	cout  chan chan string
	f     func(string) string
	cf    chan func(string) string
	open  bool
	copen chan bool
}

func NewCollector() *Collector {
	node := Collector{}
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

func (node *Collector) Start() {
	if node.open {
		return
	}
	go func() {
		fmt.Printf("Collector starting...\n")
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
					fmt.Printf("Collector out closing...\n")
					return
				}
			}
		}
	}()
	node.open = true
}

func (node *Collector) Stop() {
	if !node.open {
		return
	}
	node.copen <- false
	return
}

func (node *Collector) AddNode(nextNode *Node) *Collector {
	go func() {
		for {
			node.in<- <-nextNode.out
		}
	}()
	return node
}

func (node *Collector) Connect(nextNode *Node) *Node {
	node.cout<- nextNode.in
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
	node1.Produce().Connect(delay).ConnectDistributor(distributor1).AddNode(node2).AddNode(node3).AddNode(node4)

	collector1 := NewCollector()
	collector1.AddNode(node2).AddNode(node3).AddNode(node4).Connect(node5)

	/*go func() {
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
	}()*/

	go func() {
		for {
			fmt.Printf("NODE 5: %v\n", <-node5.out)
		}
	}()
}
