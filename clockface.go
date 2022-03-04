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

type Clockface struct {
	scale  float64
	origin float64
}

func New(scale, origin float64) *Clockface {
	return &Clockface{scale: scale, origin: origin}
}

func (c *Clockface) handPosition(p Point) Point {
	return Point{
		X: c.origin + p.X*c.scale,
		Y: c.origin - p.Y*c.scale,
	}
}

func (c *Clockface) SecondHand(tm time.Time) Point {
	return c.handPosition(secondHandPoint(tm))
}
