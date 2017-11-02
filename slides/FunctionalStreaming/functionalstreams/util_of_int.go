package functionalstreams

import (
	"fmt"
	"math/rand"
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
		case in := <-node.out: fmt.Printf("%v ", in)		// HL
		}}}()
}
// END_2 OMIT


// START_3 OMIT
func (node *NodeOfInt) ProduceAtMs(n time.Duration) *NodeOfInt {
	go func() { for { select {
			default: node.in <- 0 }	               // Trigger permanently // HL
			time.Sleep(time.Millisecond * n)	      // with delay in ms // HL
	}}()
	return node
}
// END_3 OMIT



// START_4 OMIT
func (node *NodeOfInt) ProduceRandPositivAtMs(max int, ms time.Duration) *NodeOfInt {
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			select { default: node.in <- rand.Intn(max)+1 }	 // HL
			time.Sleep(time.Millisecond * ms)	// HL
		}}()
	return node
}
// END_4 OMIT



// START_PRINTF OMIT
func (node *NodeOfInt) Printf(format string) {
	go func() { for { select {
		case in := <-node.out: fmt.Printf(format, in)		// HL
	}}}()
}
// END_PRINTF OMIT
