package google

import "strings"

func intToRoman(num int) string {
	if num < 0 || num > 3999 {
		return ""
	}

	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var roman strings.Builder

	// Iterate through each value/symbol pair.
	for i := 0; i < len(values); i++ {
		// Greedily subtract the current value from num as many times
		// as possible, appending the symbol each time.
		for num >= values[i] {
			roman.WriteString(symbols[i])
			num -= values[i]
		}
	}

	return roman.String()

}
