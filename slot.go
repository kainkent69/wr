package wr

import (
	"fmt"
	"log"

	"github.com/kainkent69/wr/record"
)

type Slots struct {
	Lists []Wer
	Spins int64
	None  int64
	Total int64
	Track bool
	rand  Randomizor
	record.Record
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
		info := v.Info()
		info.Init(s)
		s.Total += info.Weights
	}
	s.rand = rand
	s._init = true
	s.Record = record.Record{
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
		info := v.Info()
		start := last
		last += info.Weights
		if rnd <= last && start <= rnd {
			selected = *info
			if slot.Track {
				info.Hit()
			}
		} else if slot.Track {
			v.Info()
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
