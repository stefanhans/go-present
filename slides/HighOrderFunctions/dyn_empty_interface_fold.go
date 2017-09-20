package main

import (
	"fmt"
)

type List []interface{}

type Monad struct {
	NeutralElement interface{}
	AssocFunc      func(interface{}, interface{}) interface{}
}

// START_FOLD OMIT
func (list List) FoldToFloat(monad Monad) interface{} {
	out := monad.NeutralElement
	if i, ok := out.(int); ok {
		out = float64(i)
	}
	for _, v := range list {
		if i, ok := v.(int); ok {
			out = monad.AssocFunc(out, float64(i)).(float64)
		}
	}
	return out
}
// END_FOLD OMIT

func main() {
	list := List{-2.5, -1, 2, 2, 3}
	monad := Monad{0.0, func(x, y interface{}) interface{} {
		return x.(float64) + y.(float64)
	}}
	fmt.Printf("list%v.FoldToFloat(monad) yields %v\n",
		list, list.FoldToFloat(monad))
}
