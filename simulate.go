package wr

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type Simulator struct {
	List  []Wer
	Spins int64
	Bet   int64
}

// make a simulation
func (s *Simulator) Run(rnd Randomizor) Report {
	slot := &Slots{
		Lists: s.List,
		Track: true,
	}
	slot.Init(rnd)
	run(slot, s.Spins)
	return s.report(*slot)
}

// run the result
func run(slot *Slots, spins int64) {
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
	Each       map[int64]Report
	Spent      int64
	Won        int64
	RTP        float64
	RTPContrib float64
}

// make a report
func (s *Simulator) report(slot Slots) Report {
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
	report.Spent = s.Bet * s.Spins
	for _, d := range slot.Lists {
		v := d.Info()
		totalHits := float64(slot.H)
		var won int64 = d.Reward() * v.H
		report.Won += won
		childReport := Report{
			HF:         v.HF(),
			Hit:        v.H,
			Fail:       v.F,
			IsEmpty:    v.IsEmpty,
			Won:        won,
			RTPContrib: float64(won) * 100 / float64(report.Spent),
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

	// the rtp
	report.RTP = float64(report.Won) * 100 / float64(report.Spent)
	report.RTPContrib = float64(report.Won) * 100 / float64(report.Spent)

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
