# CHANGELOG 
## v0.1.0 - 2026-05-12
### Added
- New `Simulator` interface for running batch simulations.
- `Slots` now implements `Simulator` for weighted selection simulations.
- `Bet` field added to `Report` for accurate RTP calculations.
- `RTPContrib` and `Contribution` metrics for detailed analysis of each outcome.
- `ToWer` helper function in `wr` package.
- `Rer` added Id() for simulation 

### Changed
- `Simulator` changed from a struct to an interface to support different simulation types.
- Simulation reports now include more granular data per item.

## [UNRELEASED]
- Support for range-based simulations in the `Simulator` interface.
- Improved RTP calculation for dynamic rolls.
