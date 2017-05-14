package main

import (
	"fmt"
	"math"
)

func Compose(f, g func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func main() {
	fmt.Println(Compose(math.Sqrt, func(x float64) float64 { return x*x })(4.0))
}
