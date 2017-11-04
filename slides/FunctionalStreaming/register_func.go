package main

import (
	"fmt"
	"time"
	"io/ioutil"
)

// START_1 OMIT
type NodeOfInt struct {
	in    chan int                                 // Input channel
	cin   chan chan int                            // can be exchanged.

	f     func(int) int                            // Function
	cf    chan func(int) int                       // can be exchanged.

	out   chan int                                 // Output channel
	cout  chan chan int                            // can be exchanged.

	cfuncname chan string
	func_register_map map[string]func(i int) int
	close chan bool // OMIT
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for { select {

		case in := <-node.in: node.out <- node.f(in) // Handle data (DEADLOCKS!) // HL

		case funcname := <-node.cfuncname:
			fmt.Printf("funcname: %v\n", funcname)
			//if f, ok := node.func_register_map["threeTimes"]; ok {
				//fmt.Printf("funcname: %v\n", funcname)
				node.f = node.func_register_map[funcname]
					//func(in int) int { return in * 3	}
			//}

		case node.in = <-node.cin:   	            // Change input channel
		case node.f = <-node.cf: 		            // Change function
		case node.out = <-node.cout: 	            // Change output channel
		case <-node.close: return // OMIT
		}
		}}()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in }        // Default function returns input value
	node.cf = make(chan func(int) int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.func_register_map = make(map[string]func(i int) int)

	node.cfuncname = make(chan string)
	node.close = make(chan bool) // OMIT
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfInt) Stop() {
	node.close <- true
}

// START_5 OMIT
func (node *NodeOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) { node.cf <- f }
// END_SETFUNC OMIT

// START_MAP OMIT
func (node *NodeOfInt) Map(f func(int) int) *NodeOfInt {
	nextNode := NewNodeOfInt()
	nextNode.cf <- f

	node.Connect(nextNode)
	return nextNode
}
// END_MAP OMIT


// START_3 OMIT
func (node *NodeOfInt) ProduceAtMs(n time.Duration) *NodeOfInt {
	go func() { for { select {
	default: node.in <- 0 }	               // Trigger permanently // HL
		time.Sleep(time.Millisecond * n)	      // with delay in ms // HL
	}}()
	return node
}
// END_3 OMIT


// START_PRINTF OMIT
func (node *NodeOfInt) Printf(format string) {
	go func() { for { select {
	case in := <-node.out: fmt.Printf(format, in)		// HL
	}}}()
}
// END_PRINTF OMIT


func main() {

	fmt.Println("Try to receive a string from a named pipe (mkfifo) /tmp/myPipe")

	strchan := make(chan string)

	go func() {
		for {
			b, err := ioutil.ReadFile("/tmp/myPipe")
			if err != nil {
				return
			}
			fmt.Printf("string(b): %v\n", string(b))
			strchan <- string(b)
		}
	}()

	node_1, node_2 := NewNodeOfInt(), NewNodeOfInt()

	node_2.func_register_map["twoTimes"] = func(in int) int { return in * 2	}
	node_2.func_register_map["threeTimes"] = func(in int) int { return in * 3 }

	var i int
	node_1.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	})


	go func() {
		for {
			select {
			case str := <-strchan:
				fmt.Printf("str: %v\n", str)
				node_2.cfuncname <- str
			}
		}
	}()

	node_1.Connect(node_2).Printf("%v ")
	node_1.ProduceAtMs(200)
	time.Sleep(time.Second)
	cmd := "twoTimes"
	node_2.SetFunc(node_2.func_register_map[cmd])

	time.Sleep(time.Second * 60)
}