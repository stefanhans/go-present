package main

import (
	"fmt"
)

// START OMIT
type Account struct {
	Name   string
	Amount float64
}

type ListOfAccount []Account

func (list ListOfAccount) GroupBy(f func(float64, float64) float64) map[string]float64 {
	out := make(map[string]float64)
	for _, acc := range list {
		out[acc.Name] = f(acc.Amount, out[acc.Name])
	}
	return out
}

// END OMIT

func main() {
	list := ListOfAccount{
		Account{"Peter", 12.50},
		Account{"Peter", 2.25},
		Account{"Peter", -10.0},
		Account{"Susann", 12.75},
		Account{"Susann", -12.50},
	}

	balance := func(i float64, old float64) float64 {
		return old + i
	}

	fmt.Printf("list%v.GroupBy(balance) yields\n%v\n", list, list.GroupBy(balance))
}
