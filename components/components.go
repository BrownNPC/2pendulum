package c

import rl "github.com/gen2brain/raylib-go/raylib"

type Bob struct {
	// Current Position of the Bob (pixels)
	Vec2
	// Mass in arbritary units.
	Mass float64
	// Angle of bob from vertical (radians)
	Theta float64
	// Angular velocity of Bob (rad/s)
	Omega float64
}
type DoublePendulum struct {
	// coordinates of the fixed pivot
	Origin Vec2
	//Rod length in pixels
	L1, L2     float64
	Bob1, Bob2 Bob
	// Gravitational acceleration (pixels/s2 after scaling).
	G float64
}

func NewDefaultDoublePendulum() DoublePendulum {
	return DoublePendulum{
		Origin: V2(300, 100),
		L1:     120,
		L2:     120,
		Bob1: Bob{
			Vec2:  Vec2{},
			Mass:  10,
			Theta: rl.Pi / 2,
			Omega: 0,
		},
		Bob2: Bob{
			Vec2:  Vec2{},
			Mass:  10,
			Theta: rl.Pi / 2,
			Omega: 0,
		},
		G: 9.81,
	}
}
