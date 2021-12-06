package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Returns the highest number first", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
