package main

import (
	"fmt"
)

// START OMIT
type ListOfString []string

type MapStringCount map[string]int

type stringGroupFunc func(s string) (string, int)

func (list ListOfString) Group(f func(str string) (string, int)) MapStringCount {
	var out MapStringCount = make(map[string]int)

	for _, s := range list {
		s, n := f(s)
		out[s] = out[s]+n
	}
	return out
}
// END OMIT

func main() {
	var list = ListOfString{"and", "but", "or", "or", "but", "or"}

	count := func(s string) (string, int) {
		return s, 1
	}
	fmt.Printf("list%v: \nGroup(count) returns list%v\n",
		list, list.Group(count))
}
