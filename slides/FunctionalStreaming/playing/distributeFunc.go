package main

import (
	"fmt"
	"time"
)

type Edge chan float64

type Node struct {
	cins   []Edge
	couts  []Edge

	fdist  func()
	cfdist chan func()

	cdone  chan bool
}

func NewNode(f func()) []Edge {
	node := Node{}
	node.cfdist = make(chan func())
	node.cdone = make(chan bool)

	node.fdist = f
	return node.couts
}

func (node *Node) Start(cins []Edge) {
	go func() {
		for {
			select {
			case f
			}
		}
	}
	node.fdist()
}

func (producer *Node) Produce() {
	go func() {
		for {
			select {
			case <-producer.cdone:
				fmt.Printf("\nProduce() returns\n")
				return
			case producer.fcalc = <-producer.cfcalc:
				fmt.Printf("\nnew produce function arrived\n")
			default:
				producer.couts[0] <- producer.fcalc(1.0)
			}
		}
	}()
}

func (sourceNode *Node) Connect(node Node) *Node {
	node.Start(sourceNode.couts[0])
	fmt.Printf("node started\n")
	return &node
}

func (consumer *Node) Consume() {
	go func() {
		for {
			select {
			case fin := <-consumer.cins[0]:
				consumer.fcalc(fin)
			case consumer.fcalc = <-consumer.cfcalc:
				fmt.Printf("\nnew consumer function arrived\n")
			}
		}
	}()
}

func (sourceNode *Node) Distribute(fdist func()) []Edge {
	fdist()
	return sourceNode.couts
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

	producer.cdone <- true
}
