package roman_numerals

import (
	"reflect"
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {
	testTable := []struct {
		arabic uint16
		roman  string
	}{
		{arabic: 1, roman: "I"},
		{arabic: 2, roman: "II"},
		{arabic: 3, roman: "III"},
		{arabic: 4, roman: "IV"},
		{arabic: 5, roman: "V"},
		{arabic: 9, roman: "IX"},
		{arabic: 10, roman: "X"},
		{arabic: 14, roman: "XIV"},
		{arabic: 18, roman: "XVIII"},
		{arabic: 20, roman: "XX"},
		{arabic: 39, roman: "XXXIX"},
		{arabic: 40, roman: "XL"},
		{arabic: 47, roman: "XLVII"},
		{arabic: 49, roman: "XLIX"},
		{arabic: 50, roman: "L"},
		{arabic: 90, roman: "XC"},
		{arabic: 100, roman: "C"},
		{arabic: 400, roman: "CD"},
		{arabic: 500, roman: "D"},
		{arabic: 798, roman: "DCCXCVIII"},
		{arabic: 900, roman: "CM"},
		{arabic: 1000, roman: "M"},
		{arabic: 1006, roman: "MVI"},
		{arabic: 1984, roman: "MCMLXXXIV"},
		{arabic: 3999, roman: "MMMCMXCIX"},
		{arabic: 2014, roman: "MMXIV"},
	}

	t.Run("Arabic to roman", func(t *testing.T) {
		for _, testEntry := range testTable {
			got := ConvertToRoman(testEntry.arabic)
			if !reflect.DeepEqual(got, testEntry.roman) {
				t.Errorf("%#v, got %v want %v", testEntry.arabic, got, testEntry.roman)
			}
		}
	})

	t.Run("Roman to arabic", func(t *testing.T) {
		for _, testEntry := range testTable {
			got := ConvertToArabic(testEntry.roman)
			if !reflect.DeepEqual(got, testEntry.arabic) {
				t.Errorf("%#v, got %v want %v", testEntry.roman, got, testEntry.arabic)
			}
		}
	})

	t.Run("Property test: Conversion is reversible", func(t *testing.T) {
		isReversible := func(arabic uint16) bool {
			roman := ConvertToRoman(arabic)
			newArabic := ConvertToArabic(roman)
			t.Logf("Got %v (%v), want %v", newArabic, roman, arabic)
			return newArabic == arabic
		}

		if err := quick.Check(isReversible, &quick.Config{MaxCount: 1000}); err != nil {
			t.Error("Failed checks: ", err)
		}
	})
}
