package c

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (p *DoublePendulum) Step(dt float64) {
	bob1, bob2 := p.Bob1, p.Bob2
	// Difference in angles.
	delta := p.Bob2.Theta - p.Bob1.Theta

	//D1 = (m1+m2)l1 - m2(l2)Cos(delta)^2
	D1 := (bob1.Mass+bob2.Mass)*p.L1 - bob2.Mass*p.L2*math.Pow(math.Cos(delta), 2)
	// D2 := (l2 /l1)D1
	D2 := (p.L2 / p.L1) * D1

	// Acceleration bob1
	// (self.mass_bob_2 * self.length_rod_1 * self.omega_1 ** 2 * math.sin(delta) * math.cos(delta) +
	// self.mass_bob_2 * self.g * math.sin(self.theta_2) * math.cos(delta) +
	// self.mass_bob_2 * self.length_rod_2 * self.omega_2 ** 2 * math.sin(delta) -
	// (self.mass_bob_1 + self.mass_bob_2) * self.g * math.sin(self.theta_1)) / denominator_1
	A1 := bob2.Mass*p.L1*math.Pow(bob1.Omega, 2)*math.Sin(delta)*math.Cos(delta) +
		bob2.Mass*p.G*math.Sin(bob2.Theta)*math.Cos(delta) +
		bob2.Mass*p.L2*math.Pow(bob2.Omega, 2)*math.Sin(delta) -
		(bob1.Mass+bob2.Mass)*p.G*math.Sin(bob1.Theta)
	A1 /= D1

	A2 := -bob2.Mass*p.L2*math.Pow(bob2.Omega, 2)*math.Sin(delta)*math.Cos(delta) +
		(bob1.Mass+bob2.Mass)*p.G*math.Sin(bob1.Theta)*math.Cos(delta) -
		(bob1.Mass+bob2.Mass)*p.L1*math.Pow(bob1.Omega, 2)*math.Sin(delta) -
		(bob1.Mass+bob2.Mass)*p.G*math.Sin(bob2.Theta)
	A2 /= D2
}
