package clockface_test

import (
	"testing"
	"time"

	"github.com/JRasmusBm/learn-go-with-tests/clockface"
)

func TestSecondHand(t *testing.T) {
	origin := 150.0
	scale := 90.0

	t.Run("Points up at 0 seconds", func(t *testing.T) {
		clock := clockface.New(origin, scale)
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 - 90}
		got := clock.SecondHand(tm)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Points down at 30 seconds", func(t *testing.T) {
		clock := clockface.New(origin, scale)
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 + 90}
		got := clock.SecondHand(tm)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
