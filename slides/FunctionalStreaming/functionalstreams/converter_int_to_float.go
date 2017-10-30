package functionalstreams

// START_1 OMIT
type ConverterIntToFloat struct {
	in      chan int
	cin     chan chan int
	convert func(int) float64 // HL
	out     chan float64 // HL
	cout    chan chan float64 // HL
	close   chan bool
}
// END_1 OMIT

// START_2 OMIT
func (converter *ConverterIntToFloat) Start() {
	go func() {
		for {
			select {
			case converter.in = <-converter.cin: // HL
			case in := <-converter.in: converter.out <- converter.convert(in) // HL
			case converter.out = <-converter.cout: // HL
			case <-converter.close: return
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
	converter.out = make(chan float64) // HL
	converter.cout = make(chan chan float64) // HL
	converter.close = make(chan bool)
	converter.Start()
	return &converter
}
// END_3 OMIT

func (converter *ConverterIntToFloat) Stop() {
	converter.close <- true
}

// START_CONNECT OMIT
func (converter *ConverterIntToFloat) Connect(nextNode *NodeOfFloat) *NodeOfFloat { // HL
	converter.cout <- nextNode.in
	return nextNode
}
// END_CONNECT OMIT
