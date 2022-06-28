package clockface

import (
	"math"
	"time"
)

type Clockface struct {
	scale  float64
	origin float64
}

func New(origin, scale float64) *Clockface {
	return &Clockface{scale: scale, origin: origin}
}

func (c *Clockface) handPosition(angle float64) Point {
	return Point{
		X: c.origin + math.Sin(angle)*c.scale,
		Y: c.origin - math.Cos(angle)*c.scale,
	}
}

func (c *Clockface) SecondHand(tm time.Time) Point {
	return c.handPosition(secondsInRadians(tm))
}

func (c *Clockface) MinuteHand(tm time.Time) Point {
	return c.handPosition(minutesInRadians(tm))
}

func (c *Clockface) HourHand(tm time.Time) Point {
	return c.handPosition(hoursInRadians(tm))
}
