//go:build race

package tools

// raceEnabled reports whether the race detector is on; it slows code ~20x, so wall-clock
// assertions must be skipped.
const raceEnabled = true
