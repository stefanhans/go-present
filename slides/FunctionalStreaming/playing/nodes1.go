package main

import (
	"fmt"
	"time"
)

type Source interface {
	Next()
}

type Edge chan float64

type SourceChan struct {
	n float64
	c Edge
	f func(float64) (float64, bool)
}

func (sc SourceChan) Next() {
	i, ok := sc.f(sc.n)
	if !ok {
		fmt.Printf("close(sc.c)\n")
		close(sc.c)
		return
	}
	sc.c <- i
}

func (scin *SourceChan) Take(n int) Edge {
	go func() {
		for i := 0; i < n; i++ {
			scin.Next()
			fmt.Printf("i: %v\n", i)
		}
	}()
	return scin.c
}


func (scin *SourceChan) TakeAll(msec int64) Edge {
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

func Sink(s Edge) {
	for n := range s {
		fmt.Printf("%v\n", n)
	}
}

func main() {

	sourceChan := SourceChan{}
	sourceChan.n = 0.0
	sourceChan.c = make(chan float64)
	sourceChan.f = func(n float64) (float64, bool) {
		sourceChan.n = n + 1
		return sourceChan.n, true
	}

	//fmt.Printf("%v\n", TenTimes(sourceChan.TakeAll(20)))
	Sink(sourceChan.TakeAll(20).TenTimes())

	//go Sink(TenTimes(sourceChan.TakeAll(20)))



	//time.Sleep(time.Second)

}
