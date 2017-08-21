package main

import (
	"fmt"
)

// START OMIT
type ListOfString []string

type stringGroupFunc func(s string) (string, int)

func (list ListOfString) Group(f func(str string) (string, int)) map[string]int {
	out := make(map[string]int)
	for _, s := range list {
		s, n := f(s)
		out[s] = out[s]+n
	}
	return out
}
// END OMIT

func main() {
	list := ListOfString{"and", "but", "or", "or", "but", "or"}
	count := func(s string) (string, int) {
		return s, 1
	}
	fmt.Printf("list%v.Group(count) yields %v\n", list, list.Group(count))
}
