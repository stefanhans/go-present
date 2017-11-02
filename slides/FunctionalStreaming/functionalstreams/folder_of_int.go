package functionalstreams

// START_1 OMIT
type FolderOfInt struct {
	in      chan int
	cin     chan chan int
	monad   *MonadOfInt      // HL
	cmonad  chan MonadOfInt // HL
	result  int             // HL
	cresult chan int // HL
	close   chan bool // OMIT
}

// END_1 OMIT

// START_2 OMIT
func (folder *FolderOfInt) Start() {
	folder.result = folder.monad.neutralElement // HL
	go func() { for { select {
			case in := <-folder.in: // HL
				folder.result = folder.monad.assocFunc(folder.result, in) // HL
			case folder.in = <-folder.cin: // HL
			case <-folder.cresult: folder.cresult <- folder.result // HL
			case <-folder.close: return // OMIT
	}}}()
}

// END_2 OMIT

// START_3 OMIT
func NewFolderOfInt(monad *MonadOfInt) *FolderOfInt { // HL
	folder := FolderOfInt{}
	folder.in = make(chan int)
	folder.cin = make(chan chan int)
	folder.monad = monad                  // HL
	folder.cmonad = make(chan MonadOfInt) // HL
	folder.cresult = make(chan int) // HL
	folder.close = make(chan bool) // OMIT
	folder.Start()
	return &folder
}

// END_3 OMIT

// START_RESULT OMIT
func (folder *FolderOfInt) Result() int {
	folder.cresult <-1
	return <-folder.cresult
}
// END_RESULT OMIT
