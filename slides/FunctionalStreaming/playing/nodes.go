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

func (scin *SourceChan) Take(n int) chan float64 {
	go func() {
		for i := 0; i < n; i++ {
			scin.Next()
			fmt.Printf("i: %v\n", i)
		}
	}()
	return scin.c
}


func (scin *SourceChan) TakeAll(msec int64) chan float64 {
	go func() {
		for {
			scin.Next()
			time.Sleep(time.Millisecond * time.Duration(msec))
		}
	}()
	return scin.c
}

func TenTimes(cin chan float64) (cout chan float64) {
	go func() {
		for {
			x := <-cin
			fmt.Printf("x1 %v\n", x)
			x *= 10
			fmt.Printf("x2 %v\n", x)
			cout <- x
		}
	}()

	return cout
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

	go Sink(TenTimes(sourceChan.TakeAll(20)))



	time.Sleep(time.Second)

}
