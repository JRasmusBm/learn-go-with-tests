package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 1.0
	minuteHandLength = 1.0
	hourHandLength   = 1.0
	clockCentreX     = 150
	clockCentreY     = 150
)

type Clockface struct {
	scale  float64
	origin float64
}

func New(origin, scale float64) *Clockface {
	return &Clockface{scale: scale, origin: origin}
}

func (c *Clockface) handPosition(angle, modifier float64) Point {
	return Point{
		X: c.origin + math.Sin(angle)*c.scale*modifier,
		Y: c.origin - math.Cos(angle)*c.scale*modifier,
	}
}

func (c *Clockface) SecondHand(tm time.Time) Point {
	return c.handPosition(secondsInRadians(tm), secondHandLength)
}

func (c *Clockface) MinuteHand(tm time.Time) Point {
	return c.handPosition(minutesInRadians(tm), minuteHandLength)
}

func (c *Clockface) HourHand(tm time.Time) Point {
	return c.handPosition(hoursInRadians(tm), hourHandLength)
}
