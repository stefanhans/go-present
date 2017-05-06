package main

import (
	"fmt"
)

func main() {

	// CLOSURESTART OMIT
	for j := 0; j < 5; j++ {
		func() {
			fmt.Println("What happens here? ", j)
		}()
		j++
	}
	// CLOSURE END OMIT
}