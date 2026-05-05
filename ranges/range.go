package ranges

import (
	"math"

	"github.com/kainkent69/wr"
	"github.com/kainkent69/wr/record"
)

type R struct {
	// start
	A int64
	// range
	Range int64
	// for recording and simulations
	record.Record
	Each map[int64]*record.Record
	record.Report
}

func NewR(r int64, a int64) R {
	return R{
		A:     a,
		Range: r,
		Record: record.Record{
			SReq: map[int64]int64{},
		},
	}
}

func (r *R) Roll(rand wr.Randomizor) bool {
	rng := rand.Rand(r.Range)
	if rng <= r.A {
		return true
	}
	return false
}

func (r *R) Hit(rer Rer) {
	if r.Each == nil {
		r.Each = map[int64]*record.Record{}
	}
	if r.Report.Each == nil {
		r.Report.Each = map[int64]record.Report{}
	}

	reward := rer.Reward()
	each, ok := r.Each[reward]
	if !ok {
		each = &record.Record{
			SReq: map[int64]int64{},
		}
		r.Each[reward] = each
	}

	each.Hit()
	r.Record.Hit()

	r.Won += reward
	report, ok := r.Report.Each[reward]
	if !ok {
		report = record.Report{}
	}
	report.Won += reward
	// record
	r.Report.Each[reward] = report
}

type Rer interface {
	Info() *R
	Reward() int64
}

func (s *R) NewReport(bet int64, spins int64) record.Report {
	s.Report.Hit = s.H
	s.Fail = s.F
	s.StreakResult = s.SReq
	s.MinStreak = s.SMin
	s.MaxStreak = s.SMax
	s.Report.HF = s.Record.HF()
	s.Report.SAvg = s.Record.Savg()

	s.Report.Spent = bet * spins
	for k, v := range s.Each {
		var won = s.Report.Each[k].Won
		childReport := record.Report{
			HF:         v.HF(),
			Hit:        v.H,
			Fail:       v.F,
			Won:        won,
			RTPContrib: float64(won) * 100 / float64(s.Spent),
		}

		res := float64(v.H) * 100 / float64(spins)
		if math.IsInf(res, 1) || math.IsNaN(res) {
			res = 0
		}
		childReport.Contirbution = res
		s.Report.Each[k] = childReport

	}

	// the rtp
	s.Report.RTP = float64(s.Report.Won) * 100 / float64(s.Spent)
	s.Report.RTPContrib = float64(s.Report.Won) * 100 / float64(s.Spent)
	return s.Report
}
