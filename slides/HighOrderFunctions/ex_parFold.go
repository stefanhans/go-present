package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type ListOfInt []int

type IntMonad struct {
	NeutralElement int
	AssocFunc      func(int, int) int
}

func (list ListOfInt) Fold(monad IntMonad) int {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}

// START OMIT
func (list ListOfInt) chanFold(monad IntMonad, from, to int, end chan<- int) {
	out := monad.NeutralElement
	for i := from; i < to; i++ {
		out = monad.AssocFunc(out, list[i])
	}
	end <- out
}

func (list ListOfInt) ParFold(monad IntMonad, cores int) int {
	var from, to, out int
	end := make(chan int)
	batchSize := int(math.Ceil(float64(len(list)) / float64(cores)))
	for i := 0; i < cores; i++ {
		to = int(math.Min(float64(from+batchSize), float64(len(list))))
		go list.chanFold(monad, from, to, end)
		from = to
	}
	for i := 0; i < cores; i++ {
		out += <-end
	}
	return out
}

// END OMIT

func main() {
	var list ListOfInt
	j, k := 1, 50
	for i := j; i <= k; i++ {
		list = append(list, i)
	}

	monad := IntMonad{0, func(x, y int) int {
		time.Sleep(time.Duration(1 * time.Millisecond))
		return x + y
	}}

	start := time.Now()
	fmt.Printf("list[%v-%v].Fold(monad) yields %v\n", j, k, list.Fold(monad))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("list[%v-%v].ParFold(monad, %v) yields %v\n", j, k, runtime.NumCPU(),
		list.ParFold(monad, runtime.NumCPU()))
	fmt.Println(time.Since(start))
}
