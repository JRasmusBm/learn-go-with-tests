package roman_numerals

import (
	"reflect"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("Table Test", func(t *testing.T) {
		testTable := []struct {
			arabic int
			want   string
		}{
			{arabic: 1, want: "I"},
			{arabic: 2, want: "II"},
			{arabic: 3, want: "III"},
			{arabic: 4, want: "IV"},
			{arabic: 5, want: "V"},
			{arabic: 9, want: "IX"},
			{arabic: 10, want: "X"},
			{arabic: 14, want: "XIV"},
			{arabic: 18, want: "XVIII"},
			{arabic: 20, want: "XX"},
			{arabic: 39, want: "XXXIX"},
		}

		for _, testEntry := range testTable {
			got := ConvertToRoman(testEntry.arabic)
			if !reflect.DeepEqual(got, testEntry.want) {
				t.Errorf("%#v, got %v want %v", testEntry.arabic, got, testEntry.want)
			}
		}
	})
}
