package functionalstreams

// START_RR OMIT
func (publisher *PublisherOfInt) DistributeRoundRobin() {
	publisher.cf <- func(in int) {
		if publisher.last_index == len(publisher.out_map) {
			publisher.last_index = 0
		}
		for i, subscription := range publisher.out_index {
			if i == publisher.last_index {
				go func(cout chan int, in int) { cout <- in }(subscription.cint, in)
				publisher.last_index++
				return
			}
		}
	}
}

// END_RR OMIT

// START_TOALL OMIT
func (publisher *PublisherOfInt) DistributeToAll() {
	publisher.cf <- func(in int) {
		for _, cout := range publisher.out_map {
			go func(cout chan int, in int) { cout <- in }(cout, in)
		}
	}
}
// END_TOALL OMIT

// START_FILTER OMIT
func (node *NodeOfInt) Filter(filter func(int) bool) *NodeOfInt {
	publisher := NewPublisherOfInt()
	node.ConnectPublisher(publisher)
	nextNode := NewNodeOfInt()
	publisher.SubscribePublisher("t", nextNode)

	publisher.cf <- func(in int) {
		if filter(in) {
			for _, cout := range publisher.out_map {
				go func(cout chan int, in int) { cout <- in }(cout, in)
				break
			}
		}
	}
	return nextNode
}
// END_FILTER OMIT