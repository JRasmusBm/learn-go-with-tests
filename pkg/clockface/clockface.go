package clockface

import (
	"time"
)

type Clockface struct {
	scale  float64
	origin float64
}

func New(origin, scale float64) *Clockface {
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

func (c *Clockface) MinuteHand(tm time.Time) Point {
	return c.handPosition(minuteHandPoint(tm))
}
