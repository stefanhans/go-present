package functionalstreams

// START_1 OMIT
type NodeOfInt struct {
	in    chan int
	out   chan int
	f     func(int) int
	cin   chan chan int		// HL
	cout  chan chan int		// HL
	cf    chan func(int) int // HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for { select {
		case in := <-node.in: node.out <- node.f(in) // HL
		case node.in = <-node.cin:		// HL
		case node.out = <-node.cout:	// HL
		case node.f = <-node.cf: // HL
		case <-node.close: return
		}} }()
}
// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt() *NodeOfInt {
	node := NodeOfInt{}
	node.in = make(chan int)
	node.out = make(chan int)
	node.f = func(in int) int { return in }
	node.cin = make(chan chan int)
	node.cout = make(chan chan int)
	node.cf = make(chan func(int) int)
	node.close = make(chan bool)
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
func (node *NodeOfInt) SetFunc(f func(int) int) {
	node.cf <- f
}
// END_SETFUNC OMIT

// START_CALC OMIT
func (node *NodeOfInt) Calculate(calc func(int) int) *NodeOfInt {
	nextNode := NewNodeOfInt()
	nextNode.cf <- calc

	node.Connect(nextNode)
	return nextNode
}
// END_CALC OMIT


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