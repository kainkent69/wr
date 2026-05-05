# WR

A Go library for weighted selection and range-based probability tracking, designed for slot machines, Hilo games, and other probability-driven applications.

## Features

- **Weighted Selection**: Efficiently select items based on relative weights with built-in tracking.
- **Range Probability**: Handle dynamic win conditions where the "hit" range changes per roll.
- **Advanced Statistics**: Automatically tracks Hit Frequency (HF), win streaks (min/max/avg), and streak distributions.
- **RTP Analysis**: Calculate Return to Player (RTP) and individual item contributions to total returns.
- **Simulation Suite**: Easily run large-scale simulations to verify math models.

## Installation

```bash
go get github.com/kainkent69/wr
```

## Documentation

### 1. Weighted Selection (`wr`)

The `wr` package provides a system for selecting items from a weighted list.

#### Implementation
To use weighted selection, your items should implement the `Wer` interface:

```go
type Wer interface {
    Info() *W
    Reward() int64
}
```

#### Basic Usage
```go
package main

import (
	"github.com/kainkent69/wr"
)

type SlotItem struct {
	*wr.W
	reward int64
}

func (s *SlotItem) Info() *wr.W { return s.W }
func (s *SlotItem) Reward() int64 { return s.reward }

func main() {
	items := []wr.Wer{
		&SlotItem{W: &wr.W{ID: 1, Weights: 5000}, reward: 500},
		&SlotItem{W: &wr.W{ID: 2, Weights: 5000}, reward: 1000},
		&SlotItem{W: &wr.W{ID: 3, Weights: 10000, IsEmpty: true}, reward: 0},
	}

	simulator := wr.Simulator{
		List:  items,
		Spins: 100000,
		Bet:   100,
	}

	report := simulator.Run(wr.Default)
	report.Print()
}
```

### 2. Range Selection (`ranges`)

The `ranges` package is ideal for games like Hilo where a "hit" occurs if a random roll falls within a specific range `[0, A]`.

#### Basic Usage
```go
package main

import (
	"github.com/kainkent69/wr"
	"github.com/kainkent69/wr/ranges"
)

type HiloGame struct {
	r      *ranges.R
	reward int64
}

func (h *HiloGame) Info() *ranges.R { return h.r }
func (h *HiloGame) Reward() int64   { return h.reward }

func main() {
	// Initialize a range with a maximum of 10,000
	r := ranges.NewR(10000, 0)
	game := &HiloGame{r: &r}

	// Set probability to 50%
	game.r.A = 5000
	game.reward = 200

	if game.r.Roll(wr.Default) {
		game.r.Hit(game)
	} else {
		game.r.Unhit()
	}

	report := game.r.NewReport(100, 1)
	report.Print()
}
```

## Statistical Metrics Explained

The generated `record.Report` contains detailed data:

| Metric | Description |
| :--- | :--- |
| **HF** | Hit Frequency (Hits / Failures). |
| **SAvg** | Average Winning Streak length. |
| **Hit / Fail** | Total number of winning and losing events. |
| **StreakResult** | Distribution map of streak lengths (e.g., `"3": 10` means a 3-win streak occurred 10 times). |
| **Min/Max Streak** | The shortest and longest winning streaks recorded. |
| **RTP** | Return to Player percentage. |
| **Contribution** | How much a specific outcome contributed to the total hits or RTP. |

## Example Simulation Output

```json
{
   "HF": 1.49,
   "SAvg": 2.49,
   "Hit": 59942,
   "Fail": 40058,
   "StreakResult": {
      "1": 9595,
      "2": 5879,
      "3": 3504,
      ...
   },
   "MinStreak": 1,
   "MaxStreak": 21,
   "Spent": 10000000,
   "Won": 29977500,
   "RTP": 299.77,
   "Each": { ... }
}
```

## License
MIT
