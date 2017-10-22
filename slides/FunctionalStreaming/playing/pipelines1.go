package main

import (
	"fmt"
)

type Seq interface {
	Next() (x float64, ok bool)
}

type SeqFunc func() (x float64, ok bool)

func (sf SeqFunc) Next() (x float64, ok bool) {
	return sf()
}

func Naturals() Seq {
	n := 0.0
	return SeqFunc(func() (x float64, ok bool) {
		n++
		return n, true
	})
}

func Take(n int, s Seq) Seq {
	return SeqFunc(func() (x float64, ok bool) {
		if n <= 0 {
			return 0, false
		}
		n--
		return s.Next()
	})
}

func Sink(s Seq) chan float64 {
	out := make(chan float64)
	go func() {
		for {
			i, ok := s.Next()
			if !ok {
				break
			}
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {

	taken := Take(10, Naturals())
	n, ok := taken.Next()
	fmt.Printf("%v %v\n", n, ok)
	fmt.Printf("%v\n", <-Sink(taken))

	sq := func(in <-chan float64) <-chan float64 {
		out := make(chan float64)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	genChan := Sink(Take(10, Naturals()))
	out := sq(genChan)
	for {
		i, ok := <-out
		if !ok {
			break
		}
		fmt.Printf("%v\n", i)
	}
	for n := range Sink(Take(10, Naturals())) {
		fmt.Printf("%v\n", n)
	}
}
