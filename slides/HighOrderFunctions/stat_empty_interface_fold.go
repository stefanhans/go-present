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
	AssocFunc      func(interface{}, interface{}) interface{}
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
	listInt := List{-2, -1, 2, 3, 5}
	monadInt := Monad{1, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	}}
	fmt.Printf("listInt%v.Fold(monadInt) yields %v\n",
		listInt, listInt.Fold(monadInt))

	listFloat := List{-2.5, -1.0, 2.0, 2.0, 3.0}
	monadFloat := Monad{0.0, func(x, y interface{}) interface{} {
		return x.(float64) + y.(float64)
	}}
	fmt.Printf("listFloat%v.Fold(monadFloat) yields %v\n",
		listFloat, listFloat.Fold(monadFloat))
}

/*


	listComplex := List{-2.5 + 3i, 1.5 + 2i}
	monadComplex := Monad{0 + 0i, func(x, y interface{}) interface{} {
		return x.(complex128) + y.(complex128)
	}}
	fmt.Printf("listComplex%v.Fold(monadComplex) yields %v\n",
		listComplex, listComplex.Fold(monadComplex))

*/
