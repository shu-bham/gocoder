package google

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tech Interview Steps:
// 1. Clarifying Questions:
//    - Input: A list of pairs, where each pair contains an employee's name (string) and an access time (string).
//    - Access Time Format: The time is given as a string of four digits "HHMM" representing hours and minutes.
//    - High Access Definition: An employee is considered "high-access" if they have three or more accesses within a one-hour period.
//      - For example, if an employee has accesses at "0800", "0845", and "0859", they are high-access.
//      - Accesses at "0800", "0900" are one hour apart, so "0800", "0859" are within one hour.
//    - Output: Return a list of names of all high-access employees.
//    - Edge Cases:
//      - What if an employee has multiple sets of high-access periods? (e.g., 3 times in the morning, 3 times in the evening). The name should appear only once in the result.
//      - What if the input list is empty? Return an empty list.
//      - What if no employee meets the high-access criteria? Return an empty list.
//      - Are employee names case-sensitive? (Assuming yes based on standard string comparisons).

// 2. High-Level Plan:
//    - Group access times by employee. A hash map `map[string][]int` is a good data structure for this, mapping employee names to a list of their access times.
//    - Convert time strings to a numerical format (e.g., integers representing minutes from midnight or simply the integer value of "HHMM") to make comparisons easier. For "HHMM", `HH * 60 + MM` or just treating "HHMM" as an integer works if we are careful. Since `100` represents one hour, simply converting "HHMM" to an integer is simpler.
//    - For each employee, sort their access times in ascending order.
//    - Iterate through the sorted access times for each employee using a sliding window of size 3.
//    - For each window of 3 accesses (let's say `t1, t2, t3`), check if the difference between the third and first access (`t3 - t1`) is less than 100 (which corresponds to 1 hour, e.g., 0859 - 0800 = 59).
//    - If the condition is met, the employee is a high-access employee. Add their name to a result list (or a set to handle duplicates) and move to the next employee.
//    - Return the list of high-access employee names.

// 3. Complexity Analysis:
//    - Let N be the total number of access records.
//    - Sorting the initial access_times takes O(N log N).
//    - Populating the map takes O(N).
//    - Iterating through each employee's times:
//      - Let K be the number of unique employees.
//      - Let `N_u` be the number of accesses for employee `u`. Sum of all `N_u` is N.
//      - Sorting times for each employee is already done by the initial sort.
//      - The sliding window check for each employee takes O(N_u).
//      - Total time for checking all employees is Sum(O(N_u)) = O(N).
//    - Overall Time Complexity: O(N log N) dominated by the initial sort.
//    - Space Complexity: O(N) to store the `userTimeMap`.

func Test_findHighAccessEmployees(t *testing.T) {
	t.Run("Test Case 1: Example with high-access employees", func(t *testing.T) {
		access_times := [][]string{
			{"a", "0549"}, {"b", "0457"}, {"a", "0532"},
			{"a", "0621"}, {"b", "0540"},
		}
		expected := []string{"a"}
		// Note: The implementation returns a sorted list of names, but the problem description doesn't require it.
		// Sorting the result for consistent comparison in tests.
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Should identify 'a' as a high-access employee")
	})

	t.Run("Test Case 2: No high-access employees", func(t *testing.T) {
		access_times := [][]string{
			{"a", "0100"}, {"a", "0200"}, {"a", "0300"},
		}
		expected := []string{}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Should return an empty list when no employee has 3 accesses within an hour")
	})

	t.Run("Test Case 3: Multiple high-access employees", func(t *testing.T) {
		access_times := [][]string{
			{"a", "0900"}, {"a", "0930"}, {"a", "0959"},
			{"b", "1200"}, {"b", "1215"}, {"b", "1259"},
			{"c", "1400"}, {"c", "1500"}, {"c", "1600"},
		}
		expected := []string{"a", "b"}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Should identify both 'a' and 'b' as high-access")
	})

	t.Run("Test Case 4: Employee with accesses exactly one hour apart", func(t *testing.T) {
		access_times := [][]string{
			{"d", "0800"}, {"d", "0830"}, {"d", "0900"},
		}
		expected := []string{}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Accesses exactly one hour apart should not be counted as high-access")
	})

	t.Run("Test Case 5: Empty input", func(t *testing.T) {
		access_times := [][]string{}
		expected := []string{}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Should return an empty list for empty input")
	})

	t.Run("Test Case 6: Employee with fewer than 3 accesses", func(t *testing.T) {
		access_times := [][]string{
			{"a", "0100"}, {"a", "0130"},
			{"b", "0200"},
		}
		expected := []string{}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Should not consider employees with less than 3 accesses")
	})

	t.Run("Test Case 7: Employee with multiple high-access windows", func(t *testing.T) {
		access_times := [][]string{
			{"a", "0800"}, {"a", "0810"}, {"a", "0820"}, // First window
			{"a", "1700"}, {"a", "1705"}, {"a", "1710"}, // Second window
		}
		expected := []string{"a"}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Employee name should appear only once in the result")
	})

	t.Run("Test Case 8: Complex scenario with mixed accesses", func(t *testing.T) {
		access_times := [][]string{
			{"a", "0110"}, {"b", "2250"}, {"a", "0125"}, {"a", "0150"},
			{"b", "2310"}, {"c", "0800"}, {"b", "2330"},
		}
		expected := []string{"a", "b"}
		result := findHighAccessEmployees(access_times)
		assert.ElementsMatch(t, expected, result, "Should correctly identify multiple high-access employees in a mixed log")
	})
}
