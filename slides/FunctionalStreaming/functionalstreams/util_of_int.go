package functionalstreams

import (
	"fmt"
	"time"
)

// START_1 OMIT
func (node *NodeOfInt) Produce() *NodeOfInt {
	go func() {
		for {
			select { default: node.in <- 0 }	// HL
		}}()
	return node
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Print() {
	go func() {
		for { select {
		case in := <-node.out: 				// HL
			fmt.Printf("%v ", in)		// HL
		}}}()
}
// END_2 OMIT


// START_3 OMIT
func (node *NodeOfInt) ProduceAtMs(n time.Duration) *NodeOfInt {
	go func() {
		for {
			select { default: node.in <- 0 }	// HL
			time.Sleep(time.Millisecond * n)	// HL
		}}()
	return node
}
// END_3 OMIT
