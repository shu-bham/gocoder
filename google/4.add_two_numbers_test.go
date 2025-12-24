package google

import (
	"reflect"
	"testing"
)

// Helper function to create a ListNode from a slice of integers
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

// Helper function to convert a ListNode to a slice of integers
func listToSlice(head *ListNode) []int {
	var vals []int
	current := head
	for current != nil {
		vals = append(vals, current.Val)
		current = current.Next
	}
	return vals
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "Example 1",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "Example 2",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Example 3",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name:     "One list is nil",
			l1:       []int{1, 2, 3},
			l2:       nil,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both lists are nil",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			name:     "Both lists are nil",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := createList(tc.l1)
			l2 := createList(tc.l2)

			result := addTwoNumbers(l1, l2)
			resultSlice := listToSlice(result)

			if !reflect.DeepEqual(resultSlice, tc.expected) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tc.l1, tc.l2, resultSlice, tc.expected)
			}
		})
	}
}
