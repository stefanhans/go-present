package main

import (
	"fmt"
	_ "time"
)

func main() {
	// CLOSURESTART OMIT
	f := func() { fmt.Println(" I was executed") }
	// f() OMIT
	// CLOSURE END OMIT

	// type f func()
	// var fi f = func() { fmt.Println(" I was executed") } // 2

	// CLOSURE END OMIT

	// func(f) {}(fi) // 3

	// How to execute fi?
	// func(x f) { x() }(fi) // 4

	// In the background
	// go func(x f) { x() }(fi) // 5

	// time.Sleep(1 *time.Millisecond) // LAST
}
