package wr_test

import (
	"testing"

	"github.com/kainkent69/wr"
)

type wl struct {
	*wr.W
	reward int64
}

func (w *wl) Info() *wr.W {
	return w.W
}

func (w *wl) Reward() int64 {
	return w.reward
}

func TestMain(t *testing.T) {
	list0 := []*wl{
		{
			W: &wr.W{
				ID:      1,
				Weights: 5000,
			},
			reward: 500,
		},
		{
			W: &wr.W{
				ID:      2,
				Weights: 5000,
			},
			reward: 500,
		},
		{
			W: &wr.W{
				ID:      3,
				Weights: 5000,
			},
			reward: 500,
		},
	}

	list1 := []*wl{
		{
			W: &wr.W{
				ID:      4,
				Weights: 10000,
				IsEmpty: true,
			},
			reward: 0,
		},
	}

	list := []wr.Wer{}
	for _, l := range list0 {
		list = append(list, l)
	}
	for _, l := range list1 {
		list = append(list, l)
	}

	// record the slots
	slot := wr.Slots{
		Lists: list,
		Track: true,
	}

	slot.Init(wr.Default)

	sm := wr.Simulator{
		List:  list,
		Spins: 100000,
		Bet:   100,
	}

	report := sm.Run(wr.Default)
	report.Print()

}
