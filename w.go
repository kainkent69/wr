package wr

type Wer interface {
	Info() *W
	Reward() int64
}
type W struct {
	// identifiyer it can mean a thin
	ID int64
	// its a weight that  is used  as the value
	Weights int64
	// if its supposed to be the empty value
	IsEmpty bool
	// the parent beause  the slots would init each of them
	Parent *Slots
	Record
	// initialize
	_init bool
}

// initiaalize
func (w *W) Init(slot *Slots) {
	w.Parent = slot
	w._init = true
	w.Record = Record{
		SReq: map[int64]int64{},
	}
}
