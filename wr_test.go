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
	var list []wr.Wer = []wr.Wer{
		&wl{
			W: &wr.W{
				ID:      1,
				Weights: 14000,
				IsEmpty: true,
			},
			reward: 0,
		},
		&wl{
			W: &wr.W{
				ID:      2,
				Weights: 2000,
			},
			reward: 300,
		},
		&wl{
			W: &wr.W{
				ID:      3,
				Weights: 1400,
			},
			reward: 500,
		},
		&wl{
			W: &wr.W{
				ID:      4,
				Weights: 600,
			},
			reward: 750,
		},
	}

	// record the slots
	slot := wr.Slots{
		Lists: list,
		Track: true,
	}

	slot.Init(wr.Default)

	report := slot.Simulate(100, 5000000)
	report.Print()

}
