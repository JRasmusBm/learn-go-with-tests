package clockface

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	if t.Second() > 0 {
		return Point{0, -1}
	}

	return Point{0, 1}
}
