package clockface

import "time"

type Point struct {
	X float32
	Y float32
}

func SecondHand(tm time.Time) Point {
	return Point{X: 150, Y: 60}
}
