package main

import (
	"fmt"
	"time"
)

func main() {
	threshold := 12000

	start := time.Now()
	sieve(threshold)
	fmt.Printf("Executiontime %s\n", time.Since(start))
}

// Daisy-chain generate and filter for every prime found until 'threshold' passed
func sieve(threshold int) {
	srcChan := make(chan int)

	go generate(srcChan)

	primeCount := 0

	for {
		prime := <- srcChan
		primeCount++

		//fmt.Printf("%d ", prime)
		if prime > threshold {
			fmt.Printf("%d is the %dth prime number resp. the first one over %d\n", prime, primeCount, threshold)
			return
		}
		dstChan := make(chan int)
		go filter(srcChan, dstChan, prime)
		srcChan = dstChan
	}
}

// Filter out integers divisible by 'prime' from stream ('src' to 'dst')
func filter(src <-chan int, dst chan<- int, prime int)  {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}


// Generate a stream of even positive integers (2, 4, 6 ...)
func generate(c chan<- int)  {
	for i:=2; ; i++ {
		c <- i
	}
}
