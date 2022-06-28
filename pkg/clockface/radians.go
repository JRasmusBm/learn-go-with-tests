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

func hoursInRadians(t time.Time) float64 {
	return (math.Pi / (6 / (float64(t.Hour() % 12))))
}
