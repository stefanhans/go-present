package functionalstreams

import "fmt"

// START_1 OMIT
type AggregatorOfInt struct {
	in             chan int
	cin            chan chan int
	aggregator_map map[int]int           // HL
	out            chan map[int]int      // HL
	cout           chan chan map[int]int // HL
	aggregate      func(int)             // HL
	caggregate     chan func(int)        // HL
	flush          func(int) bool        // HL
	cflush         chan func(int) bool   // HL
	close          chan bool
}

// END_1 OMIT

// START_2 OMIT
func (aggregator *AggregatorOfInt) Start() {
	go func() {
		for {
			select {
			case aggregator.in = <-aggregator.cin: // HL
			case aggregator.out = <-aggregator.cout: // HL
			case in := <-aggregator.in: // HL
				aggregator.aggregate(in) // HL
				//fmt.Printf("Map %v\n", aggregator.aggregator_map)

				if aggregator.flush(in) { // HL
						aggregator.out <- aggregator.aggregator_map // HL
				} // HL
			case aggregator.flush = <-aggregator.cflush: // HL
			case in := <-aggregator.out: // HL
				fmt.Printf("%v ", in) // HL
			case <-aggregator.close: return
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewAggregatorOfInt() *AggregatorOfInt {
	aggregator := AggregatorOfInt{}
	aggregator.in = make(chan int)
	aggregator.cin = make(chan chan int)
	aggregator.aggregator_map = make(map[int]int)                                                          // HL
	aggregator.out = make(chan map[int]int, 2)                                                                // HL
	aggregator.cout = make(chan chan map[int]int)                                                          // HL
	aggregator.aggregate = func(i int) { aggregator.aggregator_map[i] = aggregator.aggregator_map[i] + 1 } // HL
	aggregator.caggregate = make(chan func(int))                                                           // HL
	aggregator.flush = func(in int) bool { return true }                                                   // HL
	aggregator.cflush = make(chan func(int) bool)                                                          // HL
	aggregator.close = make(chan bool)
	aggregator.Start()
	return &aggregator
}

// END_3 OMIT

func (aggregator *AggregatorOfInt) Print() {
	go func() {
		for {
			select {
			case in := <-aggregator.out: // HL
				fmt.Printf("%v ", in) // HL
			}
		}
	}()
}

// START_FLUSH OMIT
func (aggregator *AggregatorOfInt) Flush() {
	aggregator.cflush <- func(i int) bool {
		return true
	}
}

// END_FLUSH OMIT

// START_BUFFER OMIT
func (aggregator *AggregatorOfInt) Buffer() {
	aggregator.cflush <- func(i int) bool {
		return false
	}
}

// END_BUFFER OMIT

// START_FUNC OMIT
func (aggregator *AggregatorOfInt) SetFunc(f func(int) bool) {
	aggregator.cflush <- f
}
func (aggregator *AggregatorOfInt) Len() int { return len(aggregator.aggregator_map) }

// END_FUNC OMIT
