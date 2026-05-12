# WR API Documentation

This document provides a technical reference for all public packages, types, and methods in the `wr` library.

---

## Table of Contents
1. [Package: wr](#package-wr)
2. [Package: ranges](#package-ranges)
3. [Package: record](#package-record)

---

## Package: `wr`

The core package for weighted selection and simulations.

### Interfaces

#### `Wer`
The interface that items must implement to be used in weighted selections.
- `Info() *W`: Returns the underlying weight tracking struct.
- `Reward() int64`: Returns the reward value associated with this item.

#### `Randomizor`
Interface for generating random numbers.
- `Rand(n int64) int64`: Should return a random integer in the range `[0, n)`.

#### `Simulator`
An interface for running batch simulations.
- **Methods**:
  - `Simulate(bet int64, spins int64) record.Report`: Runs the simulation and returns a comprehensive report.

### Types

#### `W`
The base struct for tracking weighted items.
- **Fields**:
  - `ID`: Unique identifier.
  - `Weights`: Relative probability weight.
  - `IsEmpty`: Boolean indicating if this is a "null" or losing outcome.
- **Methods**:
  - `Init(slot *Slots)`: Initializes the internal record and links to parent slots.

#### `Slots`
Manages a collection of weighted items. Implements `Simulator`.
- **Fields**:
  - `Lists`: Slice of items implementing `Wer`.
  - `Track`: Boolean to enable/disable hit tracking.
- **Methods**:
  - `Init(rand Randomizor)`: Initializes the slot system and calculates total weight.
  - `Spin() W`: Executes a weighted selection and returns the selected item.
  - `Simulate(bet int64, spins int64) record.Report`: Runs the simulation for a weighted list.

### Variables
- `Default`: A standard implementation of `Randomizor` using `math/rand/v2`.
- `Secure`: A cryptographically secure implementation of `Randomizor` using `crypto/rand`.

---

## Package: `ranges`

Package for range-based probability tracking (e.g., Hilo games).

### Interfaces

#### `Rer`
- `Info() *R`: Returns the range configuration.
- `Reward() int64`: Returns the reward for the current hit.

### Types

#### `R`
Configuration for range-based rolls.
- **Fields**:
  - `Range`: The maximum possible roll value.
- **Methods**:
  - `NewR(r int64) R`: Factory function for a new range.
  - `Spin(rand wr.Randomizor) int64`: Performs a roll and returns the result in range `[0, Range)`.
  - `Hit(rer Rer)`: Records a win.
  - `Unhit()`: Records a loss and updates streaks.
  - `Simulate(bet int64, spins int64) record.Report`: Generates a report from recorded data.

#### `Simulator`
A helper to run range simulations. Implements `wr.Simulator`.
- **Fields**:
  - `R`: The `R` instance to use.
  - `Task`: A function that defines the logic for each spin (e.g., calling `Spin` and `Hit`/`Unhit`).
- **Methods**:
  - `Simulate(bet int64, spins int64) record.Report`: Runs the `Task` for `spins` times and returns a report.

---

## Package: `record`

Shared logic for recording statistics and generating reports.

### Types

#### `Record`
The foundation for all statistical tracking.
- **Methods**:
  - `Hit()`: Increments hit count.
  - `Unhit()`: Increments failure count and processes the current streak.
  - `HF() float64`: Returns the Hit Frequency.
  - `Savg() float64`: Returns the average win streak length.

#### `Report`
The output of a simulation or session.
- **Fields**:
  - `Hit / Fail`: Total counts.
  - `StreakResult`: Map of streak lengths to their frequency.
  - `Spent / Won`: Financial tracking.
  - `Bet`: The bet amount used for the simulation.
  - `RTP`: Total Return to Player percentage.
  - `RTPContrib`: Return to Player contribution of a specific item.
  - `Contribution`: Frequency contribution of a specific item.
  - `Each`: Map of sub-reports for individual items/rewards.
- **Methods**:
  - `Print()`: Prints a JSON-formatted report to standard output.
  - `Printable() string`: Returns the JSON representation as a string.
