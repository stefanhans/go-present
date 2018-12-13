package main

import (
	"fmt"
	"time"
)

// START_1 OMIT
type NodeOfInt struct {
	in  chan int      // Input channel
	cin chan chan int // can be exchanged.

	f  func(int) int      // Function
	cf chan func(int) int // can be exchanged.

	out   chan int      // Output channel
	cout  chan chan int // can be exchanged.
	close chan bool     // OMIT
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
			case <-node.close:
				return // OMIT
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in } // Default function returns input value
	node.cf = make(chan func(int) int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.close = make(chan bool) // OMIT
	node.Start()
	return &node
}

// END_3 OMIT

// START_5 OMIT
func (node *NodeOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	node.cout <- nextNode.in
	return nextNode
}

// END_5 OMIT

// START_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) { node.cf <- f }

// END_SETFUNC OMIT

// START_1 OMIT
type BufferOfInt struct {
	in     chan int
	out    chan int
	buffer chan int
	flush  func(int) bool
	cin    chan chan int       // HL
	cout   chan chan int       // HL
	cflush chan func(int) bool // HL
	close  chan bool
}

// END_1 OMIT

// START_2 OMIT
func (buffer *BufferOfInt) Start() {
	go func() {
		for {
			select {
			case buffer.in = <-buffer.cin: // HL
			case buffer.out = <-buffer.cout: // HL
			case in := <-buffer.in: // HL
				buffer.buffer <- in   // HL
				if buffer.flush(in) { // HL
					for len(buffer.buffer) > 0 { // HL
						buf := <-buffer.buffer // HL
						buffer.out <- buf      // HL
					} // HL
				} // HL
			case buffer.flush = <-buffer.cflush: // HL
			case <-buffer.close:
				return
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewBufferOfInt() *BufferOfInt {
	buffer := BufferOfInt{}
	buffer.in = make(chan int)
	buffer.out = make(chan int)
	buffer.cin = make(chan chan int)
	buffer.cout = make(chan chan int)
	buffer.buffer = make(chan int, 1024)             // HL
	buffer.flush = func(in int) bool { return true } // HL
	buffer.cflush = make(chan func(int) bool)        // HL
	buffer.close = make(chan bool)
	buffer.Start()
	return &buffer
}

// END_3 OMIT

func (buffer *BufferOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	buffer.cout <- nextNode.in
	return nextNode
}

// START_FLUSH OMIT
func (buffer *BufferOfInt) Flush() {
	buffer.cflush <- func(i int) bool {
		return true
	}
}

// END_FLUSH OMIT

// START_BUFFER OMIT
func (buffer *BufferOfInt) Buffer() {
	buffer.cflush <- func(i int) bool {
		return false
	}
}

// END_BUFFER OMIT

// START_FUNC OMIT
func (buffer *BufferOfInt) SetFunc(f func(int) bool) {
	buffer.cflush <- f
}
func (buffer *BufferOfInt) Len() int { return len(buffer.buffer) }
func (buffer *BufferOfInt) Cap() int { return cap(buffer.buffer) }

// END_FUNC OMIT

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

// START_3 OMIT
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

// END_3 OMIT

func (node *NodeOfInt) ConnectBuffer(nextNode *BufferOfInt) *BufferOfInt {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node_in, node_out := NewNodeOfInt(), NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	buffer := NewBufferOfInt()

	node_in.ConnectBuffer(buffer).Connect(node_out).Printf("%v ")
	node_in.ProduceAtMs(200)

	buffer.SetFunc(func(i int) bool {
		return buffer.Len()%5 == 0
	})
	time.Sleep(time.Second * 5)
}
