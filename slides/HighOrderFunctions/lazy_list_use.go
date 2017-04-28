package main

import (
	"fmt"
	"time"

	. "github.com/stefanhans/go-present/slides/HighOrderFunctions/hof"
)

func main() {
	var list LazyListOfFloat
	list.Monad = FloatMonad{0, func(x float64) float64 { return x + 2 }}

	ord := 20

	start := time.Now()
	fmt.Printf("list.Get(%v) yields %v\n", ord, list.Get(ord))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("list.Get(%v) yields %v\n", ord, list.Get(ord))
	fmt.Println(time.Since(start))
}

// END OMIT
