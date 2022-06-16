package clockface

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)

	return Point{math.Sin(angle), math.Cos(angle)}
}
