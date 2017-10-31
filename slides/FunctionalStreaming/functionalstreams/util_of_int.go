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
		case in := <-node.out: 				// HL
			fmt.Printf("%v ", in)		// HL // TODO adjust to fmt.Print
		case metain := <-node.metaout:	// HL
			fmt.Printf("%v%v Print()\n\n", metain, node)		// HL

		}}}()
}
// END_2 OMIT

func (node *NodeOfInt) Printf(format string) {
	go func() {
		for { select {
		case in := <-node.out: 				// HL
			fmt.Printf(format, in)		// HL
		case metain := <-node.metaout:	// HL
			fmt.Printf("%v%v Printf(format string)\n\n", metain, node)		// HL

		}}}()
}


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



// START_4 OMIT
func (node *NodeOfInt) ProduceRandPositivAtMs(max int, ms time.Duration) *NodeOfInt {
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			select { default: node.in <- rand.Intn(max)+1 }	// HL
			time.Sleep(time.Millisecond * ms)	// HL
		}}()
	return node
}
// END_4 OMIT
