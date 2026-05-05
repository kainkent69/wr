package record

import (
	"encoding/json"
	"fmt"
	"log"
)

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
