package functionalstreams

// START_SUBSCRIPTION OMIT
type SubscriptionOfInt struct {
	name string
	cint chan int
}

// END_SUBSCRIPTION OMIT

// START_1 OMIT
type PublisherOfInt struct {
	in                   chan int
	cin                  chan chan int
	out_map              map[string]chan int    // HL
	cout_map_subscribe   chan SubscriptionOfInt // HL
	cout_map_unsubscribe chan string            // HL
	out_index            []SubscriptionOfInt	// HL
	f                    func(int)				// HL
	cf                   chan func(int)			// HL
	last_index           int
	clast_index          chan int
	close                chan bool
}

// END_1 OMIT

// START_2 OMIT
func (publisher *PublisherOfInt) Start() {
	go func() {
		for {
			select {
			case publisher.in = <-publisher.cin:
			case in := <-publisher.in: publisher.f(in) // HL
			case subscribtion := <-publisher.cout_map_subscribe: // HL
				publisher.out_map[subscribtion.name] = subscribtion.cint // HL
				publisher.out_index = append(publisher.out_index, subscribtion) // HL
			case name := <-publisher.cout_map_unsubscribe: // HL
				delete(publisher.out_map, name) // HL
				i := -1; _ = i
				for n, subscription := range publisher.out_index {
					if subscription.name == name { i = n }
				}
				publisher.out_index = append(publisher.out_index[:i], publisher.out_index[i+1:]...)
			case publisher.f = <-publisher.cf:
			case <-publisher.close:
				return
			}
		}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewPublisherOfInt() *PublisherOfInt {
	publisher := PublisherOfInt{}
	publisher.in = make(chan int)
	publisher.cin = make(chan chan int)
	publisher.out_map = make(map[string]chan int)               // HL
	publisher.cout_map_subscribe = make(chan SubscriptionOfInt) // HL
	publisher.cout_map_unsubscribe = make(chan string)          // HL
	publisher.f = func(in int) {                                // HL
		for _, cout := range publisher.out_map { // HL
			go func(cout chan int, in int) { cout <- in }(cout, in) // HL
		} // HL
	} // HL
	publisher.cf = make(chan func(int)) // HL
	publisher.close = make(chan bool)
	publisher.Start()
	return &publisher
}

// END_3 OMIT

// START_4 OMIT
func (node *NodeOfInt) ConnectPublisher(publisher *PublisherOfInt) *PublisherOfInt {
	node.cout <- publisher.in
	return publisher
}

func (publisher *PublisherOfInt) SubscribePublisher(name string, nextNode *NodeOfInt) *NodeOfInt {
	publisher.cout_map_subscribe <- SubscriptionOfInt{name, nextNode.in}
	return nextNode
}
func (publisher *PublisherOfInt) UnsubscribePublisher(name string) {
	publisher.cout_map_unsubscribe <- name
}

// END_4 OMIT

// START_TEE OMIT
func (node *NodeOfInt) Tee() (*NodeOfInt, *NodeOfInt) {
	publisher := NewPublisherOfInt()
	node.Produce().ConnectPublisher(publisher)

	nodeA := NewNodeOfInt()
	nodeB := NewNodeOfInt()

	publisher.SubscribePublisher("A", nodeA)
	publisher.SubscribePublisher("B", nodeB)

	return nodeA, nodeB
}

// END_TEE OMIT