package google

import (
	"testing"
)

func TestFindOrderDFS(t *testing.T) {
	testCases := []struct {
		name        string
		numCourses  int
		prereq      [][]int
		expectCycle bool
	}{
		{
			name:        "Simple case with 2 courses",
			numCourses:  2,
			prereq:      [][]int{{1, 0}},
			expectCycle: false,
		},
		{
			name:        "No prerequisites",
			numCourses:  4,
			prereq:      [][]int{},
			expectCycle: false,
		},
		{
			name:        "Cycle detected",
			numCourses:  2,
			prereq:      [][]int{{1, 0}, {0, 1}},
			expectCycle: true,
		},
		{
			name:        "Complex graph",
			numCourses:  4,
			prereq:      [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expectCycle: false,
		},
		{
			name:        "Single course",
			numCourses:  1,
			prereq:      [][]int{},
			expectCycle: false,
		},
		{
			name:        "Zero courses",
			numCourses:  0,
			prereq:      [][]int{},
			expectCycle: false,
		},
		{
			name:        "Disconnected components",
			numCourses:  3,
			prereq:      [][]int{{1, 0}},
			expectCycle: false,
		},
		{
			name:        "Multiple prerequisites for one course",
			numCourses:  3,
			prereq:      [][]int{{0, 1}, {0, 2}},
			expectCycle: false,
		},
		{
			name:        "Long chain of prerequisites",
			numCourses:  5,
			prereq:      [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}},
			expectCycle: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			order := FindOrderDFS(tc.numCourses, tc.prereq)

			if tc.expectCycle {
				if len(order) != 0 {
					t.Errorf("expected a cycle to be detected (empty slice), but got %v", order)
				}
			} else {
				if tc.numCourses > 0 && len(order) == 0 {
					t.Errorf("expected a valid order, but got an empty slice (cycle detected?)")
					return
				}
				if tc.numCourses == 0 && len(order) != 0 {
					t.Errorf("expected empty slice for 0 courses, got %v", order)
					return
				}

				if ok, reason := isValidTopologicalSort(tc.numCourses, tc.prereq, order); !ok {
					t.Errorf("got %v which is not a valid topological sort: %s", order, reason)
				}
			}
		})
	}
}

func TestFindOrderIndegreeBFS(t *testing.T) {
	testCases := []struct {
		name        string
		numCourses  int
		prereq      [][]int
		expectCycle bool
	}{
		{
			name:        "Simple case with 2 courses",
			numCourses:  2,
			prereq:      [][]int{{1, 0}},
			expectCycle: false,
		},
		{
			name:        "No prerequisites",
			numCourses:  4,
			prereq:      [][]int{},
			expectCycle: false,
		},
		{
			name:        "Cycle detected",
			numCourses:  2,
			prereq:      [][]int{{1, 0}, {0, 1}},
			expectCycle: true,
		},
		{
			name:        "Complex graph",
			numCourses:  4,
			prereq:      [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expectCycle: false,
		},
		{
			name:        "Single course",
			numCourses:  1,
			prereq:      [][]int{},
			expectCycle: false,
		},
		{
			name:        "Zero courses",
			numCourses:  0,
			prereq:      [][]int{},
			expectCycle: false,
		},
		{
			name:        "Disconnected components",
			numCourses:  3,
			prereq:      [][]int{{1, 0}},
			expectCycle: false,
		},
		{
			name:        "Multiple prerequisites for one course",
			numCourses:  3,
			prereq:      [][]int{{0, 1}, {0, 2}},
			expectCycle: false,
		},
		{
			name:        "Long chain of prerequisites",
			numCourses:  5,
			prereq:      [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}},
			expectCycle: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			order := FindOrderIndegreeBFS(tc.numCourses, tc.prereq)

			if tc.expectCycle {
				if len(order) != 0 {
					t.Errorf("expected a cycle to be detected (empty slice), but got %v", order)
				}
			} else {
				if tc.numCourses > 0 && len(order) == 0 {
					t.Errorf("expected a valid order, but got an empty slice (cycle detected?)")
					return
				}
				if tc.numCourses == 0 && len(order) != 0 {
					t.Errorf("expected empty slice for 0 courses, got %v", order)
					return
				}

				if ok, reason := isValidTopologicalSort(tc.numCourses, tc.prereq, order); !ok {
					t.Errorf("got %v which is not a valid topological sort: %s", order, reason)
				}
			}
		})
	}
}

// isValidTopologicalSort checks if the result is a valid topological sort for the given prerequisites.
func isValidTopologicalSort(numCourses int, prereq [][]int, result []int) (bool, string) {
	if len(result) != numCourses {
		return false, "result length is not equal to numCourses"
	}

	pos := make(map[int]int, numCourses)
	for i, course := range result {
		if _, exists := pos[course]; exists {
			return false, "duplicate course in result"
		}
		pos[course] = i
	}

	if len(pos) != numCourses {
		// This case is technically covered by the first check if there are no duplicates,
		// but it's good for logical completeness.
		return false, "not all courses are present in result"
	}

	for _, p := range prereq {
		course := p[0]
		pre := p[1]
		posCourse, okCourse := pos[course]
		posPre, okPre := pos[pre]

		if !okCourse || !okPre {
			// This should not happen if all courses are in the result
			return false, "course from prerequisite not in result"
		}

		if posPre > posCourse {
			return false, "prerequisite order violated"
		}
	}
	return true, ""
}
