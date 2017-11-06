package main

import (
	"fmt"
	"strings"
	"time"
)

// START_1 OMIT
type NodeOfInt struct {
	in  chan int      // Input channel
	cin chan chan int // can be exchanged.

	f  func(int) int      // Function
	cf chan func(int) int // can be exchanged.

	out  chan int      // Output channel
	cout chan chan int // can be exchanged.

	name    string
	metain  chan string
	cmetain chan chan string

	methods map[string]func(string) string

	metaout  chan string
	cmetaout chan chan string

	close chan bool // OMIT
}

// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for {
			select {

			case in := <-node.in:
				node.out <- node.f(in) // Handle data (DEADLOCKS!) // HL

			case node.in = <-node.cin: // Change input channel
			case node.f = <-node.cf: // Change function
			case node.out = <-node.cout: // Change output channel

			case methodName := <-node.metain:
				cmdAndString := strings.SplitAfterN(methodName, " ", 2)
				/*
					fmt.Printf("methodName: %v\n", methodName)
					fmt.Printf("cmdAndString: %v\n", cmdAndString)
					fmt.Printf("len(cmdAndString): %v\n", len(cmdAndString))
					fmt.Printf("cmdAndString[0]: %v\n", cmdAndString[0])
					fmt.Printf("cmdAndString[1]: %v\n", cmdAndString[1])
				*/

				if fmeta, ok := node.methods[strings.Trim(cmdAndString[0], " ")]; ok {
					reply := fmeta(cmdAndString[1])
					//fmt.Printf("cmdAndString[0]: %v\ncmdAndString[1]: %v\n", cmdAndString[0], cmdAndString[1])
					node.metaout <- reply
				} else {
					fmt.Printf("UNKNOWN METHOD: %v%v\n", cmdAndString[0], cmdAndString[1])
				}

			case node.metain = <-node.cmetain: // Change metain channel
			case node.metaout = <-node.cmetaout: // Change metaout channel

			case str := <-node.metaout:
				fmt.Printf("%v ", str)

			case <-node.close:
				return // OMIT
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt(name string) *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in } // Default function returns input value
	node.cf = make(chan func(int) int)
	node.out = make(chan int)
	node.cout = make(chan chan int)

	node.name = name
	node.metain = make(chan string)
	node.cmetain = make(chan chan string)

	node.methods = make(map[string]func(string) string)
	node.methods["print"] = func(str string) string {
		return fmt.Sprintf("print %v\nPRINT: (%T) \"%v\"", str, node, node.name)
	}
	node.methods["log"] = func(str string) string {
		return fmt.Sprintf("log %v\nLOG: (%T) \"%v\"", str, node, node.name)
	}

	node.metaout = make(chan string)
	node.cmetaout = make(chan chan string)

	node.close = make(chan bool) // OMIT
	node.Start()
	return &node
}

// END_3 OMIT

// START_5 OMIT
func (node *NodeOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	node.cout <- nextNode.in
	node.cmetaout <- nextNode.metain
	return nextNode
}

// END_5 OMIT

// START_NODE_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) { node.cf <- f }

// END_NODE_SETFUNC OMIT

// START_PRINTF OMIT
func (node *NodeOfInt) Printf(format string) {
	go func() {
		for {
			select {
			case in := <-node.out:
				fmt.Printf(format, in) // HL
			}
		}
	}()
}

// END_PRINTF OMIT

// START_PRINTF OMIT
func (node *NodeOfInt) MetaPrintf(format string) {
	go func() {
		for {
			select {
			case metaout := <-node.metaout:
				fmt.Printf(format, strings.SplitAfterN(metaout, "\n", 2)[1]) // HL
			}
		}
	}()
}

// END_PRINTF OMIT

// START_ProduceAtMs OMIT
func (node *NodeOfInt) ProduceAtMs(n time.Duration) *NodeOfInt {
	go func() {
		for {
			select {
			default:
				node.in <- 0
			} // Trigger permanently // HL
			time.Sleep(time.Millisecond * n) // with delay in ms // HL
		}
	}()
	return node
}

// END_ProduceAtMs OMIT

// START_NODE_SETFUNC OMIT
func (node *NodeOfInt) ExecCmd(cmd string) { node.metain <- cmd + " " }

// END_NODE_SETFUNC OMIT

func main() {
	node_1, node_2, node_3 := NewNodeOfInt("node_1"), NewNodeOfInt("node_2"), NewNodeOfInt("node_3") // nodes' creation // HL
	var i int                                                                                        //
	node_1.SetFunc(func(in int) int { i++; return in + i })                                          //
	node_2.SetFunc(func(in int) int { return in * 2 })                                               //

	node_1.Connect(node_2).Connect(node_3).Printf("%v ") // stream configuration // HL
	node_1.ProduceAtMs(50)                               // sending data  // HL

	node_3.MetaPrintf("\n%v\n")

	node_1.ExecCmd("print")
	node_2.ExecCmd("log")

	time.Sleep(time.Second)
	fmt.Println()

	node_2.SetFunc(func(in int) int { return in * 10 }) // change function // HL
	time.Sleep(time.Second)
}
