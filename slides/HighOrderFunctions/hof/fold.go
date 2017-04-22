package hof

type ListFoldMonad struct {
	NeutralElement int
	AssocFunc      func(int, int) int
}

func (list ListOfInt) Fold(monad ListFoldMonad) int {
	out := monad.NeutralElement
	for _, i := range list {
		out = monad.AssocFunc(out, i)
	}
	return out
}
