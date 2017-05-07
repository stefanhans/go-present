package main

import (
	"fmt"
	_ "time"
	"time"
)

func main() {
	// CLOSURESTART OMIT
	// type f func() // OMIT
	var function myf = func() { fmt.Println(" I was executed") }
	// CLOSURE END OMIT

	// func(f) {}(fi) // 3

	// How to execute fi?
	// func(x f) { x() }(fi) // 4

	// In the background
	// go func(x f) { x() }(fi) // 5

	time.Sleep(1 *time.Millisecond)
}
