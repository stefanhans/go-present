package main

import (
	"fmt"
	"math/rand"
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
			fmt.Printf("cs%v ", <-consumer.cout)
			time.Sleep(time.Millisecond * 10)
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
		return 0.0
	})
	consumer.Consume()
	fmt.Printf("consumer started\n")

	tentimes := NewNode(func(fin float64) float64 {
		return fin * 10
	})

	producer.
	Connect(tentimes).
		Connect(tentimes).
		Connect(NewNode(func(fin float64) float64 { return fin / 3 })).
		Connect(NewNode(func(fin float64) float64 { return fin * 3 })).
		Consume()
	fmt.Printf("all connected\n")
	time.Sleep(time.Second)

	producer.cf <- func(fin float64) float64 {
		return 2.0
	}
	time.Sleep(time.Second)

	rand.Seed(time.Now().UnixNano())
	producer.cf <- func(fin float64) float64 {
		return rand.Float64()
	}
	time.Sleep(time.Second)

	floats := [7]float64{0.0, 3.0, 4.0, 5.0, 7.0, 8.0, 0.0}
	i := 0
	producer.cf <- func(fin float64) float64 {
		if i < len(floats) {
			j := i
			i++
			return floats[j]
		}
		producer.cdone <- true
		return 0.0
	}
}
