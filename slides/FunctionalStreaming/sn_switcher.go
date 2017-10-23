package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

// START_6 OMIT
func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.Cf <- func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	}

	switcher := NewSwitchOfInt()                  					// HL
	switcher.Cf <- func(in int) bool { return (in%2 == 0) } 	// HL

	node_2 := NewNodeOfInt()
	node_2.Cf <- func(in int) int { return in * 10 }
	node_3 := NewNodeOfInt()
	node_3.Cf <- func(in int) int { return in * 2 }

	node_1.Produce().ConnectSwitch(switcher).ConnectToTrue(node_2).Consume() 	// HL
	switcher.ConnectToFalse(node_3).Consume() 									// HL
	time.Sleep(time.Second)
}
// END_6 OMIT

/*

	fmt.Println("\n")
	switcher.ConnectNodes(node_2, node_3) // HL
	node_2.Consume() // HL
	node_3.Consume() // HL
	time.Sleep(time.Second)

 */
