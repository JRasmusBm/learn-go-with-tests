package clockface

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) Equals(other Point) bool {
	const threshold = 1e-7

	return math.Abs(p.X-other.X) < threshold && math.Abs(p.Y-other.Y) < threshold
}
