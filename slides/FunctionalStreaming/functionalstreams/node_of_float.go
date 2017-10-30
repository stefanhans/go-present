package functionalstreams

// START_1 OMIT
type NodeOfFloat struct {
	in    chan float64
	out   chan float64
	f     func(float64) float64
	cin   chan chan float64		// HL
	cout  chan chan float64		// HL
	cf    chan func(float64) float64 // HL
	close chan bool
}
// END_1 OMIT

// START_2 OMIT
func (node *NodeOfFloat) Start() {
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
func NewNodeOfFloat() *NodeOfFloat {
	node := NodeOfFloat{}
	node.in = make(chan float64)
	node.out = make(chan float64)
	node.f = func(in float64) float64 { return in }
	node.cin = make(chan chan float64)
	node.cout = make(chan chan float64)
	node.cf = make(chan func(float64) float64)
	node.close = make(chan bool)
	node.Start()
	return &node
}
// END_3 OMIT

func (node *NodeOfFloat) Stop() {
	node.close <- true
}

// START_5 OMIT
func (node *NodeOfFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	node.cout <- nextNode.in
	return nextNode
}
// END_5 OMIT

/*
// START_SETFUNC OMIT
func (node *NodeOfFloat) SetFunc(f func(float64) float64) {
	node.cf <- f
}
// END_SETFUNC OMIT

// START_CALC OMIT
func (node *NodeOfFloat) Calculate(calc func(float64) float64) *NodeOfFloat {
	nextNode := NewNodeOfFloat()
	nextNode.cf <- calc

	node.Connect(nextNode)
	return nextNode
}
// END_CALC OMIT


func (node *NodeOfFloat) ConnectBuffer(nextNode *BufferOfInt) *BufferOfInt {
	node.cout <- nextNode.in
	return nextNode
}


func (node *NodeOfFloat) ConnectAggregator(nextNode *AggregatorOfInt) *AggregatorOfInt {
	node.cout <- nextNode.in
	return nextNode
}

func (node *NodeOfFloat) ConnectFolder(nextNode *FolderOfInt) *FolderOfInt {
	node.cout <- nextNode.in
	return nextNode
}
*/



