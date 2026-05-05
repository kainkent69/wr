package ranges_test

import (
	"testing"

	"github.com/kainkent69/wr"
	"github.com/kainkent69/wr/ranges"
)

type MockHilo struct {
	r      *ranges.R
	reward int64
	Spins  int64
	Bet    int64
}

func (m *MockHilo) Info() *ranges.R {
	return m.r
}

func (m *MockHilo) Reward() int64 {
	return m.reward
}

func (m *MockHilo) Run(prob int64, bet int64) {
	m.r.A = m.r.Range / (100 / prob)
	res := m.Info().Roll(wr.Default)
	m.reward = bet * prob / (100 + 3)
	m.Spins++
	if res {
		m.r.Hit(m)
	} else {
		m.r.Unhit()
	}

}

func TestRange(t *testing.T) {
	n := ranges.NewR(10000, 0)
	var M = &MockHilo{
		r: &n,
	}
	var prob int64 = 50
	var bet int64 = 100
	for range 100 {
		M.Run(prob, bet)
	}

	report := M.r.NewReport(bet, M.Spins)
	report.Print()
}
