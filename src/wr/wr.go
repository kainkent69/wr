package wr

import (
	"log"
)

type Randomizor interface {
	Rand(n int64) int64
}

// for recording
type Slots struct {
	Lists []W
	Spins int64
	None  int64
	Total int64
	Track bool
	_init bool
	rand  Randomizor
	recorder.Recorder
}

func (s *Slots) Check() {
	if !s._init {
		log.Fatal("slot is not initialized")
	}
}

// spin and get result

// init
func (s *Slots) Init(rand Randomizor) {
	s.Check()
	for _, v := range s.Lists {
		v.Init(s)
		s.Total += v.Weights
	}
	s.rand = rand
	s._init = true
}

// spins  the slot
func (slot *Slots) Spin() W {
	rnd := slot.rand.Rand(slot.Total)
	last := int64(0)
	for _, v := range slot.Lists {
		start := last
		last += v.Weights
		if rnd <= last && start <= rnd {
			return v
		}
	}
	log.Fatal("weights isnt evenly distributed")
	return W{}
}

// the wiehghts
type W struct {
	// identifiyer it can mean a thin
	ID int64
	// its a weight that  is used  as the value
	Weights int64
	// if its supposed to be the empty value
	IsEmpty bool
	// the parent beause  the slots would init each of them
	Parent *Slots
	_init  bool
}

// the hit frequency
func (w W) HitFrequency(againts int64) {}

// initiaalize
func (w *W) Init(slot *Slots) {
	w.Parent = slot
	w._init = true
}
