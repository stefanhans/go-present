package main

import (
	"fmt"
)

type Aggregate interface {
	GetMonadType() MonadType
	GetNeutralElement() MonadType
	GetAssocFunc() func(...MonadType) MonadType
}

type ListOfInt []int
type ListOfFloat []float64

type MonadType interface{}

type IntAggregate struct {
	Type           MonadType
	NeutralElement MonadType
	AssocFunc      func(MonadType, MonadType) MonadType
}

// START OMIT
func (list ListOfInt) Fold(monad IntAggregate) MonadType {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}
func (list ListOfFloat) Fold(monad IntAggregate) MonadType {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}

// END OMIT

func main() {

	var list = ListOfInt{-2, -1, 2, 2, 3}
	monad := IntAggregate{0, 0, func(x, y MonadType) MonadType { return x.(int) + y.(int) }}
	fmt.Printf("List %v: Fold(monad) yields %v\n", list, list.Fold(monad))

	var listFloat = ListOfFloat{-2.5, -1.0, 2, 2, 3}
	monadFloat := IntAggregate{0.0, 0.0, func(x, y MonadType) MonadType { return x.(float64) + y.(float64) }}
	fmt.Printf("List %v: Fold(monad) yields %v\n", listFloat, listFloat.Fold(monadFloat))
}

// END OMIT
