package main

import (
	"fmt"
	"time"

	. "github.com/stefanhans/go-present/slides/FunctionalStreaming/functionalstreams"
)

func main() {
	node_in := NewNodeOfInt()

	aggr_map_1 := make(map[int]int)
	aggr_func := func(i int, aggr_map *map[int]int) { (*aggr_map)[i] = (*aggr_map)[i] + 1 }
	aggregator := NewAggregatorOfInt(&aggr_map_1, aggr_func)

	node_in.ConnectAggregator(aggregator)
	node_in.ProduceRandPositivAtMs(4, 10)

	time.Sleep(time.Second)
	fmt.Printf("%v\n", aggr_map_1)
	time.Sleep(time.Second)

	aggr_map_2 := make(map[int]int)
	aggr_map_1 = (*aggregator.Reset(&aggr_map_2))
	time.Sleep(time.Second)

	fmt.Printf("%v\n", aggr_map_1)
	fmt.Printf("%v\n", aggr_map_2)
}
