package wr

import (
	crand "crypto/rand"
	"math/big"
	"math/rand/v2"
)

/*
Interface to use for randomization it can be in any algo as long as they return a number
*/

type random struct{}

func (random) Rand(end int64) int64 {
	return rand.Int64N(end)
}

// the default random generator
var Default = random{}

type secure struct{}

func (secure) Rand(end int64) int64 {
	if end <= 0 {
		return 0
	}
	res, err := crand.Int(crand.Reader, big.NewInt(end))
	if err != nil {
		return 0
	}
	return res.Int64()
}

// the secure random generator
var Secure = secure{}

func ToWer(data Wer) Wer {
	return data
}
