package main

import (
	"fmt"
	"time"
)

type Source interface {
	Next()
}

type SourceChan struct {
	n float64
	c chan float64
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

func Take(n int, scin *SourceChan) chan float64 {
	go func() {
		for i := 0; i < n; i++ {
			scin.Next()
			fmt.Printf("i: %v\n", i)
		}
	}()
	return scin.c
}

func Sink(s chan float64) {
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

	go Sink(Take(20, &sourceChan))



	time.Sleep(time.Second)

}
