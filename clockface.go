package clockface

import "time"

type Point struct {
	X float32
	Y float32
}

func SecondHand(tm time.Time) Point {
	var origin float32 = 150.0

	if tm.Second() > 0 {
		return Point{X: 150, Y: origin + 90}
	}

	return Point{X: 150, Y: origin - 90}
}


