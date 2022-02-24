package clockface

import (
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	t.Run("Points up at 0 seconds", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := Point{X: 150, Y: 150 - 90}
		got := SecondHand(tm)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
