/*
# Roman to Integer

## Problem Description

Given a roman numeral, convert it to an integer.
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
- I: 1
- V: 5
- X: 10
- L: 50
- C: 100
- D: 500
- M: 1000

Usually, roman numerals are written largest to smallest from left to right. However, there are six instances where subtraction is used:
- I can be placed before V (5) and X (10) to make 4 and 9.
- X can be placed before L (50) and C (100) to make 40 and 90.
- C can be placed before D (500) and M (1000) to make 400 and 900.

## Tech Interview Approach Protocol

### 1. Clarification and Edge Cases
Before coding, clarify ambiguities and confirm assumptions.

- **Confirm Input/Output:**
    - **Input:** "The input will be a string."
    - **Output:** "The function should return an integer."

- **Ask Clarifying Questions:**
    - **Validity:** "Can I assume the input string is always a valid Roman numeral, or should I validate it? For instance, could I see inputs like 'IIII' or 'VX'?" (Most problems guarantee valid input).
    - **Range:** "What is the range of numbers the numeral can represent? Is it the standard 1 to 3999?" (This confirms you won't see symbols repeated more than allowed, e.g., for numbers > 3999).
    - **Case Sensitivity:** "Are the Roman numerals always uppercase?"

- **Discuss Edge Cases:**
    - **Empty String:** "What is the expected output for an empty input string? Should it be 0?"
    - **Malformed/Invalid Input:** "If validation is required, how should I handle an invalid input? Return an error, or a specific value like -1?"

### 2. Approach 1: Left-to-Right Pass with Lookahead

This is a very common and intuitive approach. The core idea is that if a smaller symbol appears before a larger symbol, it's a subtraction.

- **Algorithm:**
    1. Create a map of Roman symbols to their integer values.
    2. Initialize a running total `ans`.
    3. Iterate through the string from left to right, up to the second-to-last character.
    4. At each character `s[i]`, compare its value with the value of the next character `s[i+1]`.
        - If `value(s[i]) < value(s[i+1])`, it's a subtractive case (like 'IV' or 'IX'). Subtract `value(s[i])` from `ans`.
        - Otherwise, add `value(s[i])` to `ans`.
    5. After the loop, add the value of the last character `s[len(s)-1]` to the total. This character is always added.

- **Complexity Analysis:**
    - **Time Complexity:** O(N), where N is the length of the string, because we do a single pass.
    - **Space Complexity:** O(1), as the map size is constant regardless of the input string length.

The `romanToInt` function in this file implements this logic.

### 3. Approach 2: Handling Special Subtractive Symbols

This approach is also valid and sometimes simpler to reason about. It involves explicitly checking for the two-character subtractive symbols.

- **Algorithm:**
    1. Create a map that includes both single-character symbols ("I", "V") and two-character subtractive symbols ("IV", "IX", "XL", etc.).
    2. Iterate through the string with an index `i`.
    3. In each step, check if the substring `s[i:i+2]` exists in our map.
        - If it does, add the corresponding value to our total and advance the index `i` by 2.
        - If it doesn't, it must be a single-character symbol. Add the value of `s[i:i+1]` to the total and advance the index `i` by 1.
    4. Continue until the end of the string.

- **Complexity Analysis:**
    - **Time Complexity:** O(N), as we still iterate through the string once.
    - **Space Complexity:** O(1), the map size is constant.

The `romanToIntV2` function implements this logic. Both approaches are efficient and acceptable in an interview.

*/
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
