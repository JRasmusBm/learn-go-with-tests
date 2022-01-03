package racer

import (
	"testing"
)

func TestRacer(t *testing.T) {
	t.Run("Returns the faster url", func(t *testing.T) {
		slowUrl := "http://www.facebook.com"
		fastUrl := "http://www.quii.co.uk"

		got := Racer(slowUrl, fastUrl)
		want := fastUrl

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Input order does not matter", func(t *testing.T) {
		slowUrl := "http://www.facebook.com"
		fastUrl := "http://www.quii.co.uk"

		got := Racer(fastUrl, slowUrl)
		want := fastUrl

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
