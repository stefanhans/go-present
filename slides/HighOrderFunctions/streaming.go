package main

import (
	"fmt"
	"time"
)

type streamFunc func(int) int

var (
	tenTimes  streamFunc = func(i int) int { return 1 * 10 }
	fiveTimes streamFunc = func(i int) int { return 1 * 5 }
)

func main() {
	cout := make(chan int)
	cstop := make(chan bool)
	cfunc := make(chan streamFunc)

	go func() {
		for {
			cout <- 1
			time.Sleep(time.Millisecond * 100)
		}
	}()

	sf := tenTimes

	go func() {
		for {
			select {
			case i := <-cout:
				fmt.Printf("%v ", sf(i))
			case sf = <-cfunc:
				fmt.Printf("\nsf changed\n")
			case <-cstop:
				fmt.Println()
				break
			}
		}
	}()

	time.Sleep(time.Second)
	cfunc <- fiveTimes
	time.Sleep(time.Second)
}
