package roman_numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

loop:
	for i := arabic; i > 0; i-- {
		switch i {
		case 5:
			result.WriteString("V")
			break loop
		case 4:
			result.WriteString("IV")
			break loop
		default:
			result.WriteString("I")
		}
	}

	return result.String()
}
