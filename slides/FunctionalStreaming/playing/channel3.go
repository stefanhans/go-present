package main

import (
	"fmt"
	"time"
)

type Edge chan float64

type Connector interface {
	Connect(func(fin float64) float64) Node
}


type Producer struct {
	n float64
	cout Edge
	cdone chan bool
	f func(float64) (float64)
}

func NewProducer(f func(float64) float64) Producer {
	producer := Producer{}
	producer.cout = make(Edge)
	producer.cdone = make(chan bool)

	producer.f = f
	return producer
}

func (producer *Producer) Start() {
	go func() {
		for {
			select {
			case <-producer.cdone:
				fmt.Printf("\nreturn source's goroutine\n")
				return
			default:
				producer.cout <- producer.f(1.0)
			}
		}
	}()
}

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

func (oldNode Node) Connect(f func(fin float64) float64) Node {
	node := NewNode(f)
	node.Start(oldNode.cin)
	fmt.Printf("node started\n")
	return node
}

type Consumer struct {
	c        Edge
	f        func(Edge)
	buffer   []float64
	buffered bool
}

func main() {
	producer := NewProducer(func(fin float64) float64 {
		return 27.0
	})
	producer.Start()
	fmt.Printf("producer started\n")

	firstNode := NewNode(func(fin float64) float64 {
		return fin * 2
	})
	firstNode.Start(producer.cout)
	fmt.Printf("firstNode started\n")

	node := firstNode.Connect(func(fin float64) float64 { return fin * 3 })

	go func() {
		for {
			fmt.Printf("cs%v ", <-node.cout)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	fmt.Printf("consumer started\n")

	time.Sleep(time.Second)

	producer.cdone <- true

}
