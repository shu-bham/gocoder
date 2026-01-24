package google

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	symbolTable := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if symbolTable[s[i]] < symbolTable[s[i+1]] {
			ans -= symbolTable[s[i]]
		} else {
			ans += symbolTable[s[i]]
		}
	}
	ans += symbolTable[s[len(s)-1]]
	return ans
}

func romanToIntV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	symbolTable := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	ans := 0
	for i := 0; i < len(s); {
		// Check for two-character symbols first
		if i+1 < len(s) {
			twoCharSymbol := s[i : i+2]
			if val, ok := symbolTable[twoCharSymbol]; ok {
				ans += val
				i += 2
				continue
			}
		}

		// Fallback to one-character symbol
		oneCharSymbol := s[i : i+1]
		ans += symbolTable[oneCharSymbol]
		i++
	}

	return ans
}
