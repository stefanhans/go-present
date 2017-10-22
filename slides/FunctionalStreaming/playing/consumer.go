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

func (producer *Node) Produce() {
	go func() {
		for {
			select {
			case <-producer.cdone:
				fmt.Printf("\nProduce() returns\n")
				return
			case producer.f = <-producer.cf:
				fmt.Printf("\nnew produce function arrived\n")
			default:
				producer.cout <- producer.f(1.0)
			}
		}
	}()
}

func (sourceNode *Node) Connect(node Node) *Node {
	node.Start(sourceNode.cout)
	fmt.Printf("node started\n")
	return &node
}

func (consumer *Node) Consume() {
	go func() {
		for {
			select {
			case fin := <-consumer.cin:
				consumer.f(fin)
			case consumer.f = <-consumer.cf:
				fmt.Printf("\nnew consumer function arrived\n")
			}
		}
	}()
}

func main() {
	producer := NewNode(func(fin float64) float64 {
		return 1.0
	})
	producer.Produce()
	fmt.Printf("producer started\n")

	consumer := NewNode(func(fin float64) float64 {
		fmt.Printf("cs%v ", fin)
		time.Sleep(time.Millisecond * 100)
		return 0.0
	})

	producer.
		Connect(consumer).
		Consume()
	fmt.Printf("all connected and consumer started\n")
	time.Sleep(time.Second)

	consumer.cf <- func(fin float64) float64 {
		fmt.Printf("%v (with delay) ", fin)
		time.Sleep(time.Millisecond * 100)
		return 0.0
	}

	time.Sleep(time.Second)

	producer.cdone <- true
}
