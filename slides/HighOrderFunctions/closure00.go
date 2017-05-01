package main

import (
	"fmt"
	_ "time"
)

func main() {

	fmt.Println("'func() {} ()' is the simpliest closure or anonymous function.")

	// CLOSURESTART OMIT
	func() {}
	// CLOSURE END OMIT

	// func() {}() // Fixed
}
