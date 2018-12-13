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

// START_MonadOfInt_1 OMIT
type MonadOfInt struct {
	neutralElement int
	assocFunc      func(int, int) int
}

// END_MonadOfInt_1 OMIT

// START_MonadOfInt_2 OMIT
func NewMonadOfInt(neutralElement int, assocFunc func(int, int) int) *MonadOfInt { // HL
	monad := MonadOfInt{}
	monad.neutralElement = neutralElement // HL
	monad.assocFunc = assocFunc           // HL
	return &monad
}

// END_MonadOfInt_2 OMIT

// START_FolderOfInt_1 OMIT
type FolderOfInt struct {
	in      chan int
	cin     chan chan int
	monad   *MonadOfInt     // HL
	cmonad  chan MonadOfInt // HL
	result  int             // HL
	cresult chan int        // HL
	close   chan bool       // OMIT
}

// END_FolderOfInt_1 OMIT

// START_FolderOfInt_2 OMIT
func (folder *FolderOfInt) Start() {
	folder.result = folder.monad.neutralElement // HL
	go func() {
		for {
			select {
			case in := <-folder.in: // HL
				folder.result = folder.monad.assocFunc(folder.result, in) // HL
			case folder.in = <-folder.cin: // HL
			case <-folder.cresult:
				folder.cresult <- folder.result // HL
			case <-folder.close:
				return // OMIT
			}
		}
	}()
}

// END_FolderOfInt_2 OMIT

// START_FolderOfInt_3 OMIT
func NewFolderOfInt(monad *MonadOfInt) *FolderOfInt { // HL
	folder := FolderOfInt{}
	folder.in = make(chan int)
	folder.cin = make(chan chan int)
	folder.monad = monad                  // HL
	folder.cmonad = make(chan MonadOfInt) // HL
	folder.cresult = make(chan int)       // HL
	folder.close = make(chan bool)        // OMIT
	folder.Start()
	return &folder
}

// END_FolderOfInt_3 OMIT

// START_RESULT OMIT
func (folder *FolderOfInt) Result() int {
	folder.cresult <- 1
	return <-folder.cresult
}

// END_RESULT OMIT

func (node *NodeOfInt) ConnectFolder(nextNode *FolderOfInt) *FolderOfInt {
	node.cout <- nextNode.in
	return nextNode
}

func main() {
	node_in := NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	monad := NewMonadOfInt(0, func(result int, in int) int { return result + in })
	folder := NewFolderOfInt(monad)
	node_in.ConnectFolder(folder)
	node_in.ProduceAtMs(200)

	time.Sleep(time.Second)
	fmt.Printf("%v\n", folder.Result())
	time.Sleep(time.Second)
	fmt.Printf("%v\n", folder.Result())
}
