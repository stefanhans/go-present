package functionalstreams

// START_1 OMIT
type TeeOfInt struct {
	in     chan int
	out_a  chan int 		// HL
	out_b  chan int 		// HL
	cin    chan chan int
	cout_a chan chan int 	// HL
	cout_b chan chan int 	// HL
	close  chan bool
}
// END_1 OMIT

// START_2 OMIT
func (tee *TeeOfInt) Start() {
	go func() { for {
			select {
			case tee.in = <-tee.cin:
			case tee.out_a = <-tee.cout_a: 	// HL
			case tee.out_b = <-tee.cout_b: 	// HL
			case in := <-tee.in:			// HL
				go func(in int) { tee.out_a <- in }(in) // HL
				go func(in int) { tee.out_b <- in }(in) // HL
			case <-tee.close: return
	}}}()
}
// END_2 OMIT

// START_3 OMIT
func NewTeeOfInt() *TeeOfInt {
	tee := TeeOfInt{}
	tee.in = make(chan int)
	tee.out_a = make(chan int)  		// HL
	tee.out_b = make(chan int) 			// HL
	tee.cin = make(chan chan int)
	tee.cout_a = make(chan chan int)  	// HL
	tee.cout_b = make(chan chan int) 	// HL
	tee.close = make(chan bool)
	tee.Start()
	return &tee
}
// END_3 OMIT

// START_4 OMIT
func (node *NodeOfInt) ConnectTee(nextNode *TeeOfInt) *TeeOfInt {
	node.cout <- nextNode.in
	return nextNode
}
func (node *TeeOfInt) ConnectNodes(nextNodeA, nextNodeB *NodeOfInt) (
	trueNode, falseNode *NodeOfInt) {
	node.cout_a <- nextNodeA.in
	node.cout_b <- nextNodeB.in
	return nextNodeA, nextNodeB
}
// END_4 OMIT

// tee -> node
func (node *TeeOfInt) ConnectToA(nextNode *NodeOfInt) *NodeOfInt {
	node.cout_a <- nextNode.in
	return nextNode
}
func (node *TeeOfInt) ConnectToB(nextNode *NodeOfInt) *NodeOfInt {
	node.cout_b <- nextNode.in
	return nextNode
}

// tee -> tees
func (node *TeeOfInt) ConnectTees(nextNodeA, nextNodeB *TeeOfInt) (NodeA, NodeB *TeeOfInt) {
	node.cout_a <- nextNodeA.in
	node.cout_b <- nextNodeB.in
	return nextNodeA, nextNodeB
}

// tee -> tee
func (node *TeeOfInt) ConnectTeeA(nextNode *TeeOfInt) *TeeOfInt {
	node.cout_a <- nextNode.in
	return nextNode
}
func (node *TeeOfInt) ConnectTeeB(nextNode *TeeOfInt) *TeeOfInt {
	node.cout_b <- nextNode.in
	return nextNode
}

// Fork after register
