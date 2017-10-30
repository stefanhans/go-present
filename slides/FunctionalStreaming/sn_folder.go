package main

import (
	"fmt"
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_in := NewNodeOfInt()
	var i int
	node_in.SetFunc(func(in int) int { i++; return in + i })

	monad := NewMonadOfInt(0, func(result int, in int) int { return result + in })
	folder := NewFolderOfInt(monad)
	node_in.ConnectFolder(folder)
	node_in.ProduceAtMs(200)

	time.Sleep(time.Second); fmt.Printf("%v\n", folder.Result())
	time.Sleep(time.Second); fmt.Printf("%v\n", folder.Result())
}
