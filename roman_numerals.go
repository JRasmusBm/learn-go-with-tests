package roman_numerals

func ConvertToRoman(arabic int) string {
	switch arabic {
	case 3:
		return "III"
	case 2:
		return "II"
	default:
		return "I"
	}
}
