package main

import (
	"fmt"
	"time"
)

type NodeOfFloat struct {
	in    chan float64
	out   chan float64
	f     func(float64) float64
	cin   chan chan float64
	cout  chan chan float64
	cf    chan func(float64) float64
	close chan bool
}
func (node *NodeOfFloat) Start() {
	go func() {
		for {
			select {
			case node.in = <-node.cin:
			case node.out = <-node.cout:
			case fl := <-node.in:
				node.out <- node.f(fl)
			case node.f = <-node.cf:
			case <-node.close:
				return
			}
		}
	}()
}
func NewNodeOfFloat() *NodeOfFloat {
	node := NodeOfFloat{}
	node.in = make(chan float64)
	node.out = make(chan float64)
	node.f = func(fl float64) float64 { return fl }
	node.cin = make(chan chan float64)
	node.cout = make(chan chan float64)
	node.cf = make(chan func(float64) float64)
	node.close = make(chan bool)
	node.Start()
	return &node
}
func (node *NodeOfFloat) Stop() {
	node.close <- true
}
func (node *NodeOfFloat) Produce() *NodeOfFloat {
	go func() {
		for {
			select {
			default:
				node.in <- 0.0
			}
		}
	}()
	return node
}
func (node *NodeOfFloat) Consume() {
	go func() {
		for {
			select {
			case fl := <-node.out:
				fmt.Printf("%v ", fl)
			}
		}
	}()
}

// START_1 OMIT
type SwitchOfFloat struct {
	in         chan float64
	out_true   chan float64 // HL
	out_false  chan float64 // HL
	f          func(float64) bool
	cin        chan chan float64
	cout_true  chan chan float64 // HL
	cout_false chan chan float64 // HL
	cf         chan func(float64) bool
	close      chan bool
}
// END_1 OMIT

// START_2 OMIT
func (switcher *SwitchOfFloat) Start() {
	go func() { for { select {
			case switcher.in = <-switcher.cin:
			case switcher.out_true = <-switcher.cout_true: // HL
			case switcher.out_false = <-switcher.cout_false: // HL
			case fl := <-switcher.in: // HL
				if switcher.f(fl) { switcher.out_true <- fl // HL
				} else { switcher.out_false <- fl } // HL
			case switcher.f = <-switcher.cf:
			case <-switcher.close: return
	}}}()
}
// END_2 OMIT

// START_3 OMIT
func NewSwitchOfFloat() *SwitchOfFloat {
	switcher := SwitchOfFloat{}
	switcher.in = make(chan float64)
	switcher.out_true = make(chan float64)  // HL
	switcher.out_false = make(chan float64) // HL
	switcher.f = func(fl float64) bool { return true }
	switcher.cin = make(chan chan float64)
	switcher.cout_true = make(chan chan float64)  // HL
	switcher.cout_false = make(chan chan float64) // HL
	switcher.cf = make(chan func(float64) bool)
	switcher.close = make(chan bool)
	switcher.Start()
	return &switcher
}
// END_3 OMIT

// node -> node
func (node *NodeOfFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout <- nextNode.in
	return nextNode
}

// START_4 OMIT
// node -> switch
func (node *NodeOfFloat) ConnectSwitch(nextNode *SwitchOfFloat) *SwitchOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// END_4 OMIT


// START_5 OMIT
// switch -> nodes
func (node *SwitchOfFloat) ConnectNodes(nextNodeTrue, nextNodeFalse *NodeOfFloat) (trueNode, falseNode *NodeOfFloat) {
	node.cout_true <- nextNodeTrue.in
	node.cout_false <- nextNodeFalse.in
	return nextNodeTrue, nextNodeFalse
}

// switch -> node
func (node *SwitchOfFloat) ConnectToTrue(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout_true <- nextNode.in
	return nextNode
}
func (node *SwitchOfFloat) ConnectToFalse(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout_false <- nextNode.in
	return nextNode
}
// END_5 OMIT

// switch -> switches
func (node *SwitchOfFloat) ConnectSwitches(nextNodeTrue, nextNodeFalse *SwitchOfFloat) (trueNode, falseNode *SwitchOfFloat) {
	node.cout_true <- nextNodeTrue.in
	node.cout_false <- nextNodeFalse.in
	return nextNodeTrue, nextNodeFalse
}

// switch -> switch
func (node *SwitchOfFloat) ConnectSwitchToTrue(nextNode *SwitchOfFloat) *SwitchOfFloat {
	node.cout_true <- nextNode.in
	return nextNode
}
func (node *SwitchOfFloat) ConnectSwitchToFalse(nextNode *SwitchOfFloat) *SwitchOfFloat {
	node.cout_false <- nextNode.in
	return nextNode
}

// START_6 OMIT
func main() {
	node_1 := NewNodeOfFloat()
	var i float64
	node_1.cf <- func(fl float64) float64 {
		time.Sleep(time.Millisecond * 50)
		i++
		return fl + i
	}

	switcher := NewSwitchOfFloat()                                   // HL
	switcher.cf <- func(fl float64) bool { return (int(fl)%2 == 0) } // HL

	node_2 := NewNodeOfFloat()
	node_2.cf <- func(fl float64) float64 { return fl * 10 }
	node_3 := NewNodeOfFloat()
	node_3.cf <- func(fl float64) float64 { return fl * 0.5 }

	node_1.Produce().ConnectSwitch(switcher).ConnectToTrue(node_2).Consume() // HL
	switcher.ConnectToFalse(node_3).Consume() // HL
	time.Sleep(time.Second)
	fmt.Println("\n")
	switcher.ConnectNodes(node_2, node_3) // HL
	node_2.Consume() // HL
	node_3.Consume() // HL
	time.Sleep(time.Second)
}
// END_6 OMIT
