package main

import (
	"fmt"
	"time"
)

type streamFunc func(int) int

func main() {
	cout := make(chan int)
	cstop := make(chan bool)
	cfunc := make(chan streamFunc)
	go func() {
		for {
			cout <- 1
			time.Sleep(time.Millisecond * 10)
		}
	}()

	tenTimes := func(i int) int {
		return 1 * 10
	}

	fiveTimes := func(i int) int {
		return 1 * 5
	}

	var sf streamFunc = tenTimes

	go func() {
		for {

			select {
			case i := <-cout:
				fmt.Printf("%v ", sf(i))
			case sf = <-cfunc:
			case <-cstop:
				break
			}
		}
	}()

	time.Sleep(time.Second)

	cfunc <-fiveTimes

	time.Sleep(time.Second)

}
