package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Equals(other Point) bool {
	const threshold = 1e7

	return math.Abs(p.X-other.X) < threshold && math.Abs(p.Y-other.Y) < threshold
}

func SecondHand(tm time.Time) Point {
	var origin float64 = 150.0

	if tm.Second() > 0 {
		return Point{X: 150, Y: origin + 90}
	}

	return Point{X: 150, Y: origin - 90}
}
