package clockface

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(time.Time) Point {
	return Point{0, -1}
}
