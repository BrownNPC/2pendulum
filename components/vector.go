package c

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Vec2 is a wrapper around a raylib Vector2
type Vec2 struct {
	X, Y float64
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func V2[T Number](x, y T) Vec2 {
	return Vec2{float64(x), float64(y)}
}

// vector 2 zero
var V2Z = V2(0, 0)

// convert to raylib vector
func (v Vec2) R() rl.Vector2 {
	return rl.Vector2{float64(v.X), float64(v.Y)}
}
func (v Vec2) Add(o Vec2) Vec2      { return V2(v.X+o.X, v.Y+o.Y) }
func (v Vec2) Sub(o Vec2) Vec2      { return V2(v.X-o.X, v.Y-o.Y) }
func (v Vec2) Mul(o Vec2) Vec2      { return V2(v.X*o.X, v.Y*o.Y) }
func (v Vec2) Scale(s float64) Vec2 { return V2(v.X*s, v.Y*s) }
func (v Vec2) Dot(o Vec2) float64   { return v.X*o.X + v.Y*o.Y }
func (v Vec2) Len() float64         { return float64(math.Hypot(float64(v.X), float64(v.Y))) }
func (v Vec2) Norm() Vec2 {
	l := v.Len()
	if l == 0 {
		return V2Z
	}
	return v.Scale(1 / l)
}
func (v Vec2) Dist(o Vec2) float64 {
	return v.Sub(o).Len()
}
func (v Vec2) XY() (float64, float64) {
	return v.X, v.Y
}
func (v Vec2) ToInt() (int, int) {
	return int(v.X), int(v.Y)
}
