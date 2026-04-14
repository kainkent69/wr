package main

import (
	// "encoding/json"
	// "fmt"

	"encoding/json"
	"fmt"
	"log"

	"github.com/kainkent69/wr/src/wr"
	"github.com/kainkent69/wr/src/wr/simulate"
)

func main() {
	list := []*wr.W{
		{
			ID:      1,
			Weights: 5000,
		},
		{
			ID:      2,
			Weights: 5000,
		},

		{
			ID:      3,
			Weights: 5000,
		},

		{
			ID:      4,
			Weights: 1000,
		},
	}

	// record the slots
	slot := wr.Slots{
		Lists: list,
		Track: true,
	}

	slot.Init(wr.Default)

	sm := simulate.Simulator{
		List:  list,
		Spins: 100000,
	}

	fmt.Println("runnig a report")
	report := sm.Run(wr.Default)
	fmt.Println("got report")
	fmt.Printf("the report: %+v\n", report)

	js, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal indent", err)
	}
	fmt.Println(string(js))

}
