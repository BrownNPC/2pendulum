package c

import rl "github.com/gen2brain/raylib-go/raylib"

// TicksPerSecondRunner runs the operation a specified number of times per second
// A zero value TicksPerSecondRunner is ready to use.
type TicksPerSecondRunner float32

func (t *TicksPerSecondRunner) Tick(TicksPerSecond int, operation func()) {
	// the float32 property of TicksPerSecondRunner is treated as an accumulator.

	dt := TicksPerSecondRunner(rl.GetFrameTime())
	*t += dt

	for *t > dt {
		operation()
		*t -= 1.0 / TicksPerSecondRunner(TicksPerSecond)
	}
}
