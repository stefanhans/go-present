package main

import (
	"fmt"
	"strconv"
)

type List []interface{}

type Aggregator interface {
	Aggregate(interface{}, interface{}) interface{}
}

type Monad struct {
	NeutralElement interface{}
	AssocFunc      func(interface{}, interface{}) interface{}
	Convert        (func(interface{}) interface{})
}

func (monad Monad) Aggregate(x, y interface{}) interface{} {
	return monad.AssocFunc(monad.Convert(x), monad.Convert(y))
}

// START_FOLD OMIT
func (list List) Fold(monad Monad) interface{} {
	out := monad.NeutralElement
	for _, v := range list {
		out = monad.Aggregate(out, v)
	}
	return out
}

// END_FOLD OMIT

func main() {

	var list = List{-2.5, -1, 2.0, 2, 3.0, "100", "werd"}
	monadFloat := Monad{0.0, func(m, n interface{}) interface{} {
		return m.(float64) + n.(float64)
	}, func(v interface{}) interface{} {

		if i, ok := v.(int); ok {
			return float64(i)
		}
		if f, ok := v.(float64); ok {
			return f
		}
		if s, ok := v.(string); ok {
			if is, err := strconv.Atoi(s); err == nil {
				return float64(is)
			}
			return 0.0
		}
		return 0.0
	}}
	fmt.Printf("List %v . Fold(monadFloat) yields %v\n",
		list, list.Fold(monadFloat))
}
