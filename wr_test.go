package wr_test

import (
	"github.com/kainkent69/wr/src/wr"
	"github.com/kainkent69/wr/src/wr/simulate"
)

func main() {
	list := []*wr.W{
		{
			ID:      1,
			Weights: 5000,
		},
		{
			ID:      2,
			Weights: 5000,
		},

		{
			ID:      3,
			Weights: 5000,
		},

		{
			ID:      4,
			Weights: 10000,
			IsEmpty: true,
		},
	}

	// record the slots
	slot := wr.Slots{
		Lists: list,
		Track: true,
	}

	slot.Init(wr.Default)

	sm := simulate.Simulator{
		List:  list,
		Spins: 100000,
	}

	report := sm.Run(wr.Default)
	report.Print()

}
