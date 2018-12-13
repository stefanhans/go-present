package main

import (
	"fmt"
)

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func mult(x, y uint) uint {
	return x * y
}

func stop(x, y uint) bool {
	return x < y
}

func main() {
	in := uint(8)
	fmt.Println(factorial(in))

	end := in
	out := uint(2)
	for i := uint(2); i < end; i++ {
		fmt.Printf("%v*(%v-1)=%v*%v\n", out, i, out, (i + 1))
		out = out * (i + 1)
	}
	fmt.Println(out)

	out = uint(2)
	for i := uint(2); stop(i, end); i++ {
		fmt.Printf("%v*(%v-1)=%v*%v\n", out, i, out, (i + 1))
		out = mult(out, (i + 1))
	}
	fmt.Println(out)
}
