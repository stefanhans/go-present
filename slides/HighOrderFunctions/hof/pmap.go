package hof

import "math"

func (list *ListOfInt) RefMap(f listMapFunc) {
	for i := 0; i < len(*list); i++ {
		(*list)[i] = f((*list)[i])
	}
}

func (list *ListOfInt) chanRefMap(f listMapFunc, from, to int, c chan<- bool) {
	for i := from; i < to; i++ {
		(*list)[i] = f((*list)[i])
	}
	c<-true
}

func (list *ListOfInt) ParRefMap(f listMapFunc, cores int) {
	var from, to int
	c := make(chan bool)
	batchSize := int(math.Ceil(float64(len(*list)) / float64(cores)))
	for i := 0; i < cores; i++ {
		to = int(math.Min(float64(from+batchSize), float64(len(*list))))
		go list.chanRefMap(f, from, to, c)
		from = to
	}
	for i := 0; i < cores; i++ {
		<-c
	}
}

func (list ListOfInt) PMap(f listMapFunc, cores int) ListOfInt {
	out := make(ListOfInt, len(list))
	copy(out, list)

	c := make(chan bool)
	var from, to int
	batchSize := int(math.Ceil(float64(len(out)) / float64(cores)))
	for i := 0; i < cores; i++ {
		to = int(math.Min(float64(from+batchSize), float64(len(out))))
		go (&out).chanRefMap(f, from, to, c)
		from = to
	}
	for i := 0; i < cores; i++ { <-c }
	return out
}