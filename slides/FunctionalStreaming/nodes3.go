package main

import (
	"fmt"
	"time"
)

type Source interface {
	Next()
}

type Edge chan float64

type Producer struct {
	n float64
	c Edge
	f func(float64) (float64, bool)
}

type Consumer struct {
	c        Edge
	f        func(Edge)
	buffer   []float64
	buffered bool
}

func (sc Producer) Next() {
	i, ok := sc.f(sc.n)
	if !ok {
		fmt.Printf("close(sc.c)\n")
		close(sc.c)
		return
	}
	sc.c <- i
}

func (scin *Producer) Take(n int) Edge {
	go func() {
		for i := 0; i < n; i++ {
			scin.Next()
			fmt.Printf("i: %v\n", i)
		}
	}()
	return scin.c
}

func (scin *Producer) TakeAll(msec int64) Edge {
	go func() {
		for {
			scin.Next()
			time.Sleep(time.Millisecond * time.Duration(msec))
		}
	}()
	return scin.c
}

func (cin Edge) TenTimes() Edge {
	cout := make(chan float64)
	go func() {
		for {
			x := <-cin
			x *= 10
			cout <- x
		}
	}()

	return cout
}

func (s Edge) Sink(limit int) {
	if limit <= 0 {
		for n := range s {
			fmt.Printf("%v\n", n)
		}
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%v\n", <-s)
	}
}

func main() {

	producer := Producer{}
	producer.n = 0.0
	producer.c = make(chan float64)
	producer.f = func(n float64) (float64, bool) {
		producer.n = n + 1
		return producer.n, true
	}

	consumer := Consumer{}
	consumer.buffered = false
	consumer.f = func(e Edge) {
		for n := range e {
			fmt.Printf("%v by Consumer\n", n)
		}
	}
	//fmt.Printf("%v\n", TenTimes(sourceChan.TakeAll(20)))
	//producer.TakeAll(0).TenTimes().Sink(10)
	go consumer.f(producer.TakeAll(100).TenTimes())

	time.Sleep(time.Second)


}
