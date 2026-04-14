# WR

## A Basic weighted selection used for slot machines

### download

```bash
go get github.com/kainkent69/wr
```

### Basic Usage
```go
package main_test

import (
	"github.com/kainkent69/wr/src/wr"
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
			Weights: 10000,
			IsEmpty: true,
		},
	}

	// record the slots
	slot := wr.Slots{
		Lists: list,
		Track: true,
	}

	slot.Init(wr.Default)

	sm := wr.Simulator{
		List:  list,
		Spins: 100000,
	}

	report := sm.Run(wr.Default)
	report.Print()

}

```

### Result 

*Note*: this defer because its random

```bash
{
   "HF": 3,
   "SAvg": 3,
   "Hit": 300000,
   "Fail": 100000,
   "StreakResult": {
      "3": 100000
   },
   "MinStreak": 3,
   "MaxStreak": 3,
   "Contirbution": 0,
   "IsEmpty": false,
   "Each": {
      "1": {
         "HF": 0.24951581262260875,
         "SAvg": 0,
         "Hit": 19969,
         "Fail": 80031,
         "StreakResult": null,
         "MinStreak": 0,
         "MaxStreak": 0,
         "Contirbution": 6.656333333333333,
         "IsEmpty": false,
         "Each": null
      },
      "2": {
         "HF": 0.24968757810547362,
         "SAvg": 0,
         "Hit": 19980,
         "Fail": 80020,
         "StreakResult": null,
         "MinStreak": 0,
         "MaxStreak": 0,
         "Contirbution": 6.66,
         "IsEmpty": false,
         "Each": null
      },
      "3": {
         "HF": 0.25237952109007117,
         "SAvg": 0,
         "Hit": 20152,
         "Fail": 79848,
         "StreakResult": null,
         "MinStreak": 0,
         "MaxStreak": 0,
         "Contirbution": 6.717333333333333,
         "IsEmpty": false,
         "Each": null
      },
      "4": {
         "HF": 0.6641981061425553,
         "SAvg": 0,
         "Hit": 39911,
         "Fail": 60089,
         "StreakResult": null,
         "MinStreak": 0,
         "MaxStreak": 0,
         "Contirbution": 0,
         "IsEmpty": true,
         "Each": null
      }
   }
}                                
```

