package main

import (
	"fmt"
)

func main() {

	// CLOSURESTART OMIT
	for j := 0; j < 5; j++ {
		func() {
			fmt.Println("And here? ", j)
			j *= 10
		}()
	}
	// CLOSURE END OMIT
}
