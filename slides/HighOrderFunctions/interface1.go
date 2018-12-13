package main

import (
	"fmt"
)

// START OMIT
type MonadType interface{}

type Monad struct {
	NeutralElement MonadType
	AssocFunc      func(MonadType, MonadType) MonadType
}

type ListForMonad []MonadType

func (list ListForMonad) Fold(monad Monad) MonadType {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}

// END OMIT

func main() {

	var listInt = ListForMonad{-2, -1, 2, 3, 5}
	monadInt := Monad{0, func(x, y MonadType) MonadType { return x.(int) + y.(int) }}
	fmt.Printf("List %v: Fold(listMonad) yields %v\n", listInt, listInt.Fold(monadInt))

	/*
		var listFloat = ListForMonad{-2.5, -1.0, 2, 2, 3}
		monadFloat := Monad{0.0, func(x, y MonadType) MonadType { return x.(float64) + y.(float64) }}
		fmt.Printf("List %v: Fold(monadFloat) yields %v\n", listFloat, listFloat.Fold(monadFloat))
	*/
}
