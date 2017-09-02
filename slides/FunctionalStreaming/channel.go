package main

import (
	"fmt"
	"time"
)

type Edge chan float64

type Node struct {
	cin Edge
	cout Edge
	f func(float64) float64
	cf chan func(float64) float64
}

/*
func (node *Node) Init() {
	go func() {

	}
}
*/

func main() {
	csource := make(Edge)
	go func() {
		for {
			csource <- 1.0
		}
	}()

	fmt.Printf("producer started\n")

	node := Node{}
	node.cin = csource
	node.cout = make(Edge)
	node.cf = make(chan func(float64) float64)

	node.f = func(fin float64) float64 {
		return fin*2
	}

	go func() {
		for {
			select {
			case x := <-node.cin:
				node.cout <-node.f(x)
			case node.f = <-node.cf:
				fmt.Printf("\nnew function arrived\n")
			}
		}
	}()

	fmt.Printf("node started\n")

	go func() {
		for {
			fmt.Printf("cs%v ", <-node.cout)
			time.Sleep(time.Millisecond *100)
		}
	}()

	fmt.Printf("consumer started\n")

	time.Sleep(time.Second)

	node.cf <- func(fin float64) float64 {
		return fin*3
	}

	time.Sleep(time.Second)
	fmt.Printf("\nmain last line\n")

}