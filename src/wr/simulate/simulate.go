package simulate

import (
	"encoding/json"
	"fmt"
	"log"

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
	Hit  int64
	Fail int64
	// streak result
	StreakResult map[int64]int64
	MinStreak    int64
	MaxStreak    int64
}

// make a report
func (s *Simulator) report(slot wr.Slots) Report {
	report := Report{
		Each:         map[int64]Report{},
		Hit:          slot.H,
		Fail:         slot.F,
		StreakResult: slot.SReq,
		MinStreak:    slot.SMin,
		MaxStreak:    slot.SMax,
		HF:           slot.HF(),
		SAvg:         slot.Savg(),
	}
	report.HF = slot.HF()
	report.SAvg = slot.Savg()
	for _, v := range slot.Lists {
		childReport := Report{
			HF:           v.HF(),
			SAvg:         v.Savg(),
			Hit:          v.H,
			Fail:         v.F,
			StreakResult: v.SReq,
			MinStreak:    v.SMin,
			MaxStreak:    v.SMax,
		}
		report.Each[v.ID] = childReport

	}

	return report
}

// return to something printable
func (r Report) Printable() string {
	b, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		log.Fatalf("report marshal to json failed")
	}
	return string(b)

}

// print it to stdin
func (r Report) Print() {
	fmt.Println(r.Printable())
}
