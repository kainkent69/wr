package wr

import (
	"math"

	"github.com/kainkent69/wr/record"
)

// simulator will try to run  the algortighim for  `spins` amount of times with `bet`
// used interface to be shared with wighted or ranges algorithim based
type Simulator interface {
	Simulate(bet int64, spins int64) record.Report
}

// run the result
func run(slot *Slots, spins int64) {
	for range spins {
		slot.Spin()
	}

}

// make a report
func (slot *Slots) Simulate(bet int64, spins int64) record.Report {
	run(slot, spins)
	report := record.Report{
		Bet:          bet,
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
	report.Spent = bet * spins
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
			childReport.Contribution = res
		} else {
			res := float64(v.H) * 100 / float64(spins)
			if math.IsInf(res, 1) || math.IsNaN(res) {
				res = 0
			}
			childReport.Contribution = res
		}
		report.Each[v.ID] = childReport

	}

	// the rtp
	report.RTP = float64(report.Won) * 100 / float64(report.Spent)
	report.RTPContrib = float64(report.Won) * 100 / float64(report.Spent)

	return report
}
