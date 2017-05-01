package main

import (
	"fmt"
)

func main() {

	// CLOSURESTART OMIT
	for j := 0; j < 5; j++ {
		func() {
			fmt.Println("Make a guess! ", j)
		}()
		j++
	}
	fmt.Println("Make a guess! ", j)
	// CLOSURE END OMIT
}
