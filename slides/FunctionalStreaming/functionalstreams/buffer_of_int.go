package functionalstreams

// START_1 OMIT
type BufferOfInt struct {
	in           chan int
	out          chan int
	buffer       chan int
	flush        func(int) bool
	cin          chan chan int       // HL
	cout         chan chan int       // HL
	cflush       chan func(int) bool // HL
	close        chan bool
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
				buffer.buffer <- in // HL
				if buffer.flush(in) { // HL
					for len(buffer.buffer) > 0 { // HL
						buf := <-buffer.buffer // HL
						buffer.out <- buf // HL
					} // HL
				} // HL
			case buffer.flush = <-buffer.cflush: // HL
			case <-buffer.close: return
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
	buffer.buffer = make(chan int, 1024) // HL
	buffer.flush = func(in int) bool { return true } // HL
	buffer.cflush = make(chan func(int) bool) // HL
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

