package recorder

type Recorder struct {
	// for the streak
	S int64
	// total streak
	SC int64
	// for hit frequency
	H int64
	// failure and failed
	F int64
}

// hit the data
func (r *Recorder) Hit() {
	r.H++
	r.S++
}

// unhit  for failed record
func (r *Recorder) Unhit() {
	r.F++
	r.Streak()
}

// finish the record
func (r *Recorder) Streak() {
	if r.S > 0 {
		r.SC++
		r.S = 0
	}
}

// hit frequency
func (r Recorder) HF() float64 {
	return float64(r.F) / float64(r.H)
}

// the streak avarage
func (r Recorder) SAVg() float64 {
	return float64(r.S) / float64(r.SC)
}
