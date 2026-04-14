package recorder

import "math"

type Record struct {
	// for the streak
	S int64
	// total streak
	SC int64
	// for total streak
	TS int64

	// for minimum and maximum streak
	SMin int64
	SMax int64
	SMid int64
	// record the streak by size
	SReq map[int64]int64

	// for hit frequency
	H int64
	// failure and failed
	F int64
}

// hit the data
func (r *Record) Hit() {
	r.H++
	r.S++
}

// unhit  for failed record
func (r *Record) Unhit() {
	r.F++
	r.Streak()
}

// finish the record
func (r *Record) Streak() {
	if r.S > 0 {
		// max
		if r.S > r.SMax {
			r.SMax = r.S
		}
		if r.S < r.SMin || r.SMin == 0 {
			r.SMin = r.S
		}
		r.SReq[r.S]++
		r.SC++
		r.TS += r.S
		r.S = 0
	}
}

// hit frequency
func (r Record) HF() float64 {
	res := (float64(r.H) * 100 / float64(r.F)) / 100
	if math.IsInf(res, 1) || math.IsNaN(res) {
		return float64(0)
	}
	return res
}

// the streak avarage
func (r Record) Savg() float64 {
	res := float64(r.TS) / float64(r.SC)
	if math.IsInf(res, 1) || math.IsNaN(res) {
		return float64(0)
	}
	return res
}
