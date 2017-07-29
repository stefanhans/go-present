package main

import (
	"fmt"
)

type List []interface{}

type Monad struct {
	NeutralElement interface{}
	AssocFunc      func(...interface{}) interface{}
}

// START_FOLD OMIT
func (list List) Fold(monad Monad) interface{} {
	out := monad.NeutralElement
	if i, ok := out.(int); ok {
		out = float64(i)
	}
	for _, v := range list {
		if i, ok := v.(int); ok {
			out = monad.AssocFunc(out, float64(i)).(float64)
		}
		if f, ok := v.(float64); ok {
			out = monad.AssocFunc(out, f)
		}
	}
	return out
}
// END_FOLD OMIT

func main() {
	var listFloat = List{-2.5, -1, 2.0, 2, 3.0}
	monadFloat := Monad{0.0, func(monad ...interface{}) interface{} {
		return monad[0].(float64) + monad[1].(float64)
	}}
	fmt.Printf("List %v . Fold(monadFloat) yields %v\n",
		listFloat, listFloat.Fold(monadFloat))
}
