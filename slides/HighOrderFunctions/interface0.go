package main

import (
	"fmt"
)

// START OMIT
type MonadType interface{}

type ListForMonad []MonadType

type Aggregator interface {
	GetNeutralElement() MonadType
	GetAssocFunc() func(...MonadType) MonadType
}

type Monad struct {
	NeutralElement MonadType
	AssocFunc      func(...MonadType) MonadType
}

func (monad Monad) GetNeutralElement() MonadType {
	return monad.NeutralElement
}

func (monad Monad) GetAssocFunc() func(...MonadType) MonadType {
	return monad.AssocFunc
}

func (list ListForMonad) Fold(monad Aggregator) MonadType {
	out := monad.GetNeutralElement()
	f := monad.GetAssocFunc()
	for _, i := range list {
		out = f(out, i)
	}
	return out
}

// END OMIT

func main() {

	var listInt = ListForMonad{-2, -1, 2, 3, 5}
	monadInt := Monad{1, func(monad ...MonadType) MonadType { return monad[0].(int) * monad[1].(int) }}
	fmt.Printf("List %v: Fold(listMonad) yields %v\n", listInt, listInt.Fold(monadInt))

	var listFloat = ListForMonad{-2.5, -1.0, 2.0, 2.0, 3.0}
	monadFloat := Monad{0.0, func(monad ...MonadType) MonadType { return monad[0].(float64) + monad[1].(float64) }}
	fmt.Printf("List %v: Fold(monadFloat) yields %v\n", listFloat, listFloat.Fold(monadFloat))

	var c1, nc complex128 = -2.5 + 3i, 0 + 0i
	var listComplex = ListForMonad{c1}
	monadComplex := Monad{nc, func(monad ...MonadType) MonadType { return monad[0].(complex128) + monad[1].(complex128) }}
	fmt.Printf("List %v: Fold(monadComplex) yields %v\n", listComplex, listComplex.Fold(monadComplex))

}
