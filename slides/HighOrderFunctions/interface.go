package main

import (
	"fmt"
)

// START_IF OMIT
type List []interface{}

// END_IF OMIT

// START_IMPL OMIT
type Monad struct {
	NeutralElement interface{}
	AssocFunc      func(...interface{}) interface{}
}

// END_IMPL OMIT

// START_FOLD OMIT
func (list List) Fold(monad Monad) interface{} {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}

// END_FOLD OMIT

func main() {
	var listInt = List{-2, -1, 2, 3, 5}
	monadInt := Monad{1, func(monad ...interface{}) interface{} {
		return monad[0].(int) * monad[1].(int)
	}}
	fmt.Printf("List %v . Fold(listMonad) yields %v\n",
		listInt, listInt.Fold(monadInt))

	var listFloat = List{-2.5, -1.0, 2.0, 2.0, 3.0}
	monadFloat := Monad{0.0, func(monad ...interface{}) interface{} {
		return monad[0].(float64) + monad[1].(float64)
	}}
	fmt.Printf("List %v . Fold(monadFloat) yields %v\n",
		listFloat, listFloat.Fold(monadFloat))

	var listComplex = List{-2.5 + 3i, 1.5 + 2i}
	monadComplex := Monad{0 + 0i, func(monad ...interface{}) interface{} {
		return monad[0].(complex128) + monad[1].(complex128)
	}}
	fmt.Printf("List %v . Fold(monadComplex) yields %v\n",
		listComplex, listComplex.Fold(monadComplex))
}
