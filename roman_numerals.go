package roman_numerals

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)

	for _, roman_numeral := range r {
		if roman_numeral.Symbol == symbol {
			return roman_numeral.Value
		}
	}

	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (result int) {
	for i := 0; i < len(roman); i++ {
		currentSymbol := roman[i]
		valueWithNextSymbol := 0

		if i+1 < len(roman) {
			valueWithNextSymbol = allRomanNumerals.ValueOf(currentSymbol, roman[i+1])
		}

		valueOfCurrentSymbol := allRomanNumerals.ValueOf(currentSymbol)

		if valueWithNextSymbol == 0 {
			result += valueOfCurrentSymbol
		} else {
			result += valueWithNextSymbol
			i++
		}
	}

	return
}
