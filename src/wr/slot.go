package wr

import (
	"fmt"
	"log"

	"github.com/kainkent69/wr/src/wr/recorder"
)

type Slots struct {
	Lists []*W
	Spins int64
	None  int64
	Total int64
	Track bool
	rand  Randomizor
	recorder.Record
	// init
	_init bool
}

func (s *Slots) Check() {
	if !s._init {
		log.Fatal("slot is not initialized")
	}
}

// spin and get result

// init
func (s *Slots) Init(rand Randomizor) {
	for _, v := range s.Lists {
		v.Init(s)
		s.Total += v.Weights
	}
	s.rand = rand
	s._init = true
	s.Record = recorder.Record{
		SReq: map[int64]int64{},
	}
	fmt.Printf("slot is now %t \n", s._init)
}

// spins  the slot
func (slot *Slots) Spin() W {
	rnd := slot.rand.Rand(slot.Total)
	last := int64(0)
	var selected W
	for _, v := range slot.Lists {
		start := last
		last += v.Weights
		if rnd <= last && start <= rnd {
			selected = *v
			if slot.Track {
				v.Hit()
			}
		} else if slot.Track {
			v.Unhit()
		}

	}

	if selected.IsEmpty {
		slot.Unhit()
	} else {
		slot.Hit()
	}
	return selected
}

// the random default
func DefaultSlot() Slots {
	slot := Slots{}
	slot.Init(Default)
	return slot
}
