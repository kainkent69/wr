package simulate

import (
	"github.com/kainkent69/wr/src/wr"
)

type Simulator struct {
	List  []*wr.W
	Spins int64
}

// make a simulation
func (s *Simulator) Run(rnd wr.Randomizor) Report {
	slot := &wr.Slots{
		Lists: s.List,
		Track: true,
	}
	slot.Init(rnd)
	run(slot, 100)

	return s.report(*slot)
}

// run the result
func run(slot *wr.Slots, spins int64) {
	for range spins {
		slot.Spin()
	}

}

type Report struct {
	HF   float64
	SAvg float64
	Each map[int64]Report
}

// make a report
func (s *Simulator) report(slot wr.Slots) Report {
	report := Report{
		Each: map[int64]Report{},
	}
	report.HF = slot.HF()
	report.SAvg = slot.Savg()
	for _, v := range slot.Lists {
		childReport := Report{
			HF:   v.HF(),
			SAvg: v.Savg(),
		}
		report.Each[v.ID] = childReport

	}

	return report
}
