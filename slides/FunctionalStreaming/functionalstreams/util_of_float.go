package functionalstreams

import (
	"fmt"
	"math/rand"
	"time"
)

// START_1 OMIT
func (node *NodeOfFloat) Produce() *NodeOfFloat {
	go func() {
		for {
			select {
			default:
				node.in <- 0
			} // HL
		}
	}()
	return node
}

// END_1 OMIT

// START_2 OMIT
func (node *NodeOfFloat) Print() {
	go func() {
		for {
			select {
			case in := <-node.out: // HL
				fmt.Printf("%f ", in) // HL
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func (node *NodeOfFloat) ProduceAtMs(n time.Duration) *NodeOfFloat {
	go func() {
		for {
			select {
			default:
				node.in <- 0
			} // HL
			time.Sleep(time.Millisecond * n) // HL
		}
	}()
	return node
}

// END_3 OMIT

// START_4 OMIT
func (node *NodeOfFloat) ProduceRandPositivAtMs(mult float64, ms time.Duration) *NodeOfFloat {
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			select {
			default:
				node.in <- rand.Float64() * mult
			} // HL
			time.Sleep(time.Millisecond * ms) // HL
		}
	}()
	return node
}

// END_4 OMIT
