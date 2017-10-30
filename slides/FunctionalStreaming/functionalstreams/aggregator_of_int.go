package functionalstreams

// TODO without pointer?

// START_1 OMIT
type AggregatorOfInt struct {
	in              chan int
	cin             chan chan int
	aggregator_map  *map[int]int                 // HL
	caggregator_map chan *map[int]int            // HL
	aggregate       func(int, *map[int]int)      // HL
	close           chan bool
}

// END_1 OMIT

// START_2 OMIT
func (aggregator *AggregatorOfInt) Start() {
	go func() {
		for {
			select {
			case in := <-aggregator.in: // HL
				aggregator.aggregate(in, aggregator.aggregator_map) // HL
			case aggregator.in = <-aggregator.cin: // HL
			case aggregator_map := <-aggregator.caggregator_map: // HL
				aggregator.caggregator_map <- aggregator.aggregator_map // HL
				aggregator.aggregator_map = aggregator_map              // HL
			case <-aggregator.close: return
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewAggregatorOfInt(aggr_map *map[int]int, f func(int, *map[int]int)) *AggregatorOfInt { // HL
	aggregator := AggregatorOfInt{}
	aggregator.in = make(chan int)
	aggregator.cin = make(chan chan int)
	aggregator.aggregator_map = aggr_map                       // HL
	aggregator.caggregator_map = make(chan *map[int]int)       // HL
	aggregator.aggregate = f                                   // HL
	aggregator.close = make(chan bool)
	aggregator.Start()
	return &aggregator
}

// END_3 OMIT

// START_RESET OMIT
func (aggregator *AggregatorOfInt) Reset(aggr_map *map[int]int) *map[int]int {
	aggregator.caggregator_map <- aggr_map
	return <-aggregator.caggregator_map
}

// END_RESET OMIT
