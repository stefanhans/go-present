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
	cdone chan bool
}

func NewNode(f func(float64) float64) Node {
	node := Node{}
	node.cin = make(Edge)
	node.cout = make(Edge)
	node.cdone = make(chan bool)

	node.f = f

	go func() {
		for {
			select {
			case x := <-node.cin:
				node.cout <- node.f(x)
			case <-node.cdone:
				fmt.Printf("\ndone signal arrived\n")
				close(node.cout)
			}
		}
	}()

	return node
}

func (producer *Node) Produce() {
	go func() {
		for {
			select {
			case <-producer.cdone:
				fmt.Printf("\nProduce() returns\n")
				return
			default:
				producer.cin <- producer.f(1.0)
			}
		}
	}()
}

func (sourceNode *Node) Connect(node Node) *Node {

	go func() {
		for {
			select {
			case x := <-sourceNode.cin:
				fmt.Printf("\ndone signal arrived\n")
				node.cout <- sourceNode.f(x)
			case <-sourceNode.cdone:
				fmt.Printf("\ndone signal arrived\n")
				close(sourceNode.cout)
			}
		}
	}()

	fmt.Printf("node connected\n")
	return &node
}

func (sourceNode *Node) ChangeConnection(oldNode Node, newNode Node) *Node {
	//oldNode.cdone<- true
	sourceNode.Connect(newNode)
	fmt.Printf("node connection changed\n")
	return &newNode
}

func (consumer *Node) Consume() {
	go func() {
		for {
			select {
			case fin := <-consumer.cin:
				fmt.Printf("xxxxxxxxxxxx\n")
				consumer.f(fin)
			}
		}
	}()
}

func main() {
	node := NewNode(func(fin float64) float64 {
		return fin*2
	})

	nextNode := NewNode(func(fin float64) float64 {
		return fin*3
	})

	node.Connect(nextNode)

	go func() {
		node.cin <- 1.2
	}()

	fmt.Printf("node.out: %v\n", <-nextNode.cout)
	time.Sleep(time.Second)
}
