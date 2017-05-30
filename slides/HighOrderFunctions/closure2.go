package main

import (
	"fmt"
)

func main() {

	str := "Yes"

	// CLOSURESTART OMIT
	func() {
		fmt.Println("Can a closure see the outside? ", str)
	}()
	// CLOSURE END OMIT
}
