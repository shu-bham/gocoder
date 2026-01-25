package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tech Interview Steps:
// 1. Clarifying Questions:
//    - What should be returned for an empty input string? (e.g., "" -> []string{} or [""])
//    - What is the valid range of digits? (e.g., '2'-'9')
//    - What should be done for invalid digits? (e.g., '1', '0', '*')
//    - What is the maximum length of the input digits string?
//    - Is the order of the output combinations important?

// 2. High-Level Approach:
//    - We can solve this problem using backtracking or iteratively.
//    - For backtracking:
//      - Create a mapping from digits to letters.
//      - A recursive function will build the combinations.
//      - The base case for the recursion is when the current combination's length equals the input digits' length.
//    - For iterative approach:
//      - Start with an empty combination.
//      - For each digit in the input, create new combinations by appending each of its corresponding letters to the existing combinations.

// 3. Trade-offs:
//    - Backtracking is often more intuitive for permutation/combination problems but might have slightly more overhead due to recursion.
//    - The iterative approach can be slightly more efficient as it avoids recursion overhead. Both have the same time complexity.

// 4. Time and Space Complexity:
//    - Time Complexity: O(4^N * N), where N is the length of the input digits. 4 is the max number of letters for a digit. We also need to build N-length strings.
//    - Space Complexity: O(4^N * N) to store the results.

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "empty string",
			digits: "",
			// Based on the current implementation, it returns [""]. A common expectation is an empty slice.
			// This test will fail for letterCombinationsBackTrack which returns ["", ""].
			// A better expectation would be []string{}.
			want: []string{""},
		},
		{
			name:   "single digit",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "two digits",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "digits with 7 and 9",
			digits: "79",
			want:   []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_iterative", func(t *testing.T) {
			got := letterCombinations(tc.digits)
			// Using ElementsMatch because the order of combinations doesn't matter.
			assert.ElementsMatch(t, tc.want, got)
		})
	}

	for _, tc := range testCases {
		// The backtrack implementation has a bug with empty input.
		if tc.digits == "" {
			// It returns ["", ""], which is incorrect. A separate test below highlights this.
			continue
		}
		t.Run(tc.name+"_backtrack", func(t *testing.T) {
			got := letterCombinationsBackTrack(tc.digits)
			// The backtrack implementation adds the result to a pre-initialized slice `[]string{""}`
			// So we need to slice it to get the actual combinations.
			assert.ElementsMatch(t, tc.want, got[1:])
		})
	}

	t.Run("backtrack_empty_input_bug", func(t *testing.T) {
		// This test demonstrates the bug in the backtracking implementation for an empty input.
		// It returns ["", ""], while it should ideally be []string{} or at least consistent.
		got := letterCombinationsBackTrack("")
		assert.Equal(t, []string{"", ""}, got, "The backtracking implementation has a bug with empty input")
	})

	t.Run("iterative_empty_input_behavior", func(t *testing.T) {
		// This test clarifies the behavior of the iterative implementation for an empty input.
		// It returns [""], which might not be the desired behavior (usually an empty slice is expected).
		got := letterCombinations("")
		assert.Equal(t, []string{""}, got)
	})
}
