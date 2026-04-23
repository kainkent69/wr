package wr

import "math/rand/v2"

/*
Interface to use for randomization it can be in any algo as long as they return a number
*/

type random struct{}

func (random) Rand(end int64) int64 {
	return rand.Int64N(end)
}

// the default random generator
var Default = random{}
