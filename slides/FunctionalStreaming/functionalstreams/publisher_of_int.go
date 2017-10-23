package functionalstreams

// START_SUBSCRIPTION OMIT
type SubscriptionOfInt struct {
	name string
	cint  chan int
}
// END_SUBSCRIPTION OMIT

// START_1 OMIT
type PublisherOfInt struct {
	in                   chan int
	cin                  chan chan int
	out_map              map[string]chan int    // HL
	cout_map_subscribe   chan SubscriptionOfInt // HL
	cout_map_unsubscribe chan string // HL
	close                chan bool
}
// END_1 OMIT

// START_2 OMIT
func (publisher *PublisherOfInt) Start() {
	go func() {
		for {
			select {
			case publisher.in = <-publisher.cin:
			case in := <-publisher.in: 										// HL
				for _, cout := range publisher.out_map { 					// HL
					cout <- in 												// HL
				} 															// HL
			case subscribtion := <-publisher.cout_map_subscribe:			// HL
				publisher.out_map[subscribtion.name] = subscribtion.cint		// HL
			case name := <-publisher.cout_map_unsubscribe:					// HL
				delete(publisher.out_map, name)								// HL
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
	publisher.out_map = make(map[string]chan int)                 // HL
	publisher.cout_map_subscribe = make(chan SubscriptionOfInt)   // HL
	publisher.cout_map_unsubscribe = make(chan string) // HL
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
