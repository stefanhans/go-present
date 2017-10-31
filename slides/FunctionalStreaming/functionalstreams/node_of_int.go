package functionalstreams

import (
	"fmt"
)

// START_1 OMIT
type NodeOfInt struct {
	name        string
	description string
	in          chan int
	cin         chan chan int
	f           func(int) int
	cf          chan func(int) int
	out         chan int
	cout        chan chan int
	metain      chan string
	cmetain     chan chan string
	fmeta       func(string) string
	cfmeta      chan func(string) string
	metaout     chan string
	cmetaout    chan chan string
	close       chan bool // OMIT
}

// END_1 OMIT

// START_2 OMIT
func (node *NodeOfInt) Start() {
	go func() {
		for {
			select {
			case in := <-node.in:
				node.out <- node.f(in)
			case node.in = <-node.cin:
			case node.f = <-node.cf:
			case node.out = <-node.cout:
			case metain := <-node.metain:
				node.metaout <- node.fmeta(metain)
			case node.metain = <-node.cmetain:
			case node.fmeta = <-node.cfmeta:
			case node.metaout = <-node.cmetaout:
			case <-node.close: return // OMIT
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewNodeOfInt(name string) *NodeOfInt {
	node := NodeOfInt{}
	node.name = name
	node.in = make(chan int)
	node.cin = make(chan chan int)
	node.f = func(in int) int { return in }
	node.cf = make(chan func(int) int)
	node.out = make(chan int)
	node.cout = make(chan chan int)
	node.metain = make(chan string)
	node.cmetain = make(chan chan string)
	node.fmeta = func(in string) string {
		return fmt.Sprintf("%v%v\n", in, node.String())
	}
	node.cfmeta = make(chan func(string) string)
	node.metaout = make(chan string)
	node.cmetaout = make(chan chan string)
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
	node.cmetaout <- nextNode.metain
	return nextNode
}

// END_5 OMIT

// START_STRING OMIT
func (node *NodeOfInt) String() string {
	str := "NODE: "+node.name+"\n"
	return str
}

func (node *NodeOfInt) Report() {
	node.metain <- "\nStart Report\n\n"
}

// END_STRING OMIT

// START_SETFUNC OMIT
func (node *NodeOfInt) SetFunc(f func(int) int) {
	node.cf <- f
}

// END_SETFUNC OMIT

// START_CALC OMIT
func (node *NodeOfInt) Calculate(calc func(int) int) *NodeOfInt {
	nextNode := NewNodeOfInt("<anonymous>")
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
