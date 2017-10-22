package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)
	//c5 := make(chan string)
	//c6 := make(chan string)*/

	go func() {
		c1<- "hello "
	}()

	go func() {
		c2<- "you"
		c3<- "xxx"+ <-c2
	}()
	go func() {
		c2<- ":-)"
	}()
	s1, s2 := <-c1, <-c2
	fmt.Printf("s1: %v", s1)
	fmt.Printf("s2: %v", s2)

	s3 := <-c3
	fmt.Printf("s3: %v", s3)

	go func() {
		c4<- "hallo "

	}()

	s4 := fmt.Sprint(<-c4)
	fmt.Printf("s4: %v", s4)


	i1 := make(chan int)
	i2 := make(chan int)
	go func() {
		fmt.Printf("%v\n", <-i2 )
	}()
	fmt.Printf("%v\n", <-i2<- <-i1<-2)

}