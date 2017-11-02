package functionalstreams

// START_SUBSCRIPTION OMIT
type SubscriptionToInt struct {
	name string
	cint chan int
}

// END_SUBSCRIPTION OMIT

// START_1 OMIT
type DistributorOfInt struct {
	in                   chan int
	cin                  chan chan int
	f                    func(int)				 // Distribute over subscriptions // HL
	cf                   chan func(int)			// // HL
	out_map              map[string]chan int       // Handle subscriptions // HL
	cout_map_subscribe   chan SubscriptionToInt    // // HL
	cout_map_unsubscribe chan string               // // HL
	out_index            []SubscriptionToInt	   // Subscriptions ordered by number // HL
	last_index           int                       // // HL
	clast_index          chan int 	// OMIT
	close                chan bool 	// OMIT
}

// END_1 OMIT

// START_2 OMIT
func (distributor *DistributorOfInt) Start() {
	go func() {
		for { select {
			case distributor.in = <-distributor.cin:

			case in := <-distributor.in: distributor.f(in) // HL
			case distributor.f = <-distributor.cf: // HL

			case subscription := <-distributor.cout_map_subscribe: // HL
				distributor.out_map[subscription.name] = subscription.cint // HL
				distributor.out_index = append(distributor.out_index, subscription) // HL

			case name := <-distributor.cout_map_unsubscribe: // HL
				delete(distributor.out_map, name) // HL
				// delete from distributor.out_index accordingly // HL
				// ...
				i := -1; _ = i 	// OMIT
				for n, subscription := range distributor.out_index { 	// OMIT
					if subscription.name == name { i = n }} 	// OMIT
				if i != -1 { 	// OMIT
					distributor.out_index = append(distributor.out_index[:i], 	// OMIT
						distributor.out_index[i+1:]...)} 	// OMIT
			case <-distributor.close: return 	// OMIT
		}}
	}()
}

// END_2 OMIT

// START_3 OMIT
func NewDistributorOfInt() *DistributorOfInt {
	distributor := DistributorOfInt{}
	distributor.in = make(chan int)
	distributor.cin = make(chan chan int)
	distributor.out_map = make(map[string]chan int)               // HL
	distributor.cout_map_subscribe = make(chan SubscriptionToInt) // HL
	distributor.cout_map_unsubscribe = make(chan string)          // HL
	distributor.out_index = make([]SubscriptionToInt, 0)			// HL
	distributor.f = func(in int) {                                // HL
		for _, cout := range distributor.out_map { // HL
			go func(cout chan int, in int) { cout <- in }(cout, in) // HL
		} // HL
	} // HL
	distributor.cf = make(chan func(int)) // HL
	distributor.close = make(chan bool)
	distributor.Start()
	return &distributor
}

// END_3 OMIT

// START_4 OMIT
func (node *NodeOfInt) ConnectDistributor(distributor *DistributorOfInt) *DistributorOfInt {
	node.cout <- distributor.in
	return distributor
}

func (distributor *DistributorOfInt) SubscribeDistributor(name string, nextNode *NodeOfInt) {
	distributor.cout_map_subscribe <- SubscriptionToInt{name, nextNode.in}
}
func (distributor *DistributorOfInt) UnsubscribeDistributor(name string) {
	distributor.cout_map_unsubscribe <- name
}

// END_4 OMIT


// START_TEE OMIT
func (node *NodeOfInt) Tee() (*NodeOfInt, *NodeOfInt) {
	distributor := NewDistributorOfInt()
	node.ConnectDistributor(distributor)

	nodeA := NewNodeOfInt()
	nodeB := NewNodeOfInt()

	distributor.SubscribeDistributor("A", nodeA)
	distributor.SubscribeDistributor("B", nodeB)

	return nodeA, nodeB
}
// END_TEE OMIT

// START_RR OMIT
func (distributor *DistributorOfInt) DistributeRoundRobin() {
	distributor.cf <- func(in int) {
		if distributor.last_index == len(distributor.out_map) {        // Reset index
			distributor.last_index = 0
		}
		for i, subscription := range distributor.out_index {           // Loop until last index
			if i == distributor.last_index {
				go func(cout chan int, in int) { cout <- in }(subscription.cint, in)
				distributor.last_index++
				return
	}}}
}
// END_RR OMIT

// START_TOALL OMIT
func (distributor *DistributorOfInt) DistributeToAll() {
	distributor.cf <- func(in int) {
		for _, cout := range distributor.out_map {
			go func(cout chan int, in int) { cout <- in }(cout, in)
	}}
}
// END_TOALL OMIT

// START_FILTER OMIT
func (node *NodeOfInt) Filter(filter func(int) bool) *NodeOfInt {
	distributor := NewDistributorOfInt()

	distributor.cf <- func(in int) {
		if filter(in) {
			cout := distributor.out_map["t"]
			cout <- in
		}
	}

	nextNode := NewNodeOfInt()

	node.ConnectDistributor(distributor).SubscribeDistributor("t", nextNode)

	return nextNode
}
// END_FILTER OMIT

// START_SWITCH OMIT
func (node *NodeOfInt) Switch(fswitch func(int) bool) (nodeTrue, nodeFalse *NodeOfInt) {
	distributor := NewDistributorOfInt()
	node.ConnectDistributor(distributor)

	nextNodeTrue := NewNodeOfInt()
	nextNodeFalse := NewNodeOfInt()
	distributor.SubscribeDistributor("t", nextNodeTrue)
	distributor.SubscribeDistributor("f", nextNodeFalse)

	distributor.cf <- func(in int) {
		if fswitch(in) {
			coutTrue := distributor.out_map["t"]
			coutTrue <- in
		} else {
			coutFalse := distributor.out_map["f"]
			coutFalse <- in
		}
	}
	return nextNodeTrue, nextNodeFalse
}
// END_SWITCH OMIT
