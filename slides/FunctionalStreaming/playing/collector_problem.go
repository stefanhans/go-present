package main

import (
	"fmt"
	"time"
)

func main() {
	cout_1 := make(chan int)

	go func() {
		for {
			cout_1 <- 1
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for {
			cout_1 <- 2
			time.Sleep(time.Millisecond)
		}
	}()

	cin := make(chan int)

	go func() {
		for {
			in := <-cout_1
			cin <- in
		}
	}()

	go func() {
		for {
			fmt.Printf("%v \n", <-cin)
		}
	}()
	time.Sleep(time.Second)

}
