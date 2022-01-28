package roman_numerals

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("Converts 1 to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
