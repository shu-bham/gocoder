package google

import "strings"

func intToRoman(num int) string {
	if num < 0 || num > 3999 {
		return ""
	}

	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var ans strings.Builder

	i := 0
	for num > 0 {
		for i < len(values) {
			if num >= values[i] {
				ans.WriteString(symbols[i])
				num -= values[i]
			} else {
				break
			}
		}
		i++
	}
	return ans.String()
}
