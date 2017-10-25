package functionalstreams

// START_1 OMIT
type SwitchOfInt struct {
	in         chan int
	out_true   chan int 			// HL
	out_false  chan int 			// HL
	f          func(int) bool
	cin        chan chan int
	cout_true  chan chan int 		// HL
	cout_false chan chan int 		// HL
	cf         chan func(int) bool
	close      chan bool
}
// END_1 OMIT

// START_2 OMIT
func (switcher *SwitchOfInt) Start() {
	go func() { for { select {
	case switcher.in = <-switcher.cin:
	case switcher.out_true = <-switcher.cout_true: 		// HL
	case switcher.out_false = <-switcher.cout_false: 	// HL
	case in := <-switcher.in: 							// HL
		if switcher.f(in) { switcher.out_true <- in 	// HL
		} else { switcher.out_false <- in } 			// HL
	case switcher.f = <-switcher.cf:
	case <-switcher.close: return
	}}}()
}
// END_2 OMIT

// START_3 OMIT
func NewSwitchOfInt() *SwitchOfInt {
	switcher := SwitchOfInt{}
	switcher.in = make(chan int)
	switcher.out_true = make(chan int)  				// HL
	switcher.out_false = make(chan int) 				// HL
	switcher.f = func(in int) bool { return true }
	switcher.cin = make(chan chan int)
	switcher.cout_true = make(chan chan int)  			// HL
	switcher.cout_false = make(chan chan int) 			// HL
	switcher.cf = make(chan func(int) bool)
	switcher.close = make(chan bool)
	switcher.Start()
	return &switcher
}
// END_3 OMIT

// START_4 OMIT
// node -> switch
func (node *NodeOfInt) ConnectSwitch(nextNode *SwitchOfInt) *SwitchOfInt {
	node.cout <- nextNode.in
	return nextNode
}
// END_4 OMIT


// START_5 OMIT
// switch -> nodes
func (node *SwitchOfInt) ConnectNodes(nextNodeTrue, nextNodeFalse *NodeOfInt) (
												trueNode, falseNode *NodeOfInt) {
	node.cout_true <- nextNodeTrue.in
	node.cout_false <- nextNodeFalse.in
	return nextNodeTrue, nextNodeFalse
}

// switch -> node
func (node *SwitchOfInt) ConnectToTrue(nextNode *NodeOfInt) *NodeOfInt {
	node.cout_true <- nextNode.in
	return nextNode
}
func (node *SwitchOfInt) ConnectToFalse(nextNode *NodeOfInt) *NodeOfInt {
	node.cout_false <- nextNode.in
	return nextNode
}
// END_5 OMIT

// START_CLOSE OMIT
// switch close
func (node *SwitchOfInt) CloseTrue() {
	go func() {
		for {
			select {
			case _ = <-node.out_true:
			}
		}
	}()
}
func (node *SwitchOfInt) CloseFalse() {
	go func() {
		for {
			select {
			case _ = <-node.out_false:
			}
		}
	}()
}
// END_CLOSE OMIT

// switch -> switches
func (node *SwitchOfInt) ConnectSwitches(nextNodeTrue, nextNodeFalse *SwitchOfInt) (trueNode, falseNode *SwitchOfInt) {
	node.cout_true <- nextNodeTrue.in
	node.cout_false <- nextNodeFalse.in
	return nextNodeTrue, nextNodeFalse
}

// switch -> switch
func (node *SwitchOfInt) ConnectSwitchToTrue(nextNode *SwitchOfInt) *SwitchOfInt {
	node.cout_true <- nextNode.in
	return nextNode
}
func (node *SwitchOfInt) ConnectSwitchToFalse(nextNode *SwitchOfInt) *SwitchOfInt {
	node.cout_false <- nextNode.in
	return nextNode
}


// START_SWITCH OMIT
// node(f) -> node
func (node *NodeOfInt) Switch(fswitch func(int) bool) (NodeA, NodeB *NodeOfInt) {
	switcher := NewSwitchOfInt()
	switcher.cf <- fswitch
	node.ConnectSwitch(switcher)

	nodeA := NewNodeOfInt()
	nodeB := NewNodeOfInt()
	switcher.ConnectToTrue(nodeA)
	switcher.ConnectToFalse(nodeB)

	return nodeA, nodeB
}
// END_SWITCH OMIT


// START_FILTER OMIT
// node(f) -> node
func (node *NodeOfInt) Filter(filter func(int) bool) *NodeOfInt {
	switcher := NewSwitchOfInt()
	switcher.cf <- filter

	node.ConnectSwitch(switcher)

	nextNode := NewNodeOfInt()
	switcher.ConnectToTrue(nextNode)

	switcher.CloseFalse()

	return nextNode
}
// END_FILTER OMIT

func (node *SwitchOfInt) SetFunc(f func(int) bool) {
	node.cf <- f
}

