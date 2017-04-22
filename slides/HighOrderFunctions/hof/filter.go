package hof

type ListFilterFunc func(int) bool

func (list ListOfInt) Filter(f ListFilterFunc) ListOfInt {
	var out ListOfInt
	for _, i := range list {
		if f(i) {
			out = append(out, i)
		}
	}
	return out
}
