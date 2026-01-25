package google

func letterCombinationsBackTrack(digits string) []string {
	if digits == "" {
		return []string{}
	}
	allCombinations := []string{}
	symbolMap := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}

	var backtrack func(combination string, runes []rune)
	backtrack = func(combination string, runes []rune) {
		if len(runes) == 0 {
			allCombinations = append(allCombinations, combination)
			return
		}
		currDigit := runes[0]
		for _, symbol := range symbolMap[currDigit] {
			backtrack(combination+string(symbol), runes[1:])
		}
	}
	runes := []rune(digits)
	backtrack("", runes)
	return allCombinations
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	allCombinations := []string{""}
	symbolMap := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}

	for _, digit := range digits {
		combWithDigit := make([]string, 0)
		for _, char := range symbolMap[digit] {
			for _, combination := range allCombinations {
				newStr := combination
				newStr += string(char)
				combWithDigit = append(combWithDigit, newStr)
			}
		}
		allCombinations = combWithDigit
	}
	return allCombinations
}
