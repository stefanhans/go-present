package main

import (
	"fmt"
	_ "time"
)

type Pusher interface {
	Push(*interface{}) *interface{}
}

type Node struct {
	in    chan string
	cin   chan chan string
	out   chan string
	cout  chan chan string
	f     func(string) string
	cf    chan func(string) string
}

func (node *Node) Push() {
	node.out<- node.f(<-node.in)
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

	go func() {
		fmt.Printf("node starting...\n")
		for {
			select {
			case flin := <-node.in:
				node.out <- node.f(flin)
			case node.in = <-node.cin:
			case node.out = <-node.cout:
			case node.f = <-node.cf:
			}
		}
	}()
	return &node
}

func (node *Node) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

type PushIt struct {
	i int
	fl float64
	str string
}

func (node *Node) Push(data *interface{}) *interface{} {

}

func main() {






	node1 := NewNode()
	node2 := NewNode()
	node3 := NewNode()

	node1.cf <- func(str string) string {
		return "1"
	}

	node2.cf <- func(str string) string {
		return str + " 2"
	}

	node3.cf <- func(str string) string {
		return str + " 3"
	}

	//node1.Produce().Connect(node2).Connect(node3)

	go func() {
		for {
			fmt.Printf("NODE 3: %v\n", <-node3.out)
		}
	}()
}
