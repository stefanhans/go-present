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

func MultiCompose(funcs ...func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		fout := funcs[0]
		for _, f := range funcs[1:] {
			fout = Compose(fout, f)
		}
		return fout(x)
	}
}

func main() {
	fmt.Println(MultiCompose(math.Sqrt, func(x float64) float64 { return x*x }, math.Sqrt)(4))
}
