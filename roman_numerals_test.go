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
			{arabic: 40, want: "XL"},
			{arabic: 47, want: "XLVII"},
			{arabic: 49, want: "XLIX"},
			{arabic: 50, want: "L"},
			{arabic: 90, want: "XC"},
			{arabic: 100, want: "C"},
			{arabic: 400, want: "CD"},
			{arabic: 500, want: "D"},
			{arabic: 798, want: "DCCXCVIII"},
			{arabic: 900, want: "CM"},
			{arabic: 1000, want: "M"},
			{arabic: 1006, want: "MVI"},
			{arabic: 1984, want: "MCMLXXXIV"},
			{arabic: 3999, want: "MMMCMXCIX"},
			{arabic: 2014, want: "MMXIV"},
		}

		for _, testEntry := range testTable {
			got := ConvertToRoman(testEntry.arabic)
			if !reflect.DeepEqual(got, testEntry.want) {
				t.Errorf("%#v, got %v want %v", testEntry.arabic, got, testEntry.want)
			}
		}
	})
}
