package functionalstreams

// START_TEE OMIT
func (node *NodeOfInt) Tee() (*NodeOfInt, *NodeOfInt) {
	publisher := NewPublisherOfInt()
	node.ConnectPublisher(publisher)

	nodeA := NewNodeOfInt()
	nodeB := NewNodeOfInt()

	publisher.SubscribePublisher("A", nodeA)
	publisher.SubscribePublisher("B", nodeB)

	return nodeA, nodeB
}
// END_TEE OMIT

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

	publisher.cf <- func(in int) {
		if filter(in) {
			cout := publisher.out_map["t"]
			cout <- in
		}
	}

	nextNode := NewNodeOfInt()

	node.ConnectPublisher(publisher).SubscribePublisher("t", nextNode)

	return nextNode
}
// END_FILTER OMIT

// START_SWITCH OMIT
func (node *NodeOfInt) Switch(fswitch func(int) bool) (nodeTrue, nodeFalse *NodeOfInt) {
	publisher := NewPublisherOfInt()
	node.ConnectPublisher(publisher)

	nextNodeTrue := NewNodeOfInt()
	nextNodeFalse := NewNodeOfInt()
	publisher.SubscribePublisher("t", nextNodeTrue)
	publisher.SubscribePublisher("f", nextNodeFalse)

	publisher.cf <- func(in int) {
		if fswitch(in) {
			coutTrue := publisher.out_map["t"]
			coutTrue <- in
		} else {
			coutFalse := publisher.out_map["f"]
			coutFalse <- in
		}
	}
	return nextNodeTrue, nextNodeFalse
}
// END_SWITCH OMIT
