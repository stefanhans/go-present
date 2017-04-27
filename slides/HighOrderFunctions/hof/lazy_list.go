package hof

type FloatMonad struct {
	NeutralElement float64
	AssocFunc      func(float64) float64
}

type LazyListOfFloat struct {
	Monad FloatMonad
	floats []float64
	last float64
}

func(list *LazyListOfFloat) Next() float64 {
	if list.floats == nil {
		list.last = list.Monad.NeutralElement
	} else {
		list.last = list.Monad.AssocFunc(list.last)
	}
	list.floats = append(list.floats, list.last)
	return list.last
}

func(list *LazyListOfFloat) Get(ord int) float64 {
	for i:=len(list.floats);  i<ord; i++ {
		list.Next()
	}
	return list.floats[ord-1]
}