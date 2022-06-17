package clockface

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return secondsInRadians(t)/60 + math.Pi/(30/float64(t.Minute()))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)

	return Point{math.Sin(angle), math.Cos(angle)}
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)

	return Point{math.Sin(angle), math.Cos(angle)}
}
