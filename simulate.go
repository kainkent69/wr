package wr

import (
	"math"

	"github.com/kainkent69/wr/record"
)

type Simulator struct {
	List  []Wer
	Spins int64
	Bet   int64
}

// make a simulation
func (s *Simulator) Run(rnd Randomizor) record.Report {
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

// make a report
func (s *Simulator) report(slot Slots) record.Report {
	report := record.Report{
		Each:         map[int64]record.Report{},
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
		childReport := record.Report{
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
