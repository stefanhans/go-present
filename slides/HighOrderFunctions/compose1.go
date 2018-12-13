package main

import (
	"fmt"
)

func ComposeString(f, g func(str string) string) func(str string) string {
	return func(str string) string {
		return f(g(str))
	}
}

func main() {
	fmt.Println(ComposeString(
		func(str string) string { return str + " hello" },
		func(str string) string { return str + " Peter" })("my friend"))
}
