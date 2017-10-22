package main

import (
	"fmt"
	"time"
)

type Edge chan float64

type Node struct {
	cin   Edge
	cout  Edge
	f     func(float64) float64
	cf    chan func(float64) float64
	cdone chan bool
}

func NewNode(f func(float64) float64) Node {
	node := Node{}
	node.cout = make(Edge)
	node.cf = make(chan func(float64) float64)
	node.cdone = make(chan bool)

	node.f = f
	return node
}

func (node *Node) Start(cin Edge) {
	node.cin = cin
	go func() {
		for {
			select {
			case x := <-node.cin:
				node.cout <- node.f(x)
			case node.f = <-node.cf:
				fmt.Printf("\nnew function arrived\n")
			case <-node.cdone:
				fmt.Printf("\ndone signal arrived\n")
				close(node.cout)
			}
		}
	}()
}

func (cin Edge) Connect(f func(fin float64) float64) Node {
	node := NewNode(f)
	node.Start(cin)
	fmt.Printf("node started\n")
	return node
}

func main() {
	csource := make(Edge)
	cdone := make(chan bool)
	go func() {
		for {
			select {
			case <-cdone:
				fmt.Printf("\nreturn source's goroutine\n")
				return
			default:
				csource <- 1.0
			}
		}
	}()
	fmt.Printf("producer started\n")

	/*node := NewNode(func(fin float64) float64 {
		return fin * 2
	})
	node.Start(csource)
	fmt.Printf("node started\n")*/

	node := csource.Connect(func(fin float64) float64 { return fin * 2 })

	/*
	go func() {
		for {
			if n, ok := <-node.cout; !ok {
				fmt.Printf("\nnode's out channel closed\n")
			} else {
				node.cout<- n
			}
		}
	}()
	fmt.Printf("monitor started\n")
	*/

	go func() {
		for {
			fmt.Printf("cs%v ", <-node.cout)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	fmt.Printf("consumer started\n")

	time.Sleep(time.Second)

	node.cf <- func(fin float64) float64 {
		return fin * 3
	}

	time.Sleep(time.Second)

	cdone <- true

}
