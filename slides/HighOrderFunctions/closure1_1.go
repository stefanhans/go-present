package main

import (
	"fmt"
	_ "time"
)

func main() {

	fmt.Println("func() is a function type\n" +
		"func(){} is a function literal which represents an anonymous function\n" +
		"func(){}() is a used anonymous function or a used evaluated function literal")

	// CLOSURESTART OMIT
	func()

	func() {}
	// CLOSURE END OMIT

	// func() {}() // Fixed
}
