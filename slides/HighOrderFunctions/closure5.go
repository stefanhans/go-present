package main

import (
	"fmt"
)

func main() {

	// CLOSURESTART OMIT
	for j := 0; j < 5; j++ {
		func(j int) {
			fmt.Println("OR here? ", j)
			j *= 10
		}(j)
	}
	// CLOSURE END OMIT
}
