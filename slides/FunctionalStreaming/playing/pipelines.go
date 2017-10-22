package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	genChan := gen(1, 2, 3)
	out := sq(genChan)
	for {
		i, ok := <-out
		if ! ok {
			break
		}
		fmt.Printf("%v\n", i)
	}
	for n := range sq(gen(1, 2, 3)) {
		fmt.Printf("%v\n", n)
	}
}
