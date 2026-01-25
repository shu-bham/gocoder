/*
# Integer to Roman

## Problem Description

Given an integer, convert it to a roman numeral. The input is guaranteed to be within the range from 1 to 3999.

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
- I: 1
- V: 5
- X: 10
- L: 50
- C: 100
- D: 500
- M: 1000

And the six subtractive cases:
- IV: 4, IX: 9
- XL: 40, XC: 90
- CD: 400, CM: 900

## Tech Interview Approach Protocol

### 1. Clarification and Edge Cases
Start by asking questions to fully understand the problem's boundaries.

- **Confirm Input/Output:**
    - **Input:** "I'll be receiving an integer."
    - **Output:** "The function should return a string representing that integer in Roman numerals."

- **Ask Clarifying Questions:**
    - **Range:** "The problem states the input is between 1 and 3999. Can I assume it will never be outside this range? For example, what should happen for 0, negative numbers, or numbers greater than 3999?"
    - **Notation Rules:** "Should I use the standard subtractive notation (e.g., 'IX' for 9)? Or would a simpler form like 'VIIII' be acceptable?" (Standard form is almost always implied).

- **Discuss Edge Cases:**
    - **Boundaries:** It's good to mention you'll test the boundaries of the range, like 1 and 3999.
    - **Key Conversion Points:** Mention testing numbers that specifically test the logic, such as those involving subtraction (4, 9, 40, 90, 400, 900) and those that don't (3, 8, 88).

### 2. Greedy Approach

The most effective and straightforward solution is a greedy one. The key is to process the number from largest possible values to smallest.

- **Algorithm:**
    1. Define two parallel arrays (or a single map/struct array), one for integer values and one for their corresponding Roman numeral symbols.
    2. **Crucially**, this list must include the subtractive combinations (900, 400, 90, 40, 9, 4) to handle them correctly.
    3. The list must be sorted in descending order of values, from 1000 ("M") down to 1 ("I").
    4. Initialize an empty string or string builder for the result.
    5. Iterate through the values/symbols list. For each value:
        - While the input number `num` is greater than or equal to the current value, append the corresponding symbol to your result and subtract the value from `num`.
    6. Continue until `num` becomes 0.

- **Example Walkthrough (num = 58):**
    - `values = [..., 50, 40, 10, ...]`
    - `symbols = [..., "L", "XL", "X", ...]`
    1. `num` (58) >= 50? Yes. Result is "L", `num` becomes 8.
    2. `num` (8) >= 40? No.
    3. `num` (8) >= 10? No.
    4. ...
    5. `num` (8) >= 5? Yes. Result is "LV", `num` becomes 3.
    6. `num` (3) >= 4? No.
    7. `num` (3) >= 1? Yes. Result is "LVI", `num` becomes 2.
    8. `num` (2) >= 1? Yes. Result is "LVII", `num` becomes 1.
    9. `num` (1) >= 1? Yes. Result is "LVIII", `num` becomes 0.
    10. Stop. Return "LVIII".

- **Complexity Analysis:**
    - **Time Complexity:** O(1). This might seem counter-intuitive because of the loops. However, the number of symbols is fixed and small (13 in this case). The inner `while` loop for any given symbol can execute at most 3 times (e.g., for 3000, "M" is appended three times). Therefore, the total number of operations is bounded by a constant, not the magnitude of the input number.
    - **Space Complexity:** O(1). The space used by the string builder is also bounded, as the longest Roman numeral for the given range is of a fixed maximum length (e.g., MMMDCCCLXXXVIII for 3888).

The `intToRoman` function in this file implements this greedy approach.
*/
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
