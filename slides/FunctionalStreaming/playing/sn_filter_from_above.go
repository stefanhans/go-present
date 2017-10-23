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
		fmt.Printf("node starting...\n")
		for { select {
		case node.in = <-node.cin:
		case node.out = <-node.cout:
		case fl := <-node.in: node.out <- node.f(fl)
		case node.f = <-node.cf:
		case <-node.close: fmt.Printf("node closing...\n"); return
		}} }()
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
			select { default: node.in<- 0.0 }	// HL
		}}()
	return node
}
func (node *NodeOfFloat) Consume() {
	go func() {
		for {
			select {
			case fl := <-node.out: 				// HL
				fmt.Printf("%v ", fl)		// HL
			}
		}
	}()
}

// START_1 OMIT
type FilterOfFloat struct {
	in    chan float64
	out   chan float64
	f     func(float64) bool		// HL
	cin   chan chan float64
	cout  chan chan float64
	cf    chan func(float64) bool	// HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (filter *FilterOfFloat) Start() {
	go func() {
		fmt.Printf("filter starting...\n")
		for { select {
		case filter.in = <-filter.cin:
		case filter.out = <-filter.cout:
		case fl := <-filter.in: if filter.f(fl) { filter.out <- fl }	// HL
		case filter.f = <-filter.cf:
		case <-filter.close: fmt.Printf("filter closing...\n"); return
		}} }()
}
// END_2 OMIT

// START_3 OMIT
func NewFilterOfFloat() *FilterOfFloat {
	filter := FilterOfFloat{}
	filter.in = make(chan float64)
	filter.out = make(chan float64)
	filter.f = func(fl float64) bool { return true }	// HL
	filter.cin = make(chan chan float64)
	filter.cout = make(chan chan float64)
	filter.cf = make(chan func(float64) bool)			// HL
	filter.close = make(chan bool)
	filter.Start()
	return &filter
}
// END_3 OMIT

// START_4 OMIT
// node -> node
func (node *NodeOfFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// node -> filter
func (node *NodeOfFloat) ConnectFilter(nextNode *FilterOfFloat) *FilterOfFloat {
	node.cout <- nextNode.in
	return nextNode
}

// filter -> node
func (node *FilterOfFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// filter -> filter
func (node *FilterOfFloat) ConnectFilter(nextNode *FilterOfFloat) *FilterOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// END_4 OMIT

// START_5 OMIT
func main() {
	node_1 := NewNodeOfFloat()
	node_2 := NewNodeOfFloat()
	var i float64
	node_1.cf <- func(fl float64) float64 {
		i++
		return fl+i
	}
	node_2.cf <- func(fl float64) float64 { return fl * 3.0 }

	filter_1 := NewFilterOfFloat()	// HL
	filter_2 := NewFilterOfFloat()	// HL

	filter_1.cf <- func(fl float64) bool { return (int(fl)%33 == 0) }		// HL
	filter_2.cf <- func(fl float64) bool { return (int(fl)%100 == 0) }		// HL

	node_1.Produce().ConnectFilter(filter_1).ConnectFilter(filter_2).Connect(node_2).Consume() // HL

	time.Sleep(time.Millisecond * 50)
}
// END_5 OMIT
