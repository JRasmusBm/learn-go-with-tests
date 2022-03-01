package clockface

import "time"

type Point struct {
	X float64
	Y float64
}

func SecondHand(tm time.Time) Point {
	var origin float64 = 150.0

	if tm.Second() > 0 {
		return Point{X: 150, Y: origin + 90}
	}

	return Point{X: 150, Y: origin - 90}
}
