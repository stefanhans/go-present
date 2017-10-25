package main

import (
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

// START_6 OMIT
func main() {
	node_1 := NewNodeOfInt()
	var i int
	node_1.SetFunc(func(in int) int {
		time.Sleep(time.Millisecond * 50)
		i++
		return in + i
	})

	switcher := NewSwitchOfInt()                  				// HL
	switcher.SetFunc(func(in int) bool { return (in%2 == 0) }) 	// HL

	node_2 := NewNodeOfInt()
	node_2.SetFunc(func(in int) int { return in * 10 })

	node_1.Produce().ConnectSwitch(switcher).ConnectToTrue(node_2).Print() 	// HL
	switcher.CloseFalse() 														// HL
	time.Sleep(time.Second)
}
// END_6 OMIT
