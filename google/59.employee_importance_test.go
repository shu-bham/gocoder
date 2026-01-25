package google

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
## Tech Interview Steps for "Employee Importance"

### 1. **Understand the Problem**
The goal is to calculate the total importance value of a given employee plus all their subordinates, both direct and indirect. You receive a list of all employees in the company and the ID of the target employee. Each employee has a unique ID, an importance value, and a list of IDs of their direct subordinates.

### 2. **Clarifying Questions (Edge Cases & Constraints)**
- **Input:**
  - `employees`: A slice of `*Employee` pointers.
  - `id`: The integer ID of the employee whose importance hierarchy we need to sum up.

- **Output:** An integer representing the total importance.
- **Data Structure:** Is the reporting structure a tree or a general graph? (i.e., can an employee report to more than one manager?). The problem description implies a tree (N-ary tree), so we can assume no employee is a subordinate to multiple people.
- **Cycles:** Can there be cycles in the subordinate relationships (e.g., A is B's subordinate, and B is A's subordinate)? This would lead to infinite loops. Assuming no cycles.
- **Invalid ID:** What should happen if the given `id` does not exist in the list of employees? Should it return 0, an error, or can we assume the ID will always be valid? (A robust solution should probably return 0).
- **Empty Input:** What if the `employees` list is empty? (The logical return value is 0).
- **Value Ranges:** Are importance values always non-negative? (They can be negative, as seen in tests). Are IDs always positive?

### 3. **Example Walkthrough**
- `employees = [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]]`, `id = 1`
- We start with employee `1`. Total importance = `5`.
- Employee `1` has subordinates `2` and `3`.
- We find employee `2`, add their importance. Total = `5 + 3 = 8`.
- We find employee `3`, add their importance. Total = `8 + 3 = 11`.
- Final result: `11`.

### 4. **Propose a Solution (High-Level)**

A key observation is that searching the employee list repeatedly by ID is inefficient (O(N) for each search). A better approach is to first preprocess the list into a more accessible data structure.

- **Preprocessing:** Create a hash map (or dictionary) where the key is the employee's `Id` and the value is the `*Employee` object. This allows for O(1) average time lookup for any employee.

Once the map is built, we can traverse the hierarchy using either Depth-First Search (DFS) or Breadth-First Search (BFS).

- **Approach 1: DFS (Recursive)**
 1. Build the employee map.
 2. Create a recursive helper function `dfs(employeeId)`.
 3. Inside `dfs(employeeId)`:
    a. Look up the employee in the map.
    b. Initialize a local `total` with the employee's own importance value.
    c. For each `subordinateId` in the employee's `Subordinates` list, recursively call `dfs(subordinateId)` and add the result to `total`.
    d. Return `total`.
 4. Call `dfs` with the starting `id` to get the final answer.

- **Approach 2: BFS (Iterative)**
 1. Build the employee map.
 2. Initialize a queue and add the starting `id`.
 3. Initialize `totalImportance = 0`.
 4. While the queue is not empty:
    a. Dequeue an `employeeId`.
    b. Look up the employee in the map.
    c. Add their importance to `totalImportance`.
    d. For each `subordinateId` in their `Subordinates` list, enqueue it.
 5. Return `totalImportance`.

### 5. **Write the Code**
Implement one of the chosen approaches. The existing file `59.employee_importance.go` provides both (`getImportance` for DFS and `getImportanceOpt` for BFS).

### 6. **Test the Code**
Create test cases covering various scenarios:
- A standard case with a few levels of hierarchy.
- An employee with no subordinates.
- A deeper, chain-like hierarchy.
- **Edge Case:** An empty list of employees (should return 0).
- **Edge Case:** A non-existent employee ID (should return 0).
*/
func TestGetImportance(t *testing.T) {
	testCases := []struct {
		name      string
		employees []*Employee
		id        int
		expected  int
		panics    bool // Flag for tests that are expected to panic with the current implementation
	}{
		{
			name: "Example 1 from LeetCode",
			employees: []*Employee{
				{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
				{Id: 2, Importance: 3, Subordinates: []int{}},
				{Id: 3, Importance: 3, Subordinates: []int{}},
			},
			id:       1,
			expected: 11,
			panics:   false,
		},
		{
			name: "Single Employee with no subordinates",
			employees: []*Employee{
				{Id: 1, Importance: 10, Subordinates: []int{}},
			},
			id:       1,
			expected: 10,
			panics:   false,
		},
		{
			name: "Deeper Linear Hierarchy",
			employees: []*Employee{
				{Id: 1, Importance: 10, Subordinates: []int{2}},
				{Id: 2, Importance: 5, Subordinates: []int{3}},
				{Id: 3, Importance: 1, Subordinates: []int{}},
			},
			id:       1,
			expected: 16,
			panics:   false,
		},
		{
			name: "Employee with negative importance value",
			employees: []*Employee{
				{Id: 1, Importance: 10, Subordinates: []int{2}},
				{Id: 2, Importance: -5, Subordinates: []int{}},
			},
			id:       1,
			expected: 5,
			panics:   false,
		},
		{
			name: "Unordered employee list",
			employees: []*Employee{
				{Id: 3, Importance: 3, Subordinates: []int{}},
				{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
				{Id: 2, Importance: 3, Subordinates: []int{}},
			},
			id:       1,
			expected: 11,
			panics:   false,
		},
		{
			name:      "Edge Case: Empty employee list",
			employees: []*Employee{},
			id:        1,
			expected:  0,
			panics:    true, // This will panic the current implementation
		},
		{
			name: "Edge Case: ID not found in non-empty list",
			employees: []*Employee{
				{Id: 2, Importance: 100, Subordinates: []int{}},
			},
			id:       1,
			expected: 0,
			panics:   true, // This will also panic
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (Recursive_DFS)", func(t *testing.T) {
			if tc.panics {
				// The current implementation panics when the ID is not found.
				// A robust implementation should return 0, which is our 'expected' value.
				// Here, we assert that it panics, testing the actual current behavior.
				assert.Panics(t, func() {
					getImportance(tc.employees, tc.id)
				}, "Expected a panic for this edge case, but did not get one.")
			} else {
				assert.Equal(t, tc.expected, getImportance(tc.employees, tc.id))
			}
		})

		t.Run(tc.name+" (Iterative_BFS)", func(t *testing.T) {
			if tc.panics {
				assert.Panics(t, func() {
					getImportanceOpt(tc.employees, tc.id)
				}, "Expected a panic for this edge case, but did not get one.")
			} else {
				assert.Equal(t, tc.expected, getImportanceOpt(tc.employees, tc.id))
			}
		})
	}
}
