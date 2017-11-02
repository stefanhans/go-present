package functionalstreams

// START_1 OMIT
type ConverterIntToFloat struct {
	in      chan int
	cin     chan chan int
	convert func(int) float64 // HL
	out     chan float64
	cout    chan chan float64
	close   chan bool // OMIT
}
// END_1 OMIT

// START_2 OMIT
func (converter *ConverterIntToFloat) Start() {
	go func() {
		for {
			select {
			case converter.in = <-converter.cin:
			case in := <-converter.in: converter.out <- converter.convert(in) // HL
			case converter.out = <-converter.cout:
			case <-converter.close: return // OMIT
			}
		}
	}()
}
// END_2 OMIT

// START_3 OMIT
func NewConverterIntToFloat() *ConverterIntToFloat {
	converter := ConverterIntToFloat{}
	converter.in = make(chan int)
	converter.cin = make(chan chan int)
	converter.convert = func(in int) float64 { return float64(in) } // HL
	converter.out = make(chan float64)
	converter.cout = make(chan chan float64)
	converter.close = make(chan bool) // OMIT
	converter.Start()
	return &converter
}
// END_3 OMIT

func (converter *ConverterIntToFloat) Stop() {
	converter.close <- true
}

// START_CONNECT OMIT
func (converter *ConverterIntToFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat {
	converter.cout <- nextNode.in
	return nextNode
}
// END_CONNECT OMIT
