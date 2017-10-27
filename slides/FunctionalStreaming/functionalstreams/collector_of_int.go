package functionalstreams

// START_SUBSCRIPTION OMIT
type SubscriptionOfNodeOfInt struct {
	name string
	node NodeOfInt
}
// END_SUBSCRIPTION OMIT

// START_1 OMIT
type CollectorOfInt struct {
	in_map              map[string]*NodeOfInt    // HL
	cin_map_subscribe   chan *SubscriptionOfNodeOfInt // HL
	cin_map_unsubscribe chan string            // HL
	f                    func(int)				// HL
	cf                   chan func(int)			// HL
	in                  chan int
	out                  chan int
	cout                 chan chan int
	close                chan bool
}
// END_1 OMIT

// START_2 OMIT
func (collector *CollectorOfInt) Start() {
	go func() {
		for { select {
		case collector.out = <-collector.cout:
		case collector.f = <-collector.cf: // HL

		case subscription := <-collector.cin_map_subscribe: // HL
			collector.in_map[subscription.name] = &subscription.node // HL
			subscription.node.cin <- collector.in

		case name := <-collector.cin_map_unsubscribe: // HL
			delete(collector.in_map, name) // HL

		case <-collector.close: return
		}}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewCollectorOfInt() *CollectorOfInt {
	collector := CollectorOfInt{}
	collector.out = make(chan int)
	collector.cout = make(chan chan int)
	collector.in_map = make(map[string]*NodeOfInt)               // HL
	collector.cin_map_subscribe = make(chan *SubscriptionOfNodeOfInt) 		// HL
	collector.cin_map_unsubscribe = make(chan string)          // HL
	/*
	collector.f = func(in int) {                                // HL
		for _, cout := range collector.in_map { // HL
			go func(cout chan int, in int) { cout <- in }(cout, in) // HL
		} // HL
	} // HL
	*/
	collector.cf = make(chan func(int)) // HL
	collector.close = make(chan bool)
	collector.Start()
	return &collector
}

// END_3 OMIT

/*
// START_4 OMIT
func (node *NodeOfInt) ConnectCollector(collector *CollectorOfInt) *CollectorOfInt {
	node.cout <- collector.in
	return collector
}

func (collector *CollectorOfInt) SubscribeCollector(name string, nextNode *NodeOfInt) *NodeOfInt {
	collector.cin_map_subscribe <- SubscriptionOfInt{name, nextNode.in}
	return nextNode
}
func (collector *CollectorOfInt) UnsubscribeCollector(name string) {
	collector.cin_map_unsubscribe <- name
}
// END_4 OMIT
*/