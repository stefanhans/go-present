package functionalstreams

// START_1 OMIT
type MonadOfInt struct {
	neutralElement int
	assocFunc      func(int, int) int
}
// END_1 OMIT


// START_2 OMIT
func NewMonadOfInt(neutralElement int, assocFunc func(int, int) int) *MonadOfInt { // HL
	monad := MonadOfInt{}
	monad.neutralElement = neutralElement                       // HL
	monad.assocFunc = assocFunc                       // HL
	return &monad
}
// END_2 OMIT
