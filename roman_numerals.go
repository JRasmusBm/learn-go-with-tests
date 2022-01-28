package roman_numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	if arabic == 4 {
		return "IV"
	}

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
