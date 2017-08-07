package main

import (
	"fmt"
	"strconv"
)

type List []interface{}

type Aggregator interface {
	Aggregate(interface{}, interface{}) interface{}
}

type Float64Converter interface {
	Float64Convert(interface{}) (float64, error)
}

type Monad struct {
	NeutralElement interface{}
	AssocFunc      func(interface{}, interface{}) interface{}
}

func (monad Monad) Aggregate(x, y interface{}) interface{} {
	fx := monad.NeutralElement
	if f, ok := monad.Float64Convert(x); ok {
		fx = f
	}
	fy := monad.NeutralElement
	if f, ok := monad.Float64Convert(y); ok {
		fy = f
	}
	return monad.AssocFunc(fx, fy)
}

func (monad Monad) Float64Convert(v interface{}) (float64, bool) {
	if i, ok := v.(int); ok {
		return float64(i), ok
	}
	if f, ok := v.(float64); ok {
		return f, ok
	}
	if s, ok := v.(string); ok {
		if is, err := strconv.Atoi(s); err == nil {
			return float64(is), ok
		}
		return 0.0, ok
	}
	return 0.0, false
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
	monad := Monad{0.0,
		func(m, n interface{}) interface{} {
			return m.(float64) + n.(float64)
		}}
	fmt.Printf("List %v . Fold(monad) yields %v\n",
		list, list.Fold(monad))
}
