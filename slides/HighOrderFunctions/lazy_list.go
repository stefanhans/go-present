package main

import (
	"fmt"
	"time"
)

type FloatDef struct {
	StartElement float64
	Func         func(float64) float64
}

type LazyListOfFloat struct {
	Def    FloatDef
	floats []float64
	last   float64
}

func (list *LazyListOfFloat) Next() float64 {
	if list.floats == nil {
		list.last = list.Def.StartElement
	} else {
		list.last = list.Def.Func(list.last)
	}
	list.floats = append(list.floats, list.last)
	return list.last
}

func (list *LazyListOfFloat) Get(ord int) float64 {
	for i := len(list.floats); i < ord; i++ {
		list.Next()
	}
	return list.floats[ord-1]
}

func main() {
	var list LazyListOfFloat
	list.Def = FloatDef{0, func(x float64) float64 { return x + 2 }}

	ord := 20

	start := time.Now()
	fmt.Printf("list.Get(%v) yields %v\n", ord, list.Get(ord))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("list.Get(%v) yields %v\n", ord, list.Get(ord))
	fmt.Println(time.Since(start))
}

// END OMIT
