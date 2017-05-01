package main

import (
	"fmt"
)

func main() {

	str := "Yes"

	// CLOSURESTART OMIT
	func() { fmt.Println("Can a closures see the outside? ", str) }()
	// CLOSURE END OMIT
}
