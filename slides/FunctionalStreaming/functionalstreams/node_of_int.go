package functionalstreams

// START_1 OMIT
type NodeOfInt struct {
	in    chan int                                 // Input channel
	cin   chan chan int                            // can be exchanged.

	f     func(int) int                            // Function
	cf    chan func(int) int                       // can be exchanged.

	out   chan int                                 // Output channel
	cout  chan chan int                            // can be exchanged.
	close chan bool // OMIT
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for { select {

			case in := <-node.in: node.out <- node.f(in) // Handle data (DEADLOCKS!) // HL

			case node.in = <-node.cin:   	            // Change input channel
			case node.f = <-node.cf: 		            // Change function
			case node.out = <-node.cout: 	            // Change output channel
			case <-node.close: return // OMIT
			}
	}}()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in }        // Default function returns input value
	node.cf = make(chan func(int) int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.close = make(chan bool) // OMIT
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfInt) Stop() {
	node.close <- true
}

// START_5 OMIT
func (node *NodeOfInt) Connect(nextNode *NodeOfInt) *NodeOfInt {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) { node.cf <- f }
// END_SETFUNC OMIT

// START_MAP OMIT
func (node *NodeOfInt) Map(f func(int) int) *NodeOfInt {
	nextNode := NewNodeOfInt()
	nextNode.cf <- f

	node.Connect(nextNode)
	return nextNode
}
// END_MAP OMIT


func (node *NodeOfInt) ConnectBuffer(nextNode *BufferOfInt) *BufferOfInt {
	node.cout <- nextNode.in
	return nextNode
}


func (node *NodeOfInt) ConnectAggregator(nextNode *AggregatorOfInt) *AggregatorOfInt {
	node.cout <- nextNode.in
	return nextNode
}

func (node *NodeOfInt) ConnectFolder(nextNode *FolderOfInt) *FolderOfInt {
	node.cout <- nextNode.in
	return nextNode
}

// START_CONNECTCONVERTER OMIT
func (node *NodeOfInt) ConnectConverterIntToFloat(nextNode *ConverterIntToFloat) *ConverterIntToFloat { // HL
	node.cout <- nextNode.in
	return nextNode
}
// END_CONNECTCONVERTER OMIT