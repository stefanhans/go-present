package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type ListOfInt []int

type listMapFunc func(int) int

func (list ListOfInt) Map(f listMapFunc) ListOfInt {
	var out ListOfInt
	for _, i := range list {
		out = append(out, f(i))
	}
	return out
}

func (list ListOfInt) chanMap(f listMapFunc, from, to int, end chan<- bool) {
	for i := from; i < to; i++ {
		list[i] = f(list[i])
	}
	end <- true
}

func (list ListOfInt) ParMap(f listMapFunc, cores int) ListOfInt {
	out := make(ListOfInt, len(list))
	copy(out, list)

	end := make(chan bool)
	var from, to int
	batchSize := int(math.Ceil(float64(len(out)) / float64(cores)))
	for i := 0; i < cores; i++ {
		to = int(math.Min(float64(from+batchSize), float64(len(out))))
		go (&out).chanMap(f, from, to, end)
		from = to
	}
	for i := 0; i < cores; i++ {
		<-end
	}
	return out
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
	fmt.Printf("list%v.Map(tenTimes) yields %v\n", list, list.Map(tenTimes))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("list%v: ParMap(tenTimes, %v) yields %v\n", list, runtime.NumCPU(),
		list.ParMap(tenTimes, runtime.NumCPU()))
	fmt.Println(time.Since(start))
	fmt.Printf("and list%v is immutable\n", list)
}

// END OMIT
