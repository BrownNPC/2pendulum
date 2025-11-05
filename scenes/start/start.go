package start

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Pendulum c.DoublePendulum
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	scene.Pendulum = c.NewDefaultDoublePendulum()
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	// https://coolors.co/palette/003049-d62828-f77f00-fcbf49-eae2b7
	rl.ClearBackground(rl.GetColor(0x003049))
	rl.DrawFPS(0, 0)
	if rl.IsKeyPressed(rl.KeyEqual) {
		rl.SetTargetFPS(rl.GetFPS() + 30)
	}
	if rl.IsKeyPressed(rl.KeyMinus) {
		rl.SetTargetFPS(rl.GetFPS() - 30)
	}
	scene.Pendulum.Tick(300, 0.06)

	b1x, b1y := scene.Pendulum.Bob1.Position.XY()
	b2x, b2y := scene.Pendulum.Bob2.Position.XY()

	rl.DrawLineV(c.V2(b1x, b1y).R(), scene.Pendulum.Origin.R(), rl.GetColor(0xEAE2B7FF))
	rl.DrawCircle(int32(b1x), int32(b1y), 12, rl.GetColor(0xD62828FF))

	rl.DrawLineV(c.V2(b2x, b2y).R(), c.V2(b1x, b1y).R(), rl.GetColor(0xEAE2B7FF))
	rl.DrawCircle(int32(b2x), int32(b2y), 12, rl.GetColor(0xF77F00FF))

	return false // if true is returned, Unload is called
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "someOtherSceneId" // the engine will switch to the scene that is registered with this id
}
