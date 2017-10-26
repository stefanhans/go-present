package main

import (
	"fmt"
	"time"
)

func main() {
	cint_1 := make(chan int)
	cint_2 := make(chan int)
	delay_receive := 500
	delay_send := 50
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(delay_receive))
			select {
			case in := <-cint_1:
				fmt.Printf("cint_1 %v\n", in)
			}
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(delay_receive))
			select {
			case in := <-cint_2:
				fmt.Printf("cint_2 %v\n", in)
			}
		}
	}()

	var i int
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(delay_send))
			cint_1 <- i
			i++
		}
	}()

	var j int
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(delay_send))
			cint_2 <- j
			j--
		}
	}()

	time.Sleep(time.Second)
	delay_receive = 1000
	fmt.Printf("delay_send delay_receive: %v %v\n", delay_send, delay_receive)
	time.Sleep(time.Second)
	//close(cint)
}
