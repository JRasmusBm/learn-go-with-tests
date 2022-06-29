package clockface

import "math"

type Point struct {
	X float64
	Y float64
}

func roughlyEqual(a, b float64) bool {
	const threshold = 1e-7

	return math.Abs(a-b) < threshold
}

func (p Point) Equals(other Point) bool {
	return roughlyEqual(p.X, other.X) && roughlyEqual(p.Y, other.Y)
}
