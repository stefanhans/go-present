package main

import (
	"fmt"
)

// COMPOSE START OMIT
func Compose(f, g func(x int) int) func(x int) int {
	return func(x int) int { return f(g(x)) }
}

func MultiCompose(funcs ...func(x int) int) func(x int) int {
	return func(x int) int {
		fout := funcs[0]
		for _, f := range funcs[1:] {
			fout = Compose(fout, f)
		}
		return fout(x)
	}
}
// COMPOSE END OMIT

func main() {
	fmt.Println(MultiCompose(
		func(x int) int { return x+1 },
		func(x int) int { return x*x },
		func(x int) int { return x+3 })(4))
}
