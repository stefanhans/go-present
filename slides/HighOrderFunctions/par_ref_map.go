package main

import (
	"fmt"
	"runtime"
	"time"
	"math"
)

type ListOfInt []int

type listMapFunc func(int) int

func (list ListOfInt) RefMap(f listMapFunc) {
	for i := 0; i < len(list); i++ {
		list[i] = f(list[i])
	}
}

func (list ListOfInt) chanMap(f listMapFunc, from, to int, c chan<- bool) {
	for i := from; i < to; i++ {
		(list)[i] = f((list)[i])
	}
	c<-true
}

func (list ListOfInt) ParMap(f listMapFunc, cores int) {
	var from, to int
	c := make(chan bool)
	batchSize := int(math.Ceil(float64(len(list)) / float64(cores)))
	for i := 0; i < cores; i++ {
		to = int(math.Min(float64(from+batchSize), float64(len(list))))
		go list.chanMap(f, from, to, c)
		from = to
	}
	for i := 0; i < cores; i++ { <-c }
}

func main() {
	tenTimes := func(x int) int {
		time.Sleep(time.Duration(1 * time.Millisecond))
		return x * 10
	}
	var list ListOfInt
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
	start := time.Now()
	fmt.Printf("%v.ParMap(tenTimes, %v) ", list, runtime.NumCPU())
	list.ParMap(tenTimes, runtime.NumCPU())
	fmt.Printf("yields %v\n", list)
	fmt.Print(time.Since(start))
}

// END OMIT