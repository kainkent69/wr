package simulate

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

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
	run(slot, s.Spins)
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
	Hit  int64
	Fail int64
	// streak result
	StreakResult map[int64]int64
	MinStreak    int64
	MaxStreak    int64
	Contirbution float64
	// the meaning for empty
	IsEmpty bool

	// reporter

	Each map[int64]Report
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
		totalHits := float64(slot.H)

		childReport := Report{
			HF:      v.HF(),
			Hit:     v.H,
			Fail:    v.F,
			IsEmpty: v.IsEmpty,
		}

		if !v.IsEmpty {
			res := float64(v.H) * 100 / totalHits
			if math.IsInf(res, 1) || math.IsNaN(res) {
				res = 0
			}
			childReport.Contirbution = res
		} else {
			res := float64(v.H) * 100 / float64(s.Spins)
			if math.IsInf(res, 1) || math.IsNaN(res) {
				res = 0
			}
			childReport.Contirbution = res
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
