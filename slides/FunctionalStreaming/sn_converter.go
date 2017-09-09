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

type Converter struct {
	in    chan interface{}
	cin   chan chan interface{}
	out   chan string
	cout  chan chan string
	f     func(interface{}) string
	cf    chan func(interface{}) string
	open  bool
	copen chan bool
}

func NewConverter() *Converter {
	node := Converter{}
	node.in = make(chan interface{})
	node.cin = make(chan chan interface{})
	node.out = make(chan string)
	node.cout = make(chan chan string)
	node.cf = make(chan func(interface{}) string)
	node.f = func(inf interface{}) string {
		return ""
	}
	node.copen = make(chan bool)
	node.Start()
	return &node
}

func (node *Converter) Start() {
	if node.open {
		return
	}
	go func() {
		fmt.Printf("Converter starting...\n")
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
					fmt.Printf("Converter out closing...\n")
					return
				}
			}
		}
	}()
	node.open = true
}

func (node *Converter) Stop() {
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



func (node *Node) ConnectConverter(nextNode *Converter) *Converter {
	go func() {
		for {
			intf := <-node.cout
			nextNode.in<- nextNode.f(intf)
		}
	}()
	return nextNode
}

func (node *Converter) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Node) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Node) ConnectFilter(nextNode *Filter) *Filter {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Filter) Connect(nextNode *Node) *Node {
	node.cout <- nextNode.in
	return nextNode
}

func (node *Filter) ConnectFilter(nextNode *Filter) *Filter {
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
		fmt.Printf("Filter starting...\n")
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
					fmt.Printf("Filter out closing...\n")
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

func main() {
	node1 := NewNode()
	delay := NewNode()
	filter1 := NewFilter()
	filter2 := NewFilter()
	node3 := NewNode()

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

	filter1.cf <- func(str string) bool {
		if i, err := strconv.Atoi(str); err == nil {
			if i%2 == 0 {
				return true
			}
		}
		return false
	}

	filter2.cf <- func(str string) bool {
		if i, err := strconv.Atoi(str); err == nil {
			if i < 5 {
				return true
			}
		}
		return false
	}

	node3.cf <- func(str string) string {
		return str + " 3"
	}

	converter1 := NewConverter()
	converter1.cf <- func(intf interface{}) string {
		if v, ok := intf.(int); ok {
			return strconv.Itoa(v)
		}
		if v, ok := intf.(float64); ok {
			return fmt.Sprint(v)
		}
		if v, ok := intf.(string); ok {
			return v
		}
		return ""
	}

	go func() {
		for {
			fmt.Printf("converter1: %v\n", <-converter1.out)
		}
	}()

	converter1.in<- "hallo"

	/*node1.Produce().ConnectConverter(converter1).Connect(delay).ConnectFilter(filter1).ConnectFilter(filter2).Connect(node3)

	go func() {
		for {
			fmt.Printf("NODE 3: %v\n", <-node3.out)
		}
	}()*/


}
